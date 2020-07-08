// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsShelfSesMasterSwStateWantMaster *NsShelfSesMasterSwState
var NsShelfSesMasterSwStateReleaseMaster *NsShelfSesMasterSwState
var NsShelfSesMasterSwStateNotMaster *NsShelfSesMasterSwState
var NsShelfSesMasterSwStateUnknown *NsShelfSesMasterSwState
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
