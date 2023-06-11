package timex

import "time"

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
