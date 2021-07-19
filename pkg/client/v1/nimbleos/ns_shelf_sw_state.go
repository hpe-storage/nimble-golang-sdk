// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsShelfSwState Enum.

type NsShelfSwState string

const (
	cNsShelfSwStateAvailable NsShelfSwState = "available"
	cNsShelfSwStateOnline    NsShelfSwState = "online"
	cNsShelfSwStateForeign   NsShelfSwState = "foreign"
	cNsShelfSwStateUnknown   NsShelfSwState = "unknown"
)

var pNsShelfSwStateAvailable NsShelfSwState
var pNsShelfSwStateOnline NsShelfSwState
var pNsShelfSwStateForeign NsShelfSwState
var pNsShelfSwStateUnknown NsShelfSwState

// NsShelfSwStateAvailable enum exports
var NsShelfSwStateAvailable *NsShelfSwState

// NsShelfSwStateOnline enum exports
var NsShelfSwStateOnline *NsShelfSwState

// NsShelfSwStateForeign enum exports
var NsShelfSwStateForeign *NsShelfSwState

// NsShelfSwStateUnknown enum exports
var NsShelfSwStateUnknown *NsShelfSwState

func init() {
	pNsShelfSwStateAvailable = cNsShelfSwStateAvailable
	NsShelfSwStateAvailable = &pNsShelfSwStateAvailable

	pNsShelfSwStateOnline = cNsShelfSwStateOnline
	NsShelfSwStateOnline = &pNsShelfSwStateOnline

	pNsShelfSwStateForeign = cNsShelfSwStateForeign
	NsShelfSwStateForeign = &pNsShelfSwStateForeign

	pNsShelfSwStateUnknown = cNsShelfSwStateUnknown
	NsShelfSwStateUnknown = &pNsShelfSwStateUnknown

}
