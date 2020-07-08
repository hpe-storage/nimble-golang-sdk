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

// Export
var NsShelfPortStatusConnected *NsShelfPortStatus
var NsShelfPortStatusDisconnected *NsShelfPortStatus
var NsShelfPortStatusDisabled *NsShelfPortStatus
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
