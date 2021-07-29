// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsAlarmSeverityLevel Enum.

type NsAlarmSeverityLevel string

const (
	cNsAlarmSeverityLevelCritical NsAlarmSeverityLevel = "critical"
	cNsAlarmSeverityLevelWarning  NsAlarmSeverityLevel = "warning"
)

var pNsAlarmSeverityLevelCritical NsAlarmSeverityLevel
var pNsAlarmSeverityLevelWarning NsAlarmSeverityLevel

// NsAlarmSeverityLevelCritical enum exports
var NsAlarmSeverityLevelCritical *NsAlarmSeverityLevel

// NsAlarmSeverityLevelWarning enum exports
var NsAlarmSeverityLevelWarning *NsAlarmSeverityLevel

func init() {
	pNsAlarmSeverityLevelCritical = cNsAlarmSeverityLevelCritical
	NsAlarmSeverityLevelCritical = &pNsAlarmSeverityLevelCritical

	pNsAlarmSeverityLevelWarning = cNsAlarmSeverityLevelWarning
	NsAlarmSeverityLevelWarning = &pNsAlarmSeverityLevelWarning

}
