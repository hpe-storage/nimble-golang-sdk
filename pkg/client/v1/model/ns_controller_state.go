// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsControllerStateStartActive *NsControllerState
var NsControllerStateStartStandby *NsControllerState
var NsControllerStateStale *NsControllerState
var NsControllerStateStandby *NsControllerState
var NsControllerStateActive *NsControllerState
var NsControllerStateSolo *NsControllerState
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
