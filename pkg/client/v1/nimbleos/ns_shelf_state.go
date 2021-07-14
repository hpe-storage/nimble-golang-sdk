// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsShelfState Enum.

type NsShelfState string

const (
 cNsShelfStateDiscovering NsShelfState = "discovering"
 cNsShelfStateDisconnected NsShelfState = "disconnected"
 cNsShelfStateVoid NsShelfState = "void"
 cNsShelfStateAvailable NsShelfState = "available"
 cNsShelfStateOnline NsShelfState = "online"
 cNsShelfStateFaulty NsShelfState = "faulty"
 cNsShelfStateForeign NsShelfState = "foreign"
 cNsShelfStateUnknown NsShelfState = "unknown"
)

var pNsShelfStateDiscovering NsShelfState
var pNsShelfStateDisconnected NsShelfState
var pNsShelfStateVoid NsShelfState
var pNsShelfStateAvailable NsShelfState
var pNsShelfStateOnline NsShelfState
var pNsShelfStateFaulty NsShelfState
var pNsShelfStateForeign NsShelfState
var pNsShelfStateUnknown NsShelfState

// NsShelfStateDiscovering enum exports
var NsShelfStateDiscovering *NsShelfState

// NsShelfStateDisconnected enum exports
var NsShelfStateDisconnected *NsShelfState

// NsShelfStateVoid enum exports
var NsShelfStateVoid *NsShelfState

// NsShelfStateAvailable enum exports
var NsShelfStateAvailable *NsShelfState

// NsShelfStateOnline enum exports
var NsShelfStateOnline *NsShelfState

// NsShelfStateFaulty enum exports
var NsShelfStateFaulty *NsShelfState

// NsShelfStateForeign enum exports
var NsShelfStateForeign *NsShelfState

// NsShelfStateUnknown enum exports
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

