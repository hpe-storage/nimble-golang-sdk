// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsSeverityLevel Enum.

type NsSeverityLevel string

const (
	cNsSeverityLevelCritical NsSeverityLevel = "critical"
	cNsSeverityLevelWarning  NsSeverityLevel = "warning"
	cNsSeverityLevelInfo     NsSeverityLevel = "info"
	cNsSeverityLevelNotice   NsSeverityLevel = "notice"
)

var pNsSeverityLevelCritical NsSeverityLevel
var pNsSeverityLevelWarning NsSeverityLevel
var pNsSeverityLevelInfo NsSeverityLevel
var pNsSeverityLevelNotice NsSeverityLevel

// Export
var NsSeverityLevelCritical *NsSeverityLevel
var NsSeverityLevelWarning *NsSeverityLevel
var NsSeverityLevelInfo *NsSeverityLevel
var NsSeverityLevelNotice *NsSeverityLevel

func init() {
	pNsSeverityLevelCritical = cNsSeverityLevelCritical
	NsSeverityLevelCritical = &pNsSeverityLevelCritical

	pNsSeverityLevelWarning = cNsSeverityLevelWarning
	NsSeverityLevelWarning = &pNsSeverityLevelWarning

	pNsSeverityLevelInfo = cNsSeverityLevelInfo
	NsSeverityLevelInfo = &pNsSeverityLevelInfo

	pNsSeverityLevelNotice = cNsSeverityLevelNotice
	NsSeverityLevelNotice = &pNsSeverityLevelNotice

}
