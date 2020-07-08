// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsShelfHwStateDiscovering *NsShelfHwState
var NsShelfHwStateDisconnected *NsShelfHwState
var NsShelfHwStateVoid *NsShelfHwState
var NsShelfHwStateReady *NsShelfHwState
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
