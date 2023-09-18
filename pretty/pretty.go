package pretty

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"
)

// Pretty converts the input json into a more human readable format where each
// element is on it's own line with clear indentation.
func Pretty(json []byte) []byte { return PrettyOptions(json, nil) }

// PrettyOptions is like Pretty but with customized options.
func PrettyOptions(json []byte, option *OptionsConfig) []byte {
	if option == nil {
		option = DefaultOptionsConfig
	}
	buf := make([]byte, 0, len(json))
	if len(option.Prefix) != 0 {
		buf = append(buf, option.Prefix...)
	}
	buf, _, _, _ = appendPrettyAny(buf, json, 0, true,
		option.Width, option.Prefix, option.Indent, option.SortKeys,
		0, 0, -1)
	if len(buf) > 0 {
		buf = append(buf, '\n')
	}
	return buf
}

// Ugly removes insignificant space characters from the input json byte slice
// and returns the compacted result.
func Ugly(json []byte) []byte {
	buf := make([]byte, 0, len(json))
	return ugly(buf, json)
}

// UglyInPlace removes insignificant space characters from the input json
// byte slice and returns the compacted result. This method reuses the
// input json buffer to avoid allocations. Do not use the original bytes
// slice upon return.
func UglyInPlace(json []byte) []byte { return ugly(json, json) }

func ugly(dst, src []byte) []byte {
	dst = dst[:0]
	for i := 0; i < len(src); i++ {
		if src[i] > ' ' {
			dst = append(dst, src[i])
			if src[i] == '"' {
				for i = i + 1; i < len(src); i++ {
					dst = append(dst, src[i])
					if src[i] == '"' {
						j := i - 1
						for ; ; j-- {
							if src[j] != '\\' {
								break
							}
						}
						if (j-i)%2 != 0 {
							break
						}
					}
				}
			}
		}
	}
	return dst
}

func isNaNOrInf(src []byte) bool {
	return src[0] == 'i' || //Inf
		src[0] == 'I' || // inf
		src[0] == '+' || // +Inf
		src[0] == 'N' || // Nan
		(src[0] == 'n' && len(src) > 1 && src[1] != 'u') // nan
}

func appendPrettyAny(buf, json []byte, i int, pretty bool, width int, prefix, indent string, sortKeys bool, tabs, nl, max int) ([]byte, int, int, bool) {
	for ; i < len(json); i++ {
		if json[i] <= ' ' {
			continue
		}
		if json[i] == '"' {
			return appendPrettyString(buf, json, i, nl)
		}

		if (json[i] >= '0' && json[i] <= '9') || json[i] == '-' || isNaNOrInf(json[i:]) {
			return appendPrettyNumber(buf, json, i, nl)
		}
		if json[i] == '{' {
			return appendPrettyObject(buf, json, i, '{', '}', pretty, width, prefix, indent, sortKeys, tabs, nl, max)
		}
		if json[i] == '[' {
			return appendPrettyObject(buf, json, i, '[', ']', pretty, width, prefix, indent, sortKeys, tabs, nl, max)
		}
		switch json[i] {
		case 't':
			return append(buf, 't', 'r', 'u', 'e'), i + 4, nl, true
		case 'f':
			return append(buf, 'f', 'a', 'l', 's', 'e'), i + 5, nl, true
		case 'n':
			return append(buf, 'n', 'u', 'l', 'l'), i + 4, nl, true
		}
	}
	return buf, i, nl, true
}

type pair struct {
	keyStart, keyEnd     int
	valueStart, valueEnd int
}

type byKeyVal struct {
	sorted bool
	json   []byte
	buf    []byte
	pairs  []pair
}

func (b *byKeyVal) Len() int {
	return len(b.pairs)
}

func (b *byKeyVal) Less(i, j int) bool {
	if b.isLess(i, j, byKey) {
		return true
	}
	if b.isLess(j, i, byKey) {
		return false
	}
	return b.isLess(i, j, byVal)
}

func (b *byKeyVal) Swap(i, j int) {
	b.pairs[i], b.pairs[j] = b.pairs[j], b.pairs[i]
	b.sorted = true
}

type byKind int

const (
	byKey byKind = 0
	byVal byKind = 1
)

type jsonType int

const (
	jsonNull jsonType = iota
	jsonFalse
	jNumber
	jsonString
	jsonTrue
	jsonJson
)

func getJsonType(v []byte) jsonType {
	if len(v) == 0 {
		return jsonNull
	}
	switch v[0] {
	case '"':
		return jsonString
	case 'f':
		return jsonFalse
	case 't':
		return jsonTrue
	case 'n':
		return jsonNull
	case '[', '{':
		return jsonJson
	default:
		return jNumber
	}
}

func (a *byKeyVal) isLess(i, j int, kind byKind) bool {
	k1 := a.json[a.pairs[i].keyStart:a.pairs[i].keyEnd]
	k2 := a.json[a.pairs[j].keyStart:a.pairs[j].keyEnd]
	var v1, v2 []byte
	if kind == byKey {
		v1 = k1
		v2 = k2
	} else {
		v1 = bytes.TrimSpace(a.buf[a.pairs[i].valueStart:a.pairs[i].valueEnd])
		v2 = bytes.TrimSpace(a.buf[a.pairs[j].valueStart:a.pairs[j].valueEnd])
		if len(v1) >= len(k1)+1 {
			v1 = bytes.TrimSpace(v1[len(k1)+1:])
		}
		if len(v2) >= len(k2)+1 {
			v2 = bytes.TrimSpace(v2[len(k2)+1:])
		}
	}
	t1 := getJsonType(v1)
	t2 := getJsonType(v2)
	if t1 < t2 {
		return true
	}
	if t1 > t2 {
		return false
	}
	if t1 == jsonString {
		s1 := parseString(v1)
		s2 := parseString(v2)
		return string(s1) < string(s2)
	}
	if t1 == jNumber {
		n1, _ := strconv.ParseFloat(string(v1), 64)
		n2, _ := strconv.ParseFloat(string(v2), 64)
		return n1 < n2
	}
	return string(v1) < string(v2)

}

func parseString(s []byte) []byte {
	for i := 1; i < len(s); i++ {
		if s[i] == '\\' {
			var str string
			json.Unmarshal(s, &str)
			return []byte(str)
		}
		if s[i] == '"' {
			return s[1:i]
		}
	}
	return nil
}

func appendPrettyObject(buf, json []byte, i int, open, close byte, pretty bool, width int, prefix, indent string, sortKeys bool, tabs, nl, max int) ([]byte, int, int, bool) {
	var ok bool
	if width > 0 {
		if pretty && open == '[' && max == -1 {
			// here we try to create a single line array
			max := width - (len(buf) - nl)
			if max > 3 {
				s1, s2 := len(buf), i
				buf, i, _, ok = appendPrettyObject(buf, json, i, '[', ']', false, width, prefix, "", sortKeys, 0, 0, max)
				if ok && len(buf)-s1 <= max {
					return buf, i, nl, true
				}
				buf = buf[:s1]
				i = s2
			}
		} else if max != -1 && open == '{' {
			return buf, i, nl, false
		}
	}
	buf = append(buf, open)
	i++
	var pairs []pair
	if open == '{' && sortKeys {
		pairs = make([]pair, 0, 8)
	}
	var n int
	for ; i < len(json); i++ {
		if json[i] <= ' ' {
			continue
		}
		if json[i] == close {
			if pretty {
				if open == '{' && sortKeys {
					buf = sortPairs(json, buf, pairs)
				}
				if n > 0 {
					nl = len(buf)
					if buf[nl-1] == ' ' {
						buf[nl-1] = '\n'
					} else {
						buf = append(buf, '\n')
					}
				}
				if buf[len(buf)-1] != open {
					buf = appendTabs(buf, prefix, indent, tabs)
				}
			}
			buf = append(buf, close)
			return buf, i + 1, nl, open != '{'
		}
		if open == '[' || json[i] == '"' {
			if n > 0 {
				buf = append(buf, ',')
				if width != -1 && open == '[' {
					buf = append(buf, ' ')
				}
			}
			var p pair
			if pretty {
				nl = len(buf)
				if buf[nl-1] == ' ' {
					buf[nl-1] = '\n'
				} else {
					buf = append(buf, '\n')
				}
				if open == '{' && sortKeys {
					p.keyStart = i
					p.valueStart = len(buf)
				}
				buf = appendTabs(buf, prefix, indent, tabs+1)
			}
			if open == '{' {
				buf, i, nl, _ = appendPrettyString(buf, json, i, nl)
				if sortKeys {
					p.keyEnd = i
				}
				buf = append(buf, ':')
				if pretty {
					buf = append(buf, ' ')
				}
			}
			buf, i, nl, ok = appendPrettyAny(buf, json, i, pretty, width, prefix, indent, sortKeys, tabs+1, nl, max)
			if max != -1 && !ok {
				return buf, i, nl, false
			}
			if pretty && open == '{' && sortKeys {
				p.valueEnd = len(buf)
				if p.keyStart > p.keyEnd || p.valueStart > p.valueEnd {
					// bad data. disable sorting
					sortKeys = false
				} else {
					pairs = append(pairs, p)
				}
			}
			i--
			n++
		}
	}
	return buf, i, nl, open != '{'
}

func sortPairs(json, buf []byte, pairs []pair) []byte {
	if len(pairs) == 0 {
		return buf
	}
	_valStart := pairs[0].valueStart
	_valEnd := pairs[len(pairs)-1].valueEnd
	_keyVal := byKeyVal{false, json, buf, pairs}
	sort.Stable(&_keyVal)
	if !_keyVal.sorted {
		return buf
	}
	n := make([]byte, 0, _valEnd-_valStart)
	for i, p := range pairs {
		n = append(n, buf[p.valueStart:p.valueEnd]...)
		if i < len(pairs)-1 {
			n = append(n, ',')
			n = append(n, '\n')
		}
	}
	return append(buf[:_valStart], n...)
}

func appendPrettyString(buf, json []byte, i, nl int) ([]byte, int, int, bool) {
	s := i
	i++
	for ; i < len(json); i++ {
		if json[i] == '"' {
			var sc int
			for j := i - 1; j > s; j-- {
				if json[j] == '\\' {
					sc++
				} else {
					break
				}
			}
			if sc%2 == 1 {
				continue
			}
			i++
			break
		}
	}
	return append(buf, json[s:i]...), i, nl, true
}

func appendPrettyNumber(buf, json []byte, i, nl int) ([]byte, int, int, bool) {
	s := i
	i++
	for ; i < len(json); i++ {
		if json[i] <= ' ' || json[i] == ',' || json[i] == ':' || json[i] == ']' || json[i] == '}' {
			break
		}
	}
	return append(buf, json[s:i]...), i, nl, true
}

func appendTabs(buf []byte, prefix, indent string, tabs int) []byte {
	if len(prefix) != 0 {
		buf = append(buf, prefix...)
	}
	if len(indent) == 2 && indent[0] == ' ' && indent[1] == ' ' {
		for i := 0; i < tabs; i++ {
			buf = append(buf, ' ', ' ')
		}
	} else {
		for i := 0; i < tabs; i++ {
			buf = append(buf, indent...)
		}
	}
	return buf
}

// Style is the color style
type Style struct {
	Key, String, Number [2]string
	True, False, Null   [2]string
	Escape              [2]string
	Brackets            [2]string
	Append              func(dst []byte, c byte) []byte
}

func hexp(p byte) byte {
	switch {
	case p < 10:
		return p + '0'
	default:
		return (p - 10) + 'a'
	}
}

// TerminalStyle is for terminals
var TerminalStyle *Style

func init() {
	TerminalStyle = &Style{
		Key:      [2]string{"\x1B[1m\x1B[94m", "\x1B[0m"},
		String:   [2]string{"\x1B[32m", "\x1B[0m"},
		Number:   [2]string{"\x1B[33m", "\x1B[0m"},
		True:     [2]string{"\x1B[36m", "\x1B[0m"},
		False:    [2]string{"\x1B[36m", "\x1B[0m"},
		Null:     [2]string{"\x1B[2m", "\x1B[0m"},
		Escape:   [2]string{"\x1B[35m", "\x1B[0m"},
		Brackets: [2]string{"\x1B[1m", "\x1B[0m"},
		Append: func(dst []byte, c byte) []byte {
			if c < ' ' && (c != '\r' && c != '\n' && c != '\t' && c != '\v') {
				dst = append(dst, "\\u00"...)
				dst = append(dst, hexp((c>>4)&0xF))
				return append(dst, hexp((c)&0xF))
			}
			return append(dst, c)
		},
	}
}

// Color will colorize the json. The style parma is used for customizing
// the colors. Passing nil to the style param will use the default
// TerminalStyle.
func Color(src []byte, style *Style) []byte {
	if style == nil {
		style = TerminalStyle
	}
	appendStyle := style.Append
	if appendStyle == nil {
		appendStyle = func(dst []byte, c byte) []byte {
			return append(dst, c)
		}
	}
	type innerStack struct {
		kind byte
		key  bool
	}
	var dst []byte
	var stack []innerStack
	for i := 0; i < len(src); i++ {
		if src[i] == '"' {
			key := len(stack) > 0 && stack[len(stack)-1].key
			if key {
				dst = append(dst, style.Key[0]...)
			} else {
				dst = append(dst, style.String[0]...)
			}
			dst = appendStyle(dst, '"')
			esc := false
			useEsc := 0
			for i = i + 1; i < len(src); i++ {
				if src[i] == '\\' {
					if key {
						dst = append(dst, style.Key[1]...)
					} else {
						dst = append(dst, style.String[1]...)
					}
					dst = append(dst, style.Escape[0]...)
					dst = appendStyle(dst, src[i])
					esc = true
					if i+1 < len(src) && src[i+1] == 'u' {
						useEsc = 5
					} else {
						useEsc = 1
					}
				} else if esc {
					dst = appendStyle(dst, src[i])
					if useEsc == 1 {
						esc = false
						dst = append(dst, style.Escape[1]...)
						if key {
							dst = append(dst, style.Key[0]...)
						} else {
							dst = append(dst, style.String[0]...)
						}
					} else {
						useEsc--
					}
				} else {
					dst = appendStyle(dst, src[i])
				}
				if src[i] == '"' {
					j := i - 1
					for ; ; j-- {
						if src[j] != '\\' {
							break
						}
					}
					if (j-i)%2 != 0 {
						break
					}
				}
			}
			if esc {
				dst = append(dst, style.Escape[1]...)
			} else if key {
				dst = append(dst, style.Key[1]...)
			} else {
				dst = append(dst, style.String[1]...)
			}
		} else if src[i] == '{' || src[i] == '[' {
			stack = append(stack, innerStack{src[i], src[i] == '{'})
			dst = append(dst, style.Brackets[0]...)
			dst = appendStyle(dst, src[i])
			dst = append(dst, style.Brackets[1]...)
		} else if (src[i] == '}' || src[i] == ']') && len(stack) > 0 {
			stack = stack[:len(stack)-1]
			dst = append(dst, style.Brackets[0]...)
			dst = appendStyle(dst, src[i])
			dst = append(dst, style.Brackets[1]...)
		} else if (src[i] == ':' || src[i] == ',') && len(stack) > 0 && stack[len(stack)-1].kind == '{' {
			stack[len(stack)-1].key = !stack[len(stack)-1].key
			dst = append(dst, style.Brackets[0]...)
			dst = appendStyle(dst, src[i])
			dst = append(dst, style.Brackets[1]...)
		} else {
			var kind byte
			if (src[i] >= '0' && src[i] <= '9') || src[i] == '-' || isNaNOrInf(src[i:]) {
				kind = '0'
				dst = append(dst, style.Number[0]...)
			} else if src[i] == 't' {
				kind = 't'
				dst = append(dst, style.True[0]...)
			} else if src[i] == 'f' {
				kind = 'f'
				dst = append(dst, style.False[0]...)
			} else if src[i] == 'n' {
				kind = 'n'
				dst = append(dst, style.Null[0]...)
			} else {
				dst = appendStyle(dst, src[i])
			}
			if kind != 0 {
				for ; i < len(src); i++ {
					if src[i] <= ' ' || src[i] == ',' || src[i] == ':' || src[i] == ']' || src[i] == '}' {
						i--
						break
					}
					dst = appendStyle(dst, src[i])
				}
				if kind == '0' {
					dst = append(dst, style.Number[1]...)
				} else if kind == 't' {
					dst = append(dst, style.True[1]...)
				} else if kind == 'f' {
					dst = append(dst, style.False[1]...)
				} else if kind == 'n' {
					dst = append(dst, style.Null[1]...)
				}
			}
		}
	}
	return dst
}

// Spec strips out comments and trailing commas and convert the input to a
// valid JSON per the official spec: https://tools.ietf.org/html/rfc8259
//
// The resulting JSON will always be the same length as the input and it will
// include all of the same line breaks at matching offsets. This is to ensure
// the result can be later processed by a external parser and that that
// parser will report messages or errors with the correct offsets.
func Spec(src []byte) []byte {
	return spec(src, nil)
}

// SpecInPlace is the same as Spec, but this method reuses the input json
// buffer to avoid allocations. Do not use the original bytes slice upon return.
func SpecInPlace(src []byte) []byte {
	return spec(src, src)
}

func spec(src, dst []byte) []byte {
	dst = dst[:0]
	for i := 0; i < len(src); i++ {
		if src[i] == '/' {
			if i < len(src)-1 {
				if src[i+1] == '/' {
					dst = append(dst, ' ', ' ')
					i += 2
					for ; i < len(src); i++ {
						if src[i] == '\n' {
							dst = append(dst, '\n')
							break
						} else if src[i] == '\t' || src[i] == '\r' {
							dst = append(dst, src[i])
						} else {
							dst = append(dst, ' ')
						}
					}
					continue
				}
				if src[i+1] == '*' {
					dst = append(dst, ' ', ' ')
					i += 2
					for ; i < len(src)-1; i++ {
						if src[i] == '*' && src[i+1] == '/' {
							dst = append(dst, ' ', ' ')
							i++
							break
						} else if src[i] == '\n' || src[i] == '\t' ||
							src[i] == '\r' {
							dst = append(dst, src[i])
						} else {
							dst = append(dst, ' ')
						}
					}
					continue
				}
			}
		}
		dst = append(dst, src[i])
		if src[i] == '"' {
			for i = i + 1; i < len(src); i++ {
				dst = append(dst, src[i])
				if src[i] == '"' {
					j := i - 1
					for ; ; j-- {
						if src[j] != '\\' {
							break
						}
					}
					if (j-i)%2 != 0 {
						break
					}
				}
			}
		} else if src[i] == '}' || src[i] == ']' {
			for j := len(dst) - 2; j >= 0; j-- {
				if dst[j] <= ' ' {
					continue
				}
				if dst[j] == ',' {
					dst[j] = ' '
				}
				break
			}
		}
	}
	return dst
}
