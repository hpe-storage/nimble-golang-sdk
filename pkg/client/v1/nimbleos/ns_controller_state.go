// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsControllerState Enum.

type NsControllerState string

const (
	cNsControllerStateStartActive  NsControllerState = "start_active"
	cNsControllerStateStartStandby NsControllerState = "start_standby"
	cNsControllerStateStale        NsControllerState = "stale"
	cNsControllerStateStandby      NsControllerState = "standby"
	cNsControllerStateActive       NsControllerState = "active"
	cNsControllerStateSolo         NsControllerState = "solo"
	cNsControllerStateNone         NsControllerState = "none"
)

var pNsControllerStateStartActive NsControllerState
var pNsControllerStateStartStandby NsControllerState
var pNsControllerStateStale NsControllerState
var pNsControllerStateStandby NsControllerState
var pNsControllerStateActive NsControllerState
var pNsControllerStateSolo NsControllerState
var pNsControllerStateNone NsControllerState

// NsControllerStateStartActive enum exports
var NsControllerStateStartActive *NsControllerState

// NsControllerStateStartStandby enum exports
var NsControllerStateStartStandby *NsControllerState

// NsControllerStateStale enum exports
var NsControllerStateStale *NsControllerState

// NsControllerStateStandby enum exports
var NsControllerStateStandby *NsControllerState

// NsControllerStateActive enum exports
var NsControllerStateActive *NsControllerState

// NsControllerStateSolo enum exports
var NsControllerStateSolo *NsControllerState

// NsControllerStateNone enum exports
var NsControllerStateNone *NsControllerState

func init() {
	pNsControllerStateStartActive = cNsControllerStateStartActive
	NsControllerStateStartActive = &pNsControllerStateStartActive

	pNsControllerStateStartStandby = cNsControllerStateStartStandby
	NsControllerStateStartStandby = &pNsControllerStateStartStandby

	pNsControllerStateStale = cNsControllerStateStale
	NsControllerStateStale = &pNsControllerStateStale

	pNsControllerStateStandby = cNsControllerStateStandby
	NsControllerStateStandby = &pNsControllerStateStandby

	pNsControllerStateActive = cNsControllerStateActive
	NsControllerStateActive = &pNsControllerStateActive

	pNsControllerStateSolo = cNsControllerStateSolo
	NsControllerStateSolo = &pNsControllerStateSolo

	pNsControllerStateNone = cNsControllerStateNone
	NsControllerStateNone = &pNsControllerStateNone

}
