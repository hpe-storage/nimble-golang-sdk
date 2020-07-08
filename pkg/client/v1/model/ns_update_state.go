// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsUpdateStateNormal *NsUpdateState
var NsUpdateStatePaused *NsUpdateState
var NsUpdateStateUpdating *NsUpdateState
var NsUpdateStateInvalid *NsUpdateState
var NsUpdateStateTimedOut *NsUpdateState
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
