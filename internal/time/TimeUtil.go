package time

import (
	"regexp"
	"strconv"
)

// Calculates the travel duration in minutes based on a string of the format 'X hours Y minutes'. The hours part is
// optional
func CalcDurationInMinutes(duration string) int {
	r, _ := regexp.Compile("(?:(\\d+) hours )?(\\d+) mins")
	submatches := r.FindSubmatch([]byte(duration))
	hours, _ := strconv.Atoi(string(submatches[1]))
	minutes, _ := strconv.Atoi(string(submatches[2]))
	return hours*60 + minutes
}
