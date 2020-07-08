// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsShelfState Enum.

type NsShelfState string

const (
	cNsShelfStateDiscovering  NsShelfState = "discovering"
	cNsShelfStateDisconnected NsShelfState = "disconnected"
	cNsShelfStateVoid         NsShelfState = "void"
	cNsShelfStateAvailable    NsShelfState = "available"
	cNsShelfStateOnline       NsShelfState = "online"
	cNsShelfStateFaulty       NsShelfState = "faulty"
	cNsShelfStateForeign      NsShelfState = "foreign"
	cNsShelfStateUnknown      NsShelfState = "unknown"
)

var pNsShelfStateDiscovering NsShelfState
var pNsShelfStateDisconnected NsShelfState
var pNsShelfStateVoid NsShelfState
var pNsShelfStateAvailable NsShelfState
var pNsShelfStateOnline NsShelfState
var pNsShelfStateFaulty NsShelfState
var pNsShelfStateForeign NsShelfState
var pNsShelfStateUnknown NsShelfState

// Export
var NsShelfStateDiscovering *NsShelfState
var NsShelfStateDisconnected *NsShelfState
var NsShelfStateVoid *NsShelfState
var NsShelfStateAvailable *NsShelfState
var NsShelfStateOnline *NsShelfState
var NsShelfStateFaulty *NsShelfState
var NsShelfStateForeign *NsShelfState
var NsShelfStateUnknown *NsShelfState

func init() {
	pNsShelfStateDiscovering = cNsShelfStateDiscovering
	NsShelfStateDiscovering = &pNsShelfStateDiscovering

	pNsShelfStateDisconnected = cNsShelfStateDisconnected
	NsShelfStateDisconnected = &pNsShelfStateDisconnected

	pNsShelfStateVoid = cNsShelfStateVoid
	NsShelfStateVoid = &pNsShelfStateVoid

	pNsShelfStateAvailable = cNsShelfStateAvailable
	NsShelfStateAvailable = &pNsShelfStateAvailable

	pNsShelfStateOnline = cNsShelfStateOnline
	NsShelfStateOnline = &pNsShelfStateOnline

	pNsShelfStateFaulty = cNsShelfStateFaulty
	NsShelfStateFaulty = &pNsShelfStateFaulty

	pNsShelfStateForeign = cNsShelfStateForeign
	NsShelfStateForeign = &pNsShelfStateForeign

	pNsShelfStateUnknown = cNsShelfStateUnknown
	NsShelfStateUnknown = &pNsShelfStateUnknown

}
