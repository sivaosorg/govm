package timex

import "regexp"

const (
	DateTimeFormatYearMonthDTimezoneHourMinuteSecond    = "2006-01-02T15:04:05"
	DateTimeFormatYearMonthDay                          = "2006-01-02"
	DateTimeFormYearMonthDayHourMinuteSecond            = "2006-01-02 15:04:05"
	DateTimeFormRFC3339NanoTimezone                     = "2006-01-02 15:04:05.999999999 -07:00"
	DateTimeFormatDotYearMonthDayHourMinuteSecond       = "2006.01.02.15.04.05"
	DateTimeFormatCommaDashYearMonthDayHourMinuteSecond = "2006:01:02_15:04:05"
)

const (
	FormatHourMinuteSecond    = "15:04:05"
	DotFormatHourMinuteSecond = "15.04.05"
)

var HasTimeRegexp = regexp.MustCompile(`(\s+|^\s*|T)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))(\s*$|[Z+-])`) // match 15:04:05, 15:04:05.000, 15:04:05.000000 15, 2017-01-01 15:04, 2021-07-20T00:59:10Z, 2021-07-20T00:59:10+08:00, 2021-07-20T00:00:10-07:00 etc
var OnlyTimeRegexp = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)                // match 15:04:05, 15, 15:04:05.000, 15:04:05.000000, etc
