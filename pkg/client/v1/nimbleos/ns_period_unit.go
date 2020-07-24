// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsPeriodUnit Enum.

type NsPeriodUnit string

const (
	cNsPeriodUnitHours   NsPeriodUnit = "hours"
	cNsPeriodUnitWeeks   NsPeriodUnit = "weeks"
	cNsPeriodUnitMinutes NsPeriodUnit = "minutes"
	cNsPeriodUnitDays    NsPeriodUnit = "days"
)

var pNsPeriodUnitHours NsPeriodUnit
var pNsPeriodUnitWeeks NsPeriodUnit
var pNsPeriodUnitMinutes NsPeriodUnit
var pNsPeriodUnitDays NsPeriodUnit

// NsPeriodUnitHours enum exports
var NsPeriodUnitHours *NsPeriodUnit

// NsPeriodUnitWeeks enum exports
var NsPeriodUnitWeeks *NsPeriodUnit

// NsPeriodUnitMinutes enum exports
var NsPeriodUnitMinutes *NsPeriodUnit

// NsPeriodUnitDays enum exports
var NsPeriodUnitDays *NsPeriodUnit

func init() {
	pNsPeriodUnitHours = cNsPeriodUnitHours
	NsPeriodUnitHours = &pNsPeriodUnitHours

	pNsPeriodUnitWeeks = cNsPeriodUnitWeeks
	NsPeriodUnitWeeks = &pNsPeriodUnitWeeks

	pNsPeriodUnitMinutes = cNsPeriodUnitMinutes
	NsPeriodUnitMinutes = &pNsPeriodUnitMinutes

	pNsPeriodUnitDays = cNsPeriodUnitDays
	NsPeriodUnitDays = &pNsPeriodUnitDays

}
