// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsShelfHwState Enum.

type NsShelfHwState string

const (
	cNsShelfHwStateDiscovering  NsShelfHwState = "discovering"
	cNsShelfHwStateDisconnected NsShelfHwState = "disconnected"
	cNsShelfHwStateVoid         NsShelfHwState = "void"
	cNsShelfHwStateReady        NsShelfHwState = "ready"
	cNsShelfHwStateFaulty       NsShelfHwState = "faulty"
)

var pNsShelfHwStateDiscovering NsShelfHwState
var pNsShelfHwStateDisconnected NsShelfHwState
var pNsShelfHwStateVoid NsShelfHwState
var pNsShelfHwStateReady NsShelfHwState
var pNsShelfHwStateFaulty NsShelfHwState

// NsShelfHwStateDiscovering enum exports
var NsShelfHwStateDiscovering *NsShelfHwState

// NsShelfHwStateDisconnected enum exports
var NsShelfHwStateDisconnected *NsShelfHwState

// NsShelfHwStateVoid enum exports
var NsShelfHwStateVoid *NsShelfHwState

// NsShelfHwStateReady enum exports
var NsShelfHwStateReady *NsShelfHwState

// NsShelfHwStateFaulty enum exports
var NsShelfHwStateFaulty *NsShelfHwState

func init() {
	pNsShelfHwStateDiscovering = cNsShelfHwStateDiscovering
	NsShelfHwStateDiscovering = &pNsShelfHwStateDiscovering

	pNsShelfHwStateDisconnected = cNsShelfHwStateDisconnected
	NsShelfHwStateDisconnected = &pNsShelfHwStateDisconnected

	pNsShelfHwStateVoid = cNsShelfHwStateVoid
	NsShelfHwStateVoid = &pNsShelfHwStateVoid

	pNsShelfHwStateReady = cNsShelfHwStateReady
	NsShelfHwStateReady = &pNsShelfHwStateReady

	pNsShelfHwStateFaulty = cNsShelfHwStateFaulty
	NsShelfHwStateFaulty = &pNsShelfHwStateFaulty

}
