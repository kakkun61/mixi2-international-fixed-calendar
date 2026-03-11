package international_fixed_calendar

import "time"

type Time time.Time
type Month int
type Day int

const (
	NoMonth Month = iota
	January
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
	LeapDay Day = 29
	YearDay Day = 30
)

func (t Time) Month() Month {
	day := time.Time(t).YearDay()
	if day <= 6*28 {
		return Month((day-1)/28 + 1)
	}
	if isLeap(time.Time(t).Year()) {
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

func (t Time) Day() Day {
	day := time.Time(t).YearDay()
	if day <= 6*28 {
		return Day((day-1)%28 + 1)
	}
	if isLeap(time.Time(t).Year()) {
		if day == 6*28+1 {
			return LeapDay
		}
		day--
	}
	if day == 13*28+1 {
		return YearDay
	}
	return Day((day-1)%28 + 1)
}

// copy from time.isLeap
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
