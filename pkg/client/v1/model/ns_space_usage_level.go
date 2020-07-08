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

// Export
var NsSpaceUsageLevelNormal *NsSpaceUsageLevel
var NsSpaceUsageLevelCritical *NsSpaceUsageLevel
var NsSpaceUsageLevelWarning *NsSpaceUsageLevel

func init() {
	pNsSpaceUsageLevelNormal = cNsSpaceUsageLevelNormal
	NsSpaceUsageLevelNormal = &pNsSpaceUsageLevelNormal

	pNsSpaceUsageLevelCritical = cNsSpaceUsageLevelCritical
	NsSpaceUsageLevelCritical = &pNsSpaceUsageLevelCritical

	pNsSpaceUsageLevelWarning = cNsSpaceUsageLevelWarning
	NsSpaceUsageLevelWarning = &pNsSpaceUsageLevelWarning

}
