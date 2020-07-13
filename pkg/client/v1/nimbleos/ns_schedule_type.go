// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsScheduleType Enum.

type NsScheduleType string

const (
	cNsScheduleTypeExternalTrigger NsScheduleType = "external_trigger"
	cNsScheduleTypeRegular         NsScheduleType = "regular"
)

var pNsScheduleTypeExternalTrigger NsScheduleType
var pNsScheduleTypeRegular NsScheduleType

// NsScheduleTypeExternalTrigger enum exports
var NsScheduleTypeExternalTrigger *NsScheduleType

// NsScheduleTypeRegular enum exports
var NsScheduleTypeRegular *NsScheduleType

func init() {
	pNsScheduleTypeExternalTrigger = cNsScheduleTypeExternalTrigger
	NsScheduleTypeExternalTrigger = &pNsScheduleTypeExternalTrigger

	pNsScheduleTypeRegular = cNsScheduleTypeRegular
	NsScheduleTypeRegular = &pNsScheduleTypeRegular

}
