// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsAlarmSeverityLevel Enum.

type NsAlarmSeverityLevel string

const (
	cNsAlarmSeverityLevelCritical NsAlarmSeverityLevel = "critical"
	cNsAlarmSeverityLevelWarning  NsAlarmSeverityLevel = "warning"
)

var pNsAlarmSeverityLevelCritical NsAlarmSeverityLevel
var pNsAlarmSeverityLevelWarning NsAlarmSeverityLevel

// Export
var NsAlarmSeverityLevelCritical *NsAlarmSeverityLevel
var NsAlarmSeverityLevelWarning *NsAlarmSeverityLevel

func init() {
	pNsAlarmSeverityLevelCritical = cNsAlarmSeverityLevelCritical
	NsAlarmSeverityLevelCritical = &pNsAlarmSeverityLevelCritical

	pNsAlarmSeverityLevelWarning = cNsAlarmSeverityLevelWarning
	NsAlarmSeverityLevelWarning = &pNsAlarmSeverityLevelWarning

}
