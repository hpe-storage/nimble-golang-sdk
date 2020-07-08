// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsShelfSwStateAvailable *NsShelfSwState
var NsShelfSwStateOnline *NsShelfSwState
var NsShelfSwStateForeign *NsShelfSwState
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
