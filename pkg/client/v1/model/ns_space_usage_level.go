// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsSpaceUsageLevel Enum.

type NsSpaceUsageLevel string

const (
	cNsSpaceUsageLevelNormal   NsSpaceUsageLevel = "normal"
	cNsSpaceUsageLevelCritical NsSpaceUsageLevel = "critical"
	cNsSpaceUsageLevelWarning  NsSpaceUsageLevel = "warning"
)

var pNsSpaceUsageLevelNormal NsSpaceUsageLevel
var pNsSpaceUsageLevelCritical NsSpaceUsageLevel
var pNsSpaceUsageLevelWarning NsSpaceUsageLevel

// NsSpaceUsageLevelNormal enum exports
var NsSpaceUsageLevelNormal *NsSpaceUsageLevel

// NsSpaceUsageLevelCritical enum exports
var NsSpaceUsageLevelCritical *NsSpaceUsageLevel

// NsSpaceUsageLevelWarning enum exports
var NsSpaceUsageLevelWarning *NsSpaceUsageLevel

func init() {
	pNsSpaceUsageLevelNormal = cNsSpaceUsageLevelNormal
	NsSpaceUsageLevelNormal = &pNsSpaceUsageLevelNormal

	pNsSpaceUsageLevelCritical = cNsSpaceUsageLevelCritical
	NsSpaceUsageLevelCritical = &pNsSpaceUsageLevelCritical

	pNsSpaceUsageLevelWarning = cNsSpaceUsageLevelWarning
	NsSpaceUsageLevelWarning = &pNsSpaceUsageLevelWarning

}
