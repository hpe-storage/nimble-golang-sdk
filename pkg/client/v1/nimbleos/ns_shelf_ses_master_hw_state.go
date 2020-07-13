// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

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

// NsShelfSesMasterHwStateNotMaster enum exports
var NsShelfSesMasterHwStateNotMaster *NsShelfSesMasterHwState

// NsShelfSesMasterHwStateFailed enum exports
var NsShelfSesMasterHwStateFailed *NsShelfSesMasterHwState

// NsShelfSesMasterHwStateUnknown enum exports
var NsShelfSesMasterHwStateUnknown *NsShelfSesMasterHwState

// NsShelfSesMasterHwStateMaster enum exports
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
