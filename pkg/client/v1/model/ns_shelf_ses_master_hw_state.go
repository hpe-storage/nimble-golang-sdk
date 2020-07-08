// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsShelfSesMasterHwState Enum.

type NsShelfSesMasterHwState string

const (
	cNsShelfSesMasterHwStateNotMaster NsShelfSesMasterHwState = "not master"
	cNsShelfSesMasterHwStateFailed    NsShelfSesMasterHwState = "failed"
	cNsShelfSesMasterHwStateUnknown   NsShelfSesMasterHwState = "unknown"
	cNsShelfSesMasterHwStateMaster    NsShelfSesMasterHwState = "master"
)

var pNsShelfSesMasterHwStateNotMaster NsShelfSesMasterHwState
var pNsShelfSesMasterHwStateFailed NsShelfSesMasterHwState
var pNsShelfSesMasterHwStateUnknown NsShelfSesMasterHwState
var pNsShelfSesMasterHwStateMaster NsShelfSesMasterHwState

// Export
var NsShelfSesMasterHwStateNotMaster *NsShelfSesMasterHwState
var NsShelfSesMasterHwStateFailed *NsShelfSesMasterHwState
var NsShelfSesMasterHwStateUnknown *NsShelfSesMasterHwState
var NsShelfSesMasterHwStateMaster *NsShelfSesMasterHwState

func init() {
	pNsShelfSesMasterHwStateNotMaster = cNsShelfSesMasterHwStateNotMaster
	NsShelfSesMasterHwStateNotMaster = &pNsShelfSesMasterHwStateNotMaster

	pNsShelfSesMasterHwStateFailed = cNsShelfSesMasterHwStateFailed
	NsShelfSesMasterHwStateFailed = &pNsShelfSesMasterHwStateFailed

	pNsShelfSesMasterHwStateUnknown = cNsShelfSesMasterHwStateUnknown
	NsShelfSesMasterHwStateUnknown = &pNsShelfSesMasterHwStateUnknown

	pNsShelfSesMasterHwStateMaster = cNsShelfSesMasterHwStateMaster
	NsShelfSesMasterHwStateMaster = &pNsShelfSesMasterHwStateMaster

}
