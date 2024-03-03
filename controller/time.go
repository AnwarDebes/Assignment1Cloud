package controller

import "time"

// uptimeStartTime holds the start time of the application to calculate uptime.
var uptimeStartTime time.Time

// HandleUpTime initializes the uptime tracking by recording the current time.
// This function should be called at the start of the application to begin tracking uptime.
func HandleUpTime() {
	// StartUptime starts recording uptime
	uptimeStartTime = time.Now()
}

// getUptime calculates and returns the duration since the application started, effectively giving the current uptime.
// It uses the uptimeStartTime set by HandleUpTime and calculates the duration from that time to the current time.
func getUptime() time.Duration {
	// GetUptime returns the current uptime duration
	return time.Since(uptimeStartTime)
}
