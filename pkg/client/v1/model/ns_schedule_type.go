// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsScheduleType Enum.

type NsScheduleType string

const (
	cNsScheduleTypeExternalTrigger NsScheduleType = "external_trigger"
	cNsScheduleTypeRegular         NsScheduleType = "regular"
)

var pNsScheduleTypeExternalTrigger NsScheduleType
var pNsScheduleTypeRegular NsScheduleType

// Export
var NsScheduleTypeExternalTrigger *NsScheduleType
var NsScheduleTypeRegular *NsScheduleType

func init() {
	pNsScheduleTypeExternalTrigger = cNsScheduleTypeExternalTrigger
	NsScheduleTypeExternalTrigger = &pNsScheduleTypeExternalTrigger

	pNsScheduleTypeRegular = cNsScheduleTypeRegular
	NsScheduleTypeRegular = &pNsScheduleTypeRegular

}
