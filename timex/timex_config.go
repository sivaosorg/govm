package timex

import "regexp"

// Deprecated: unsupported
const (
	DateTimeFormatYearMonthDTimezoneHourMinuteSecond  = "2006-01-02T15:04:05"
	DateTimeFormatYearMonthDay                        = "2006-01-02"
	DateTimeFormYearMonthDayHourMinuteSecond          = "2006-01-02 15:04:05"
	DateTimeFormRFC3339NanoTimezone                   = "2006-01-02 15:04:05.999999999 -07:00"
	DateTimeFormDotYearMonthDayHourMinuteSecond       = "2006.01.02.15.04.05"
	DateTimeFormCommaDashYearMonthDayHourMinuteSecond = "2006:01:02_15:04:05"
	DateTimeForm20060102150405999999                  = "2006-01-02 15:04:05.999999"
)

// Recommended
const (
	TimeRFC01T150405 = "15:04:05"
	TimeRFC02D150405 = "15.04.05"
)

// Recommended
const (
	TimeFormat20060102T150405999999       = "2006-01-02T15:04:05.999999"
	TimeFormat20060102T150405             = "2006-01-02T15:04:05"
	TimeFormat20060102150405              = "2006-01-02 15:04:05"
	TimeFormat02012006150405              = "02-01-2006 15:04:05"
	TimeFormatRFC0102012006150405         = "02/01/2006 15:04:05"
	TimeFormat20060102                    = "2006-01-02"
	TimeFormatRFC0102012006               = "02/01/2006"
	TimeFormat20060102150405999999        = "2006-01-02 15:04:05.999999"
	TimeFormat20060102150405999999RFC3339 = "2006-01-02 15:04:05.999999999 -07:00"
)

var ApplyTimeRegexp = regexp.MustCompile(`(\s+|^\s*|T)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))(\s*$|[Z+-])`) // match 15:04:05, 15:04:05.000, 15:04:05.000000 15, 2017-01-01 15:04, 2021-07-20T00:59:10Z, 2021-07-20T00:59:10+08:00, 2021-07-20T00:00:10-07:00 etc
var OnlyTimeRegexp = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)                  // match 15:04:05, 15, 15:04:05.000, 15:04:05.000000, etc

// Timezone constants representing default timezones for specific regions.
const (
	// DefaultTimezoneVietnam is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Vietnam, which is "Asia/Ho_Chi_Minh".
	DefaultTimezoneVietnam = "Asia/Ho_Chi_Minh"

	// DefaultTimezoneNewYork is a constant that holds the IANA Time Zone identifier
	// for the default timezone in New York, USA, which is "America/New_York".
	DefaultTimezoneNewYork = "America/New_York"

	// DefaultTimezoneLondon is a constant that holds the IANA Time Zone identifier
	// for the default timezone in London, United Kingdom, which is "Europe/London".
	DefaultTimezoneLondon = "Europe/London"

	// DefaultTimezoneTokyo is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Tokyo, Japan, which is "Asia/Tokyo".
	DefaultTimezoneTokyo = "Asia/Tokyo"

	// DefaultTimezoneSydney is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Sydney, Australia, which is "Australia/Sydney".
	DefaultTimezoneSydney = "Australia/Sydney"

	// DefaultTimezoneParis is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Paris, France, which is "Europe/Paris".
	DefaultTimezoneParis = "Europe/Paris"

	// DefaultTimezoneMoscow is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Moscow, Russia, which is "Europe/Moscow".
	DefaultTimezoneMoscow = "Europe/Moscow"

	// DefaultTimezoneLosAngeles is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Los Angeles, USA, which is "America/Los_Angeles".
	DefaultTimezoneLosAngeles = "America/Los_Angeles"

	// DefaultTimezoneManila is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Manila, Philippines, which is "Asia/Manila".
	DefaultTimezoneManila = "Asia/Manila"

	// DefaultTimezoneKualaLumpur is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Kuala Lumpur, Malaysia, which is "Asia/Kuala_Lumpur".
	DefaultTimezoneKualaLumpur = "Asia/Kuala_Lumpur"

	// DefaultTimezoneJakarta is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Jakarta, Indonesia, which is "Asia/Jakarta".
	DefaultTimezoneJakarta = "Asia/Jakarta"

	// DefaultTimezoneYangon is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Yangon, Myanmar, which is "Asia/Yangon".
	DefaultTimezoneYangon = "Asia/Yangon"

	// DefaultTimezoneAuckland is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Auckland, New Zealand, which is "Pacific/Auckland".
	DefaultTimezoneAuckland = "Pacific/Auckland"

	// DefaultTimezoneBangkok is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Bangkok, Thailand, which is "Asia/Bangkok".
	DefaultTimezoneBangkok = "Asia/Bangkok"

	// DefaultTimezoneDelhi is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Delhi, India, which is "Asia/Kolkata".
	DefaultTimezoneDelhi = "Asia/Kolkata"

	// DefaultTimezoneDubai is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Dubai, United Arab Emirates, which is "Asia/Dubai".
	DefaultTimezoneDubai = "Asia/Dubai"

	// DefaultTimezoneCairo is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Cairo, Egypt, which is "Africa/Cairo".
	DefaultTimezoneCairo = "Africa/Cairo"

	// DefaultTimezoneAthens is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Athens, Greece, which is "Europe/Athens".
	DefaultTimezoneAthens = "Europe/Athens"

	// DefaultTimezoneRome is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Rome, Italy, which is "Europe/Rome".
	DefaultTimezoneRome = "Europe/Rome"

	// DefaultTimezoneJohannesburg is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Johannesburg, South Africa, which is "Africa/Johannesburg".
	DefaultTimezoneJohannesburg = "Africa/Johannesburg"

	// DefaultTimezoneStockholm is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Stockholm, Sweden, which is "Europe/Stockholm".
	DefaultTimezoneStockholm = "Europe/Stockholm"

	// DefaultTimezoneOslo is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Oslo, Norway, which is "Europe/Oslo".
	DefaultTimezoneOslo = "Europe/Oslo"

	// DefaultTimezoneHelsinki is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Helsinki, Finland, which is "Europe/Helsinki".
	DefaultTimezoneHelsinki = "Europe/Helsinki"

	// DefaultTimezoneKiev is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Kiev, Ukraine, which is "Europe/Kiev".
	DefaultTimezoneKiev = "Europe/Kiev"

	// DefaultTimezoneBeijing is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Beijing, China, which is "Asia/Shanghai".
	DefaultTimezoneBeijing = "Asia/Shanghai"

	// DefaultTimezoneSingapore is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Singapore, which is "Asia/Singapore".
	DefaultTimezoneSingapore = "Asia/Singapore"

	// DefaultTimezoneIslamabad is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Islamabad, Pakistan, which is "Asia/Karachi".
	DefaultTimezoneIslamabad = "Asia/Karachi"

	// DefaultTimezoneColombo is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Colombo, Sri Lanka, which is "Asia/Colombo".
	DefaultTimezoneColombo = "Asia/Colombo"

	// DefaultTimezoneDhaka is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Dhaka, Bangladesh, which is "Asia/Dhaka".
	DefaultTimezoneDhaka = "Asia/Dhaka"

	// DefaultTimezoneKathmandu is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Kathmandu, Nepal, which is "Asia/Kathmandu".
	DefaultTimezoneKathmandu = "Asia/Kathmandu"

	// DefaultTimezoneBrisbane is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Brisbane, Australia, which is "Australia/Brisbane".
	DefaultTimezoneBrisbane = "Australia/Brisbane"

	// DefaultTimezoneWellington is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Wellington, New Zealand, which is "Pacific/Auckland".
	DefaultTimezoneWellington = "Pacific/Auckland"

	// DefaultTimezonePortMoresby is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Port Moresby, Papua New Guinea, which is "Pacific/Port_Moresby".
	DefaultTimezonePortMoresby = "Pacific/Port_Moresby"

	// DefaultTimezoneSuva is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Suva, Fiji, which is "Pacific/Fiji".
	DefaultTimezoneSuva = "Pacific/Fiji"
)
