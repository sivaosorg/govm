package match

import (
	"unicode/utf8"
)

type result int

// Match returns true if str matches pattern. This is a very
// simple wildcard match where '*' matches on any number characters
// and '?' matches on any one character.
//
// pattern:
//
//	{ term }
//
// term:
//
//	'*'         matches any sequence of non-Separator characters
//	'?'         matches any single non-Separator character
//	c           matches character c (c != '*', '?', '\\')
//	'\\' c      matches character c
func Match(str, pattern string) bool {
	if pattern == "*" {
		return true
	}
	return match(str, pattern, 0, nil, -1) == rightMatch
}

// MatchLimit is the same as Match but will limit the complexity of the match
// operation. This is to avoid long running matches, specifically to avoid ReDos
// attacks from arbitrary inputs.
//
// How it works:
// The underlying match routine is recursive and may call itself when it
// encounters a sandwiched wildcard pattern, such as: `user:*:name`.
// Every time it calls itself a counter is incremented.
// The operation is stopped when counter > maxComplexity*len(str).
func MatchLimit(str, pattern string, maxComplexity int) (matched, stopped bool) {
	if pattern == "*" {
		return true, false
	}
	counter := 0
	r := match(str, pattern, len(str), &counter, maxComplexity)
	if r == rightStop {
		return false, true
	}
	return r == rightMatch, false
}

func match(text, pattern string, lenString int, counter *int, maxComplexity int) result {
	if maxComplexity > -1 {
		if *counter > lenString*maxComplexity {
			return rightStop
		}
		*counter++
	}

	for len(pattern) > 0 {
		var wild bool
		pc, ps := rune(pattern[0]), 1
		if pc > 0x7f {
			pc, ps = utf8.DecodeRuneInString(pattern)
		}
		var sc rune
		var ss int
		if len(text) > 0 {
			sc, ss = rune(text[0]), 1
			if sc > 0x7f {
				sc, ss = utf8.DecodeRuneInString(text)
			}
		}
		switch pc {
		case '?':
			if ss == 0 {
				return rightNoMatch
			}
		case '*':
			// Ignore repeating stars.
			for len(pattern) > 1 && pattern[1] == '*' {
				pattern = pattern[1:]
			}

			// If this star is the last character then it must be a match.
			if len(pattern) == 1 {
				return rightMatch
			}

			// Match and trim any non-wildcard suffix characters.
			var ok bool
			text, pattern, ok = matchTrimSuffix(text, pattern)
			if !ok {
				return rightNoMatch
			}

			// Check for single star again.
			if len(pattern) == 1 {
				return rightMatch
			}

			// Perform recursive wildcard search.
			r := match(text, pattern[1:], lenString, counter, maxComplexity)
			if r != rightNoMatch {
				return r
			}
			if len(text) == 0 {
				return rightNoMatch
			}
			wild = true
		default:
			if ss == 0 {
				return rightNoMatch
			}
			if pc == '\\' {
				pattern = pattern[ps:]
				pc, ps = utf8.DecodeRuneInString(pattern)
				if ps == 0 {
					return rightNoMatch
				}
			}
			if sc != pc {
				return rightNoMatch
			}
		}
		text = text[ss:]
		if !wild {
			pattern = pattern[ps:]
		}
	}
	if len(text) == 0 {
		return rightMatch
	}
	return rightNoMatch
}

// matchTrimSuffix matches and trims any non-wildcard suffix characters.
// Returns the trimmed string and pattern.
//
// This is called because the pattern contains extra data after the wildcard
// star. Here we compare any suffix characters in the pattern to the suffix of
// the target string. Basically a reverse match that stops when a wildcard
// character is reached. This is a little trickier than a forward match because
// we need to evaluate an escaped character in reverse.
//
// Any matched characters will be trimmed from both the target
// string and the pattern.
func matchTrimSuffix(text, pattern string) (string, string, bool) {
	// It's expected that the pattern has at least two bytes and the first byte
	// is a wildcard star '*'
	match := true
	for len(text) > 0 && len(pattern) > 1 {
		pc, ps := utf8.DecodeLastRuneInString(pattern)
		var esc bool
		for i := 0; ; i++ {
			if pattern[len(pattern)-ps-i-1] != '\\' {
				if i&1 == 1 {
					esc = true
					ps++
				}
				break
			}
		}
		if pc == '*' && !esc {
			match = true
			break
		}
		sc, ss := utf8.DecodeLastRuneInString(text)
		if !((pc == '?' && !esc) || pc == sc) {
			match = false
			break
		}
		text = text[:len(text)-ss]
		pattern = pattern[:len(pattern)-ps]
	}
	return text, pattern, match
}

// Allowable parses the pattern and determines the minimum and maximum allowable
// values that the pattern can represent.
// When the max cannot be determined, 'true' will be returned
// for infinite.
func Allowable(pattern string) (min, max string) {
	if pattern == "" || pattern[0] == '*' {
		return "", ""
	}

	minVal := make([]byte, 0, len(pattern))
	maxVal := make([]byte, 0, len(pattern))
	var wild bool
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '*' {
			wild = true
			break
		}
		if pattern[i] == '?' {
			minVal = append(minVal, 0)
			maxVal = append(maxVal, maxRuneBytes[:]...)
		} else {
			minVal = append(minVal, pattern[i])
			maxVal = append(maxVal, pattern[i])
		}
	}
	if wild {
		r, n := utf8.DecodeLastRune(maxVal)
		if r != utf8.RuneError {
			if r < utf8.MaxRune {
				r++
				if r > 0x7f {
					b := make([]byte, 4)
					nn := utf8.EncodeRune(b, r)
					maxVal = append(maxVal[:len(maxVal)-n], b[:nn]...)
				} else {
					maxVal = append(maxVal[:len(maxVal)-n], byte(r))
				}
			}
		}
	}
	return string(minVal), string(maxVal)
}

// IsPattern returns true if the string is a pattern.
func IsPattern(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] == '*' || value[i] == '?' {
			return true
		}
	}
	return false
}
