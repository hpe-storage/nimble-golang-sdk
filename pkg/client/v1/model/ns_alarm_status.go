// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsAlarmStatus Enum.

type NsAlarmStatus string

const (
	cNsAlarmStatusAcknowledged NsAlarmStatus = "acknowledged"
	cNsAlarmStatusOpen         NsAlarmStatus = "open"
)

var pNsAlarmStatusAcknowledged NsAlarmStatus
var pNsAlarmStatusOpen NsAlarmStatus

// Export
var NsAlarmStatusAcknowledged *NsAlarmStatus
var NsAlarmStatusOpen *NsAlarmStatus

func init() {
	pNsAlarmStatusAcknowledged = cNsAlarmStatusAcknowledged
	NsAlarmStatusAcknowledged = &pNsAlarmStatusAcknowledged

	pNsAlarmStatusOpen = cNsAlarmStatusOpen
	NsAlarmStatusOpen = &pNsAlarmStatusOpen

}
