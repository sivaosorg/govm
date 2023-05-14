package utils

import "time"

// SetTimezoneFallback
func SetTimezoneFallback(at time.Time, timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	now := at.In(loc)
	return now, err
}

// SetTimezone
func SetTimezone(at time.Time, timezone string) time.Time {
	t, err := SetTimezoneFallback(at, timezone)
	if err != nil {
		return at
	}
	return t
}
