// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsPeriodUnitHours *NsPeriodUnit
var NsPeriodUnitWeeks *NsPeriodUnit
var NsPeriodUnitMinutes *NsPeriodUnit
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
