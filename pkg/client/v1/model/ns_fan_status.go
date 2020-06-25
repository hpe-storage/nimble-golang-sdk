// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsFanStatus Enum.
 
type NsFanStatus string 

const (
	NSFANSTATUS_FAN_FAILED NsFanStatus = "fan_failed"
	NSFANSTATUS_FAN_OKAY NsFanStatus = "fan_okay"
	NSFANSTATUS_FAN_ALERTED NsFanStatus = "fan_alerted"
	NSFANSTATUS_FAN_UNKNOWN NsFanStatus = "fan_unknown"

) 
