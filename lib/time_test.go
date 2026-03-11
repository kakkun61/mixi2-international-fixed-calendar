package international_fixed_calendar

import (
	"testing"
	"time"
)

func TestTimeMonth(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want Month
	}{
		{
			name: "regular day in first month",
			date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: January,
		},
		{
			name: "leap day has no month",
			date: time.Date(2024, time.June, 17, 0, 0, 0, 0, time.UTC),
			want: NoMonth,
		},
		{
			name: "day after leap day starts sol in leap year",
			date: time.Date(2024, time.June, 18, 0, 0, 0, 0, time.UTC),
			want: Sol,
		},
		{
			name: "last regular month in leap year remains december",
			date: time.Date(2024, time.December, 30, 0, 0, 0, 0, time.UTC),
			want: December,
		},
		{
			name: "year day in leap year has no month",
			date: time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC),
			want: NoMonth,
		},
		{
			name: "year day has no month",
			date: time.Date(2025, time.December, 31, 0, 0, 0, 0, time.UTC),
			want: NoMonth,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Time{gregorian: tt.date}.Month()
			if got != tt.want {
				t.Fatalf("Month() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeDay(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want int
	}{
		{
			name: "regular day in first month",
			date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "leap day is returned explicitly",
			date: time.Date(2024, time.June, 17, 0, 0, 0, 0, time.UTC),
			want: LeapDay,
		},
		{
			name: "day after leap day resets to first day of sol",
			date: time.Date(2024, time.June, 18, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "last regular day before year day in leap year",
			date: time.Date(2024, time.December, 30, 0, 0, 0, 0, time.UTC),
			want: 28,
		},
		{
			name: "year day in leap year is returned explicitly",
			date: time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC),
			want: YearDay,
		},
		{
			name: "year day is returned explicitly",
			date: time.Date(2025, time.December, 31, 0, 0, 0, 0, time.UTC),
			want: YearDay,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Time{gregorian: tt.date}.Day()
			if got != tt.want {
				t.Fatalf("Day() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeWeekday(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want time.Weekday
	}{
		{
			name: "regular day in first month",
			date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: time.Sunday,
		},
		{
			name: "leap day has no weekday",
			date: time.Date(2024, time.June, 17, 0, 0, 0, 0, time.UTC),
			want: NoWeekday,
		},
		{
			name: "day after leap day starts sol and is a weekday",
			date: time.Date(2024, time.June, 18, 0, 0, 0, 0, time.UTC),
			want: time.Sunday,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Time{gregorian: tt.date}.Weekday()
			if got != tt.want {
				t.Fatalf("Weekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeDate(t *testing.T) {
	tests := []struct {
		name  string
		year  int
		month Month
		day   int
	}{
		{
			name:  "regular day in first month",
			year:  2025,
			month: January,
			day:   1,
		},
		{
			name:  "leap day has no month",
			year:  2024,
			month: NoMonth,
			day:   LeapDay,
		},
		{
			name:  "day after leap day starts sol",
			year:  2024,
			month: Sol,
			day:   1,
		},
		{
			name:  "last regular day before year day in leap year",
			year:  2024,
			month: December,
			day:   28,
		},
		{
			name:  "year day in leap year has no month",
			year:  2024,
			month: NoMonth,
			day:   YearDay,
		},
		{
			name:  "year day has no month",
			year:  2025,
			month: NoMonth,
			day:   YearDay,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Date(tt.year, tt.month, tt.day, 0, 0, 0, 0, time.UTC)
			if got.Year() != tt.year {
				t.Fatalf("Year() = %v, want %v", got.Year(), tt.year)
			}
			if got.Month() != tt.month {
				t.Fatalf("Month() = %v, want %v", got.Month(), tt.month)
			}
			if got.Day() != tt.day {
				t.Fatalf("Day() = %v, want %v", got.Day(), tt.day)
			}
		})
	}
}

func TestIsLeap(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		{
			name: "common year",
			year: 2025,
			want: false,
		},
		{
			name: "typical leap year",
			year: 2024,
			want: true,
		},
		{
			name: "century non leap year",
			year: 1900,
			want: false,
		},
		{
			name: "400 year leap year",
			year: 2000,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isLeap(tt.year)
			if got != tt.want {
				t.Fatalf("isLeap(%d) = %v, want %v", tt.year, got, tt.want)
			}
		})
	}
}
