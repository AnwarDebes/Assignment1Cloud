package controller

import "time"

var uptimeStartTime time.Time

func HandleUpTime() {
	// StartUptime starts recording uptime
	uptimeStartTime = time.Now()
}
func getUptime() time.Duration {
	// GetUptime returns the current uptime duration
	return time.Since(uptimeStartTime)
}
