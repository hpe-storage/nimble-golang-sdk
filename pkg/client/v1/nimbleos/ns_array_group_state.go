// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsArrayGroupState Enum.

type NsArrayGroupState string

const (
	cNsArrayGroupStateInvalid     NsArrayGroupState = "invalid"
	cNsArrayGroupStateInitialized NsArrayGroupState = "initialized"
	cNsArrayGroupStateUnused      NsArrayGroupState = "unused"
	cNsArrayGroupStateRemoving    NsArrayGroupState = "removing"
)

var pNsArrayGroupStateInvalid NsArrayGroupState
var pNsArrayGroupStateInitialized NsArrayGroupState
var pNsArrayGroupStateUnused NsArrayGroupState
var pNsArrayGroupStateRemoving NsArrayGroupState

// NsArrayGroupStateInvalid enum exports
var NsArrayGroupStateInvalid *NsArrayGroupState

// NsArrayGroupStateInitialized enum exports
var NsArrayGroupStateInitialized *NsArrayGroupState

// NsArrayGroupStateUnused enum exports
var NsArrayGroupStateUnused *NsArrayGroupState

// NsArrayGroupStateRemoving enum exports
var NsArrayGroupStateRemoving *NsArrayGroupState

func init() {
	pNsArrayGroupStateInvalid = cNsArrayGroupStateInvalid
	NsArrayGroupStateInvalid = &pNsArrayGroupStateInvalid

	pNsArrayGroupStateInitialized = cNsArrayGroupStateInitialized
	NsArrayGroupStateInitialized = &pNsArrayGroupStateInitialized

	pNsArrayGroupStateUnused = cNsArrayGroupStateUnused
	NsArrayGroupStateUnused = &pNsArrayGroupStateUnused

	pNsArrayGroupStateRemoving = cNsArrayGroupStateRemoving
	NsArrayGroupStateRemoving = &pNsArrayGroupStateRemoving

}
