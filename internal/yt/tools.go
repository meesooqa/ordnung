package yt

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseYtDuration(durationStr string) (time.Duration, error) {
	if !strings.HasPrefix(durationStr, "PT") {
		return 0, fmt.Errorf("invalid duration format, missing PT prefix")
	}
	durationStr = strings.TrimPrefix(durationStr, "PT")

	var total time.Duration

	// hours
	if idx := strings.Index(durationStr, "H"); idx != -1 {
		hoursStr := durationStr[:idx]
		hours, err := strconv.Atoi(hoursStr)
		if err != nil {
			return 0, fmt.Errorf("invalid hours: %v", err)
		}
		total += time.Duration(hours) * time.Hour
		durationStr = durationStr[idx+1:]
	}

	// minutes
	if idx := strings.Index(durationStr, "M"); idx != -1 {
		minutesStr := durationStr[:idx]
		minutes, err := strconv.Atoi(minutesStr)
		if err != nil {
			return 0, fmt.Errorf("invalid minutes: %v", err)
		}
		total += time.Duration(minutes) * time.Minute
		durationStr = durationStr[idx+1:]
	}

	// seconds
	if idx := strings.Index(durationStr, "S"); idx != -1 {
		secondsStr := durationStr[:idx]
		seconds, err := strconv.ParseFloat(secondsStr, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid seconds: %v", err)
		}
		total += time.Duration(seconds * float64(time.Second))
	}

	return total, nil
}
