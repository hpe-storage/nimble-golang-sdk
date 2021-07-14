// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsShelfSensorState Enum.

type NsShelfSensorState string

const (
 cNsShelfSensorStateMissing NsShelfSensorState = "Missing"
 cNsShelfSensorStateFailed NsShelfSensorState = "Failed"
 cNsShelfSensorStateOk NsShelfSensorState = "OK"
 cNsShelfSensorStateAlerted NsShelfSensorState = "Alerted"
)

var pNsShelfSensorStateMissing NsShelfSensorState
var pNsShelfSensorStateFailed NsShelfSensorState
var pNsShelfSensorStateOk NsShelfSensorState
var pNsShelfSensorStateAlerted NsShelfSensorState

// NsShelfSensorStateMissing enum exports
var NsShelfSensorStateMissing *NsShelfSensorState

// NsShelfSensorStateFailed enum exports
var NsShelfSensorStateFailed *NsShelfSensorState

// NsShelfSensorStateOk enum exports
var NsShelfSensorStateOk *NsShelfSensorState

// NsShelfSensorStateAlerted enum exports
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

