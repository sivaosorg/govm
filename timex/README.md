## Timex

## Usage

Calculating time based on current time

```go

time.Now() // 2013-11-18 17:51:49.123456789 Mon

timex.BeginningOfMinute()        // 2013-11-18 17:51:00 Mon
timex.BeginningOfHour()          // 2013-11-18 17:00:00 Mon
timex.BeginningOfDay()           // 2013-11-18 00:00:00 Mon
timex.BeginningOfWeek()          // 2013-11-17 00:00:00 Sun
timex.BeginningOfMonth()         // 2013-11-01 00:00:00 Fri
timex.BeginningOfQuarter()       // 2013-10-01 00:00:00 Tue
timex.BeginningOfYear()          // 2013-01-01 00:00:00 Tue

timex.EndOfMinute()              // 2013-11-18 17:51:59.999999999 Mon
timex.EndOfHour()                // 2013-11-18 17:59:59.999999999 Mon
timex.EndOfDay()                 // 2013-11-18 23:59:59.999999999 Mon
timex.EndOfWeek()                // 2013-11-23 23:59:59.999999999 Sat
timex.EndOfMonth()               // 2013-11-30 23:59:59.999999999 Sat
timex.EndOfQuarter()             // 2013-12-31 23:59:59.999999999 Tue
timex.EndOfYear()                // 2013-12-31 23:59:59.999999999 Tue

timex.WeekStartDay = time.Monday // Set Monday as first day, default is Sunday
timex.EndOfWeek()                // 2013-11-24 23:59:59.999999999 Sun
```

Calculating time based on another time

```go
t := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.timex().Location())
timex.With(t).EndOfMonth()   // 2013-02-28 23:59:59.999999999 Thu
```

Calculating time based on configuration

```go
location, err := time.LoadLocation("Asia/Ho_Chi_Minh")

myConfig := &timex.Config{
	WeekStartDay: time.Monday,
	TimeLocation: location,
	TimeFormats: []string{"2006-01-02 15:04:05"},
}

t := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.Now().Location()) // // 2013-11-18 17:51:49.123456789 Mon
myConfig.With(t).BeginningOfWeek()         // 2013-11-18 00:00:00 Mon

myConfig.Parse("2002-10-12 22:14:01")     // 2002-10-12 22:14:01
myConfig.Parse("2002-10-12 22:14")        // returns error 'can't parse string as time: 2002-10-12 22:14'
```

### Monday/Sunday

Don't be bothered with the `WeekStartDay` setting, you can use `Monday`, `Sunday`

```go
timex.Monday()              // 2013-11-18 00:00:00 Mon
timex.Monday("17:44")       // 2013-11-18 17:44:00 Mon
timex.Sunday()              // 2013-11-24 00:00:00 Sun (Next Sunday)
timex.Sunday("18:19:24")    // 2013-11-24 18:19:24 Sun (Next Sunday)
timex.EndOfSunday()         // 2013-11-24 23:59:59.999999999 Sun (End of next Sunday)

t := time.Date(2013, 11, 24, 17, 51, 49, 123456789, time.Now().Location()) // 2013-11-24 17:51:49.123456789 Sun
timex.With(t).Monday()              // 2013-11-18 00:00:00 Mon (Last Monday if today is Sunday)
timex.With(t).Monday("17:44")       // 2013-11-18 17:44:00 Mon (Last Monday if today is Sunday)
timex.With(t).Sunday()              // 2013-11-24 00:00:00 Sun (Beginning Of Today if today is Sunday)
timex.With(t).Sunday("18:19:24")    // 2013-11-24 18:19:24 Sun (Beginning Of Today if today is Sunday)
timex.With(t).EndOfSunday()         // 2013-11-24 23:59:59.999999999 Sun (End of Today if today is Sunday)
```

### Parse String to Time

```go
time.Now() // 2013-11-18 17:51:49.123456789 Mon

// Parse(string) (time.Time, error)
t, err := timex.Parse("2017")                // 2017-01-01 00:00:00, nil
t, err := timex.Parse("2017-10")             // 2017-10-01 00:00:00, nil
t, err := timex.Parse("2017-10-13")          // 2017-10-13 00:00:00, nil
t, err := timex.Parse("1999-12-12 12")       // 1999-12-12 12:00:00, nil
t, err := timex.Parse("1999-12-12 12:20")    // 1999-12-12 12:20:00, nil
t, err := timex.Parse("1999-12-12 12:20:21") // 1999-12-12 12:20:21, nil
t, err := timex.Parse("10-13")               // 2013-10-13 00:00:00, nil
t, err := timex.Parse("12:20")               // 2013-11-18 12:20:00, nil
t, err := timex.Parse("12:20:13")            // 2013-11-18 12:20:13, nil
t, err := timex.Parse("14")                  // 2013-11-18 14:00:00, nil
t, err := timex.Parse("99:99")               // 2013-11-18 12:20:00, Can't parse string as time: 99:99

// MustParse must parse string to time or it will panic
timex.MustParse("2013-01-13")             // 2013-01-13 00:00:00
timex.MustParse("02-17")                  // 2013-02-17 00:00:00
timex.MustParse("2-17")                   // 2013-02-17 00:00:00
timex.MustParse("8")                      // 2013-11-18 08:00:00
timex.MustParse("2002-10-12 22:14")       // 2002-10-12 22:14:00
timex.MustParse("99:99")                  // panic: Can't parse string as time: 99:99
```

Extend `timex` to support more formats is quite easy, just update `timex.TimeFormats` with other time layouts, e.g:

```go
timex.TimeFormats = append(timex.TimeFormats, "02 Jan 2006 15:04")
```
