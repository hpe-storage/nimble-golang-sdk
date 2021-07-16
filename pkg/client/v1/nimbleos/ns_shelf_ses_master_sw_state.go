// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsShelfSesMasterSwState Enum.

type NsShelfSesMasterSwState string

const (
	cNsShelfSesMasterSwStateWantMaster    NsShelfSesMasterSwState = "want master"
	cNsShelfSesMasterSwStateReleaseMaster NsShelfSesMasterSwState = "release master"
	cNsShelfSesMasterSwStateNotMaster     NsShelfSesMasterSwState = "not master"
	cNsShelfSesMasterSwStateUnknown       NsShelfSesMasterSwState = "unknown"
	cNsShelfSesMasterSwStateMaster        NsShelfSesMasterSwState = "master"
)

var pNsShelfSesMasterSwStateWantMaster NsShelfSesMasterSwState
var pNsShelfSesMasterSwStateReleaseMaster NsShelfSesMasterSwState
var pNsShelfSesMasterSwStateNotMaster NsShelfSesMasterSwState
var pNsShelfSesMasterSwStateUnknown NsShelfSesMasterSwState
var pNsShelfSesMasterSwStateMaster NsShelfSesMasterSwState

// NsShelfSesMasterSwStateWantMaster enum exports
var NsShelfSesMasterSwStateWantMaster *NsShelfSesMasterSwState

// NsShelfSesMasterSwStateReleaseMaster enum exports
var NsShelfSesMasterSwStateReleaseMaster *NsShelfSesMasterSwState

// NsShelfSesMasterSwStateNotMaster enum exports
var NsShelfSesMasterSwStateNotMaster *NsShelfSesMasterSwState

// NsShelfSesMasterSwStateUnknown enum exports
var NsShelfSesMasterSwStateUnknown *NsShelfSesMasterSwState

// NsShelfSesMasterSwStateMaster enum exports
var NsShelfSesMasterSwStateMaster *NsShelfSesMasterSwState

func init() {
	pNsShelfSesMasterSwStateWantMaster = cNsShelfSesMasterSwStateWantMaster
	NsShelfSesMasterSwStateWantMaster = &pNsShelfSesMasterSwStateWantMaster

	pNsShelfSesMasterSwStateReleaseMaster = cNsShelfSesMasterSwStateReleaseMaster
	NsShelfSesMasterSwStateReleaseMaster = &pNsShelfSesMasterSwStateReleaseMaster

	pNsShelfSesMasterSwStateNotMaster = cNsShelfSesMasterSwStateNotMaster
	NsShelfSesMasterSwStateNotMaster = &pNsShelfSesMasterSwStateNotMaster

	pNsShelfSesMasterSwStateUnknown = cNsShelfSesMasterSwStateUnknown
	NsShelfSesMasterSwStateUnknown = &pNsShelfSesMasterSwStateUnknown

	pNsShelfSesMasterSwStateMaster = cNsShelfSesMasterSwStateMaster
	NsShelfSesMasterSwStateMaster = &pNsShelfSesMasterSwStateMaster

}
