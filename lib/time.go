package international_fixed_calendar

import "time"

type Time struct {
	gregorian time.Time
}

type Month int

const (
	NoMonth Month = -1
	January Month = iota
	February
	March
	April
	May
	June
	Sol
	July
	August
	September
	October
	November
	December
)

const (
	LeapDay int = -1
	YearDay int = -2
)

const (
	NoWeekday = -1
)

func FromGregorian(t time.Time) Time {
	return Time{gregorian: t}
}

func ToGregorian(t Time) time.Time {
	return t.gregorian
}

func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	if loc == nil {
		panic("time: missing Location in call to Date")
	}

	// Normalize month, overflowing into year.
	if month != NoMonth {
		m := int(month) - 1
		year, m = norm(year, m, 13)
		month = Month(m) + 1
	}

	if day == LeapDay {
		return Time{gregorian: time.Date(year, time.June, 17, hour, min, sec, nsec, loc)}
	}
	if day == YearDay {
		return Time{gregorian: time.Date(year, time.December, 31, hour, min, sec, nsec, loc)}
	}
	gDay := (int(month)-1)*28 + day // initially year day
	if month >= Sol {
		gDay = gDay + 1 // add leap day
	}

	if gDay <= 31 {
		return Time{gregorian: time.Date(year, time.January, gDay, hour, min, sec, nsec, loc)}
	}
	gDay = gDay - 31
	februaryDays := 28
	if isLeap(year) {
		februaryDays = 29
	}
	if gDay <= februaryDays {
		return Time{gregorian: time.Date(year, time.February, gDay, hour, min, sec, nsec, loc)}
	}
	gDay = gDay - februaryDays
	if gDay <= 31 {
		return Time{gregorian: time.Date(year, time.March, gDay, hour, min, sec, nsec, loc)}
	}
	gDay = gDay - 31
	if gDay <= 30 {
		return Time{gregorian: time.Date(year, time.April, gDay, hour, min, sec, nsec, loc)}
	}
	gDay = gDay - 30
	if gDay <= 31 {
		return Time{gregorian: time.Date(year, time.May, gDay, hour, min, sec, nsec, loc)}
	}
	gDay = gDay - 31
	if gDay <= 30 {
		return Time{gregorian: time.Date(year, time.June, gDay, hour, min, sec, nsec, loc)}
	}
	gDay = gDay - 30
	if gDay <= 31 {
		return Time{gregorian: time.Date(year, time.July, gDay, hour, min, sec, nsec, loc)}
	}
	gDay = gDay - 31
	if gDay <= 31 {
		return Time{gregorian: time.Date(year, time.August, gDay, hour, min, sec, nsec, loc)}
	}
	gDay = gDay - 31
	if gDay <= 30 {
		return Time{gregorian: time.Date(year, time.September, gDay, hour, min, sec, nsec, loc)}
	}
	gDay = gDay - 30
	if gDay <= 31 {
		return Time{gregorian: time.Date(year, time.October, gDay, hour, min, sec, nsec, loc)}
	}
	gDay = gDay - 31
	if gDay <= 30 {
		return Time{gregorian: time.Date(year, time.November, gDay, hour, min, sec, nsec, loc)}
	}
	return Time{gregorian: time.Date(year, time.December, gDay-30, hour, min, sec, nsec, loc)}
}

func (t Time) Year() int {
	return t.gregorian.Year()
}

func (t Time) Month() Month {
	day := t.gregorian.YearDay()
	if day <= 6*28 {
		return Month((day-1)/28 + 1)
	}
	if isLeap(t.gregorian.Year()) {
		if day == 6*28+1 {
			return NoMonth
		}
		day--
	}
	if day == 13*28+1 {
		return NoMonth
	}
	return Month((day-1)/28 + 1)
}

func (t Time) Day() int {
	day := t.gregorian.YearDay()
	if day <= 6*28 {
		return (day-1)%28 + 1
	}
	if isLeap(t.gregorian.Year()) {
		if day == 6*28+1 {
			return LeapDay
		}
		day--
	}
	if day == 13*28+1 {
		return YearDay
	}
	return (day-1)%28 + 1
}

func (t Time) Weekday() time.Weekday {
	day := t.Day()
	if day == LeapDay || day == YearDay {
		return NoWeekday
	}
	return time.Weekday((day - 1) % 7)
}

// copy of time.isLeap
func isLeap(year int) bool {
	// year%4 == 0 && (year%100 != 0 || year%400 == 0)
	// Bottom 2 bits must be clear.
	// For multiples of 25, bottom 4 bits must be clear.
	mask := 0xf
	if year%25 != 0 {
		mask = 3
	}
	return year&mask == 0
}

// copy of time.norm
func norm(hi, lo, base int) (nhi, nlo int) {
	if lo < 0 {
		n := (-lo-1)/base + 1
		hi -= n
		lo += n * base
	}
	if lo >= base {
		n := lo / base
		hi += n
		lo -= n * base
	}
	return hi, lo
}
