// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsSensorState Enum.

type NsSensorState string

const (
 cNsSensorStateSensorOk NsSensorState = "sensor_ok"
 cNsSensorStateSensorAlertCond NsSensorState = "sensor_alert_cond"
 cNsSensorStateSensorMissing NsSensorState = "sensor_missing"
 cNsSensorStateSensorReadingUnavail NsSensorState = "sensor_reading_unavail"
 cNsSensorStateSensorFailed NsSensorState = "sensor_failed"
)

var pNsSensorStateSensorOk NsSensorState
var pNsSensorStateSensorAlertCond NsSensorState
var pNsSensorStateSensorMissing NsSensorState
var pNsSensorStateSensorReadingUnavail NsSensorState
var pNsSensorStateSensorFailed NsSensorState

// NsSensorStateSensorOk enum exports
var NsSensorStateSensorOk *NsSensorState

// NsSensorStateSensorAlertCond enum exports
var NsSensorStateSensorAlertCond *NsSensorState

// NsSensorStateSensorMissing enum exports
var NsSensorStateSensorMissing *NsSensorState

// NsSensorStateSensorReadingUnavail enum exports
var NsSensorStateSensorReadingUnavail *NsSensorState

// NsSensorStateSensorFailed enum exports
var NsSensorStateSensorFailed *NsSensorState

func init() {
 pNsSensorStateSensorOk = cNsSensorStateSensorOk
 NsSensorStateSensorOk = &pNsSensorStateSensorOk

 pNsSensorStateSensorAlertCond = cNsSensorStateSensorAlertCond
 NsSensorStateSensorAlertCond = &pNsSensorStateSensorAlertCond

 pNsSensorStateSensorMissing = cNsSensorStateSensorMissing
 NsSensorStateSensorMissing = &pNsSensorStateSensorMissing

 pNsSensorStateSensorReadingUnavail = cNsSensorStateSensorReadingUnavail
 NsSensorStateSensorReadingUnavail = &pNsSensorStateSensorReadingUnavail

 pNsSensorStateSensorFailed = cNsSensorStateSensorFailed
 NsSensorStateSensorFailed = &pNsSensorStateSensorFailed

}

