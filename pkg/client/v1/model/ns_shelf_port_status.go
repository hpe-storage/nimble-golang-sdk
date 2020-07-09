// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsShelfPortStatus Enum.

type NsShelfPortStatus string

const (
	cNsShelfPortStatusConnected    NsShelfPortStatus = "connected"
	cNsShelfPortStatusDisconnected NsShelfPortStatus = "disconnected"
	cNsShelfPortStatusDisabled     NsShelfPortStatus = "disabled"
	cNsShelfPortStatusUnknown      NsShelfPortStatus = "unknown"
)

var pNsShelfPortStatusConnected NsShelfPortStatus
var pNsShelfPortStatusDisconnected NsShelfPortStatus
var pNsShelfPortStatusDisabled NsShelfPortStatus
var pNsShelfPortStatusUnknown NsShelfPortStatus

// NsShelfPortStatusConnected enum exports
var NsShelfPortStatusConnected *NsShelfPortStatus

// NsShelfPortStatusDisconnected enum exports
var NsShelfPortStatusDisconnected *NsShelfPortStatus

// NsShelfPortStatusDisabled enum exports
var NsShelfPortStatusDisabled *NsShelfPortStatus

// NsShelfPortStatusUnknown enum exports
var NsShelfPortStatusUnknown *NsShelfPortStatus

func init() {
	pNsShelfPortStatusConnected = cNsShelfPortStatusConnected
	NsShelfPortStatusConnected = &pNsShelfPortStatusConnected

	pNsShelfPortStatusDisconnected = cNsShelfPortStatusDisconnected
	NsShelfPortStatusDisconnected = &pNsShelfPortStatusDisconnected

	pNsShelfPortStatusDisabled = cNsShelfPortStatusDisabled
	NsShelfPortStatusDisabled = &pNsShelfPortStatusDisabled

	pNsShelfPortStatusUnknown = cNsShelfPortStatusUnknown
	NsShelfPortStatusUnknown = &pNsShelfPortStatusUnknown

}
