package time

import (
	"regexp"
	"strconv"
	gotime "time"
)

type DayType int

const (
	WorkingDay DayType = iota
	WeekEnd
)

// Determine, whether its a working day or a weekend
func GetDayType(time gotime.Time) DayType {
	weekDay := time.Weekday()
	if weekDay == gotime.Sunday || weekDay == gotime.Saturday {
		return WorkingDay
	} else {
		return WeekEnd
	}
}

func (d DayType) String() string {
	if d == WorkingDay {
		return "WorkingDay"
	} else {
		return "WeekEnd"
	}
}

// Calculates the travel duration in minutes based on a string of the format 'X hours Y minutes'. The hours part is
// optional
func CalcDurationInMinutes(duration string) int {
	r, _ := regexp.Compile("(?:(\\d+) hours )?(\\d+) mins")
	submatches := r.FindSubmatch([]byte(duration))
	hours, _ := strconv.Atoi(string(submatches[1]))
	minutes, _ := strconv.Atoi(string(submatches[2]))
	return hours*60 + minutes
}
