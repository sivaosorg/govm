package timex

import (
	"fmt"
	"time"
)

func GetBeginDay(at time.Time) time.Time {
	return time.Date(at.Year(), at.Month(), at.Day(), 0, 0, 0, 0, at.Local().Location())
}

func GetEndDay(at time.Time) time.Time {
	return time.Date(at.Year(), at.Month(), at.Day(), 23, 59, 59, 0, at.Local().Location())
}

// GetPreviousBeginDay
// Get previous day with begin day
// Example: today is 2023-03-17, so then previous day is 2023-03-16 00:00:00.0
func GetPreviousBeginDay(at time.Time, previousDay int) time.Time {
	lastDay := at.AddDate(0, 0, previousDay*(-1))
	return GetBeginDay(lastDay)
}

// GetPreviousEndDay
// Get previous day with end day
// Example: today is 2023-03-17, so then previous end day is 2023-03-16 23:59:59.0
func GetPreviousEndDay(at time.Time, previousDay int) time.Time {
	lastDay := at.AddDate(0, 0, previousDay*(-1))
	return GetEndDay(lastDay)
}

// SetTimezoneX
func SetTimezoneX(at time.Time, timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	now := at.In(loc)
	return now, err
}

// SetTimezone
func SetTimezone(at time.Time, timezone string) time.Time {
	t, err := SetTimezoneX(at, timezone)
	if err != nil {
		return at
	}
	return t
}

// AddSeconds
func AddSeconds(at time.Time, numberOfSeconds int) time.Time {
	if numberOfSeconds == 0 {
		return at
	}
	return at.Add(time.Second * time.Duration(numberOfSeconds))
}

// AddMinutes
func AddMinutes(at time.Time, numberOfMinutes int) time.Time {
	if numberOfMinutes == 0 {
		return at
	}
	return at.Add(time.Minute * time.Duration(numberOfMinutes))
}

// AddHours
func AddHours(at time.Time, numberOfHours int) time.Time {
	if numberOfHours == 0 {
		return at
	}
	return at.Add(time.Hour * time.Duration(numberOfHours))
}

// AddDays
func AddDays(at time.Time, numberOfDays int) time.Time {
	if numberOfDays == 0 {
		return at
	}
	return at.Add(time.Hour * 24 * time.Duration(numberOfDays))
}

func HoursSince(t time.Time) float64 {
	duration := time.Since(t)
	hours := duration.Hours()
	return hours
}

func MinutesSince(t time.Time) float64 {
	duration := time.Since(t)
	minutes := duration.Minutes()
	return minutes
}

func SecondsSince(t time.Time) float64 {
	duration := time.Since(t)
	seconds := duration.Seconds()
	return seconds
}

func OnTime(t time.Time) bool {
	targetTime := time.Now()
	tolerance := time.Minute
	diff := t.Sub(targetTime)
	return diff >= -tolerance && diff <= tolerance
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func isLeap(t time.Time) bool {
	return isLeapYear(t.Year())
}

func ListWeekdays(start time.Time, end time.Time) []time.Time {
	var weekdays []time.Time
	for current := start; current.Before(end) || current.Equal(end); current = current.AddDate(0, 0, 1) {
		d := current.Weekday()
		if d != time.Sunday && d != time.Saturday {
			y := current.Year()
			if isLeapYear(y) && current.Month() == time.February && current.Day() == 29 {
				weekdays = append(weekdays, current)
			} else if !isLeapYear(y) && current.Day() <= time.Date(y, time.December, 31, 0, 0, 0, 0, time.UTC).Day() {
				weekdays = append(weekdays, current)
			}
		}
	}
	return weekdays
}

// New initialize Now based on configuration
func (config *Config) With(t time.Time) *Timex {
	return &Timex{Time: t, Config: config}
}

// Parse parse string to time based on configuration
func (config *Config) Parse(s ...string) (time.Time, error) {
	if config.TimeLocation == nil {
		return config.With(time.Now()).Parse(s...)
	} else {
		return config.With(time.Now().In(config.TimeLocation)).Parse(s...)
	}
}

// MustParse must parse string to time or will panic
func (config *Config) MustParse(s ...string) time.Time {
	if config.TimeLocation == nil {
		return config.With(time.Now()).MustParse(s...)
	} else {
		return config.With(time.Now().In(config.TimeLocation)).MustParse(s...)
	}
}

// With initialize Now with time
func With(t time.Time) *Timex {
	config := DefaultConfig
	if config == nil {
		config = &Config{
			WeekStartDay: WeekStartDay,
			TimeFormats:  TimeFormats,
		}
	}
	return &Timex{Time: t, Config: config}
}

// New initialize Now with time
func New(t time.Time) *Timex {
	return With(t)
}

// BeginningOfMinute beginning of minute
func BeginningOfMinute() time.Time {
	return With(time.Now()).BeginningOfMinute()
}

// BeginningOfHour beginning of hour
func BeginningOfHour() time.Time {
	return With(time.Now()).BeginningOfHour()
}

// BeginningOfDay beginning of day
func BeginningOfDay() time.Time {
	return With(time.Now()).BeginningOfDay()
}

// BeginningOfWeek beginning of week
func BeginningOfWeek() time.Time {
	return With(time.Now()).BeginningOfWeek()
}

// BeginningOfMonth beginning of month
func BeginningOfMonth() time.Time {
	return With(time.Now()).BeginningOfMonth()
}

// BeginningOfQuarter beginning of quarter
func BeginningOfQuarter() time.Time {
	return With(time.Now()).BeginningOfQuarter()
}

// BeginningOfYear beginning of year
func BeginningOfYear() time.Time {
	return With(time.Now()).BeginningOfYear()
}

// EndOfMinute end of minute
func EndOfMinute() time.Time {
	return With(time.Now()).EndOfMinute()
}

// EndOfHour end of hour
func EndOfHour() time.Time {
	return With(time.Now()).EndOfHour()
}

// EndOfDay end of day
func EndOfDay() time.Time {
	return With(time.Now()).EndOfDay()
}

// EndOfWeek end of week
func EndOfWeek() time.Time {
	return With(time.Now()).EndOfWeek()
}

// EndOfMonth end of month
func EndOfMonth() time.Time {
	return With(time.Now()).EndOfMonth()
}

// EndOfQuarter end of quarter
func EndOfQuarter() time.Time {
	return With(time.Now()).EndOfQuarter()
}

// EndOfYear end of year
func EndOfYear() time.Time {
	return With(time.Now()).EndOfYear()
}

// Monday monday

func Monday(s ...string) time.Time {
	return With(time.Now()).Monday(s...)
}

// Sunday sunday
func Sunday(s ...string) time.Time {
	return With(time.Now()).Sunday(s...)
}

// EndOfSunday end of sunday
func EndOfSunday() time.Time {
	return With(time.Now()).EndOfSunday()
}

// Quarter returns the yearly quarter
func Quarter() uint {
	return With(time.Now()).Quarter()
}

// Parse parse string to time
func Parse(s ...string) (time.Time, error) {
	return With(time.Now()).Parse(s...)
}

// ParseInLocation parse string to time in location
func ParseInLocation(loc *time.Location, s ...string) (time.Time, error) {
	return With(time.Now().In(loc)).Parse(s...)
}

// MustParse must parse string to time or will panic
func MustParse(s ...string) time.Time {
	return With(time.Now()).MustParse(s...)
}

// MustParseInLocation must parse string to time in location or will panic
func MustParseInLocation(loc *time.Location, s ...string) time.Time {
	return With(time.Now().In(loc)).MustParse(s...)
}

// Between check now between the begin, end time or not
func Between(time1, time2 string) bool {
	return With(time.Now()).Between(time1, time2)
}

// BeginningOfMinute beginning of minute
func (now *Timex) BeginningOfMinute() time.Time {
	return now.Truncate(time.Minute)
}

// BeginningOfHour beginning of hour
func (now *Timex) BeginningOfHour() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, now.Time.Hour(), 0, 0, 0, now.Time.Location())
}

// BeginningOfDay beginning of day
func (now *Timex) BeginningOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Time.Location())
}

// BeginningOfWeek beginning of week
func (now *Timex) BeginningOfWeek() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())

	if now.WeekStartDay != time.Sunday {
		weekStartDayInt := int(now.WeekStartDay)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return t.AddDate(0, 0, -weekday)
}

// BeginningOfMonth beginning of month
func (now *Timex) BeginningOfMonth() time.Time {
	y, m, _ := now.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
}

// BeginningOfQuarter beginning of quarter
func (now *Timex) BeginningOfQuarter() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

// BeginningOfHalf beginning of half year
func (now *Timex) BeginningOfHalf() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

// BeginningOfYear BeginningOfYear beginning of year
func (now *Timex) BeginningOfYear() time.Time {
	y, _, _ := now.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, now.Location())
}

// EndOfMinute end of minute
func (now *Timex) EndOfMinute() time.Time {
	return now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

// EndOfHour end of hour
func (now *Timex) EndOfHour() time.Time {
	return now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

// EndOfDay end of day
func (now *Timex) EndOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
}

// EndOfWeek end of week
func (now *Timex) EndOfWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// EndOfMonth end of month
func (now *Timex) EndOfMonth() time.Time {
	return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfQuarter end of quarter
func (now *Timex) EndOfQuarter() time.Time {
	return now.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

// EndOfHalf end of half year
func (now *Timex) EndOfHalf() time.Time {
	return now.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond)
}

// EndOfYear end of year
func (now *Timex) EndOfYear() time.Time {
	return now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

func (now *Timex) Monday(s ...string) time.Time {
	var parseTime time.Time
	var err error
	if len(s) > 0 {
		parseTime, err = now.Parse(s...)
		if err != nil {
			panic(err)
		}
	} else {
		parseTime = now.BeginningOfDay()
	}
	weekday := int(parseTime.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return parseTime.AddDate(0, 0, -weekday+1)
}

func (now *Timex) Sunday(s ...string) time.Time {
	var parseTime time.Time
	var err error
	if len(s) > 0 {
		parseTime, err = now.Parse(s...)
		if err != nil {
			panic(err)
		}
	} else {
		parseTime = now.BeginningOfDay()
	}
	weekday := int(parseTime.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return parseTime.AddDate(0, 0, (7 - weekday))
}

// EndOfSunday end of sunday
func (now *Timex) EndOfSunday() time.Time {
	return New(now.Sunday()).EndOfDay()
}

// Quarter returns the yearly quarter
func (now *Timex) Quarter() uint {
	return (uint(now.Month())-1)/3 + 1
}

func (now *Timex) parseWithFormat(s string, location *time.Location) (t time.Time, err error) {
	for _, format := range now.TimeFormats {
		t, err = time.ParseInLocation(format, s, location)

		if err == nil {
			return
		}
	}
	err = fmt.Errorf("Can't parse string as time: %v", s)
	return
}

// Parse parse string to time
func (now *Timex) Parse(s ...string) (t time.Time, err error) {
	var (
		setCurrentTime  bool
		parseTime       []int
		currentLocation = now.Location()
		onlyTimeInStr   = true
		currentTime     = FormatTimex(now.Time)
	)

	for _, str := range s {
		hasTimeInStr := HasTimeRegexp.MatchString(str) // match 15:04:05, 15
		onlyTimeInStr = hasTimeInStr && onlyTimeInStr && OnlyTimeRegexp.MatchString(str)
		if t, err = now.parseWithFormat(str, currentLocation); err == nil {
			location := t.Location()
			parseTime = FormatTimex(t)

			for i, v := range parseTime {
				// Don't reset hour, minute, second if current time str including time
				if hasTimeInStr && i <= 3 {
					continue
				}

				// If value is zero, replace it with current time
				if v == 0 {
					if setCurrentTime {
						parseTime[i] = currentTime[i]
					}
				} else {
					setCurrentTime = true
				}

				// if current time only includes time, should change day, month to current time
				if onlyTimeInStr {
					if i == 4 || i == 5 {
						parseTime[i] = currentTime[i]
						continue
					}
				}
			}
			t = time.Date(parseTime[6], time.Month(parseTime[5]), parseTime[4], parseTime[3], parseTime[2], parseTime[1], parseTime[0], location)
			currentTime = FormatTimex(t)
		}
	}
	return
}

// MustParse must parse string to time or it will panic
func (now *Timex) MustParse(s ...string) (t time.Time) {
	t, err := now.Parse(s...)
	if err != nil {
		panic(err)
	}
	return t
}

// Between check time between the begin, end time or not
func (now *Timex) Between(begin, end string) bool {
	beginTime := now.MustParse(begin)
	endTime := now.MustParse(end)
	return now.After(beginTime) && now.Before(endTime)
}

func FormatTimex(t time.Time) []int {
	hour, min, sec := t.Clock()
	year, month, day := t.Date()
	return []int{t.Nanosecond(), sec, min, hour, day, int(month), year}
}
