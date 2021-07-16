// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsFanStatus Enum.

type NsFanStatus string

const (
	cNsFanStatusFanFailed  NsFanStatus = "fan_failed"
	cNsFanStatusFanOkay    NsFanStatus = "fan_okay"
	cNsFanStatusFanAlerted NsFanStatus = "fan_alerted"
	cNsFanStatusFanUnknown NsFanStatus = "fan_unknown"
)

var pNsFanStatusFanFailed NsFanStatus
var pNsFanStatusFanOkay NsFanStatus
var pNsFanStatusFanAlerted NsFanStatus
var pNsFanStatusFanUnknown NsFanStatus

// NsFanStatusFanFailed enum exports
var NsFanStatusFanFailed *NsFanStatus

// NsFanStatusFanOkay enum exports
var NsFanStatusFanOkay *NsFanStatus

// NsFanStatusFanAlerted enum exports
var NsFanStatusFanAlerted *NsFanStatus

// NsFanStatusFanUnknown enum exports
var NsFanStatusFanUnknown *NsFanStatus

func init() {
	pNsFanStatusFanFailed = cNsFanStatusFanFailed
	NsFanStatusFanFailed = &pNsFanStatusFanFailed

	pNsFanStatusFanOkay = cNsFanStatusFanOkay
	NsFanStatusFanOkay = &pNsFanStatusFanOkay

	pNsFanStatusFanAlerted = cNsFanStatusFanAlerted
	NsFanStatusFanAlerted = &pNsFanStatusFanAlerted

	pNsFanStatusFanUnknown = cNsFanStatusFanUnknown
	NsFanStatusFanUnknown = &pNsFanStatusFanUnknown

}
