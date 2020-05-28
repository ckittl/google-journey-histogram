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
		return WeekEnd
	} else {
		return WorkingDay
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

// Let's the routine wait for the next full duration
func WaitToNextFull(dur gotime.Duration) {
	executionStart := gotime.Now().Add(dur).Truncate(dur)
	println("Wait until " + executionStart.Format("2006-01-02 15:04:05"))
	gotime.Sleep(gotime.Until(executionStart))
}

// Build one ticker, that gives the frequency and one that gives the end of the frequency ticks
func GetTickers(frequency gotime.Duration, runDuration gotime.Duration) (*gotime.Ticker, *gotime.Timer) {
	frequencyTicker := gotime.NewTicker(frequency)
	endTimer := gotime.NewTimer(runDuration)
	return frequencyTicker, endTimer
}
