// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsShelfSensorState Enum.

type NsShelfSensorState string

const (
	cNsShelfSensorStateMissing NsShelfSensorState = "Missing"
	cNsShelfSensorStateFailed  NsShelfSensorState = "Failed"
	cNsShelfSensorStateOk      NsShelfSensorState = "OK"
	cNsShelfSensorStateAlerted NsShelfSensorState = "Alerted"
)

var pNsShelfSensorStateMissing NsShelfSensorState
var pNsShelfSensorStateFailed NsShelfSensorState
var pNsShelfSensorStateOk NsShelfSensorState
var pNsShelfSensorStateAlerted NsShelfSensorState

// Export
var NsShelfSensorStateMissing *NsShelfSensorState
var NsShelfSensorStateFailed *NsShelfSensorState
var NsShelfSensorStateOk *NsShelfSensorState
var NsShelfSensorStateAlerted *NsShelfSensorState

func init() {
	pNsShelfSensorStateMissing = cNsShelfSensorStateMissing
	NsShelfSensorStateMissing = &pNsShelfSensorStateMissing

	pNsShelfSensorStateFailed = cNsShelfSensorStateFailed
	NsShelfSensorStateFailed = &pNsShelfSensorStateFailed

	pNsShelfSensorStateOk = cNsShelfSensorStateOk
	NsShelfSensorStateOk = &pNsShelfSensorStateOk

	pNsShelfSensorStateAlerted = cNsShelfSensorStateAlerted
	NsShelfSensorStateAlerted = &pNsShelfSensorStateAlerted

}
