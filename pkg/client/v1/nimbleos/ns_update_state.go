// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsUpdateState Enum.

type NsUpdateState string

const (
	cNsUpdateStateNormal   NsUpdateState = "normal"
	cNsUpdateStatePaused   NsUpdateState = "paused"
	cNsUpdateStateUpdating NsUpdateState = "updating"
	cNsUpdateStateInvalid  NsUpdateState = "invalid"
	cNsUpdateStateTimedOut NsUpdateState = "timed_out"
	cNsUpdateStateFailed   NsUpdateState = "failed"
)

var pNsUpdateStateNormal NsUpdateState
var pNsUpdateStatePaused NsUpdateState
var pNsUpdateStateUpdating NsUpdateState
var pNsUpdateStateInvalid NsUpdateState
var pNsUpdateStateTimedOut NsUpdateState
var pNsUpdateStateFailed NsUpdateState

// NsUpdateStateNormal enum exports
var NsUpdateStateNormal *NsUpdateState

// NsUpdateStatePaused enum exports
var NsUpdateStatePaused *NsUpdateState

// NsUpdateStateUpdating enum exports
var NsUpdateStateUpdating *NsUpdateState

// NsUpdateStateInvalid enum exports
var NsUpdateStateInvalid *NsUpdateState

// NsUpdateStateTimedOut enum exports
var NsUpdateStateTimedOut *NsUpdateState

// NsUpdateStateFailed enum exports
var NsUpdateStateFailed *NsUpdateState

func init() {
	pNsUpdateStateNormal = cNsUpdateStateNormal
	NsUpdateStateNormal = &pNsUpdateStateNormal

	pNsUpdateStatePaused = cNsUpdateStatePaused
	NsUpdateStatePaused = &pNsUpdateStatePaused

	pNsUpdateStateUpdating = cNsUpdateStateUpdating
	NsUpdateStateUpdating = &pNsUpdateStateUpdating

	pNsUpdateStateInvalid = cNsUpdateStateInvalid
	NsUpdateStateInvalid = &pNsUpdateStateInvalid

	pNsUpdateStateTimedOut = cNsUpdateStateTimedOut
	NsUpdateStateTimedOut = &pNsUpdateStateTimedOut

	pNsUpdateStateFailed = cNsUpdateStateFailed
	NsUpdateStateFailed = &pNsUpdateStateFailed

}
