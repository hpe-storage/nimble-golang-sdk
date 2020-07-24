// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsGroupFailoverMode Enum.

type NsGroupFailoverMode string

const (
	cNsGroupFailoverModeManual    NsGroupFailoverMode = "Manual"
	cNsGroupFailoverModeAutomatic NsGroupFailoverMode = "Automatic"
)

var pNsGroupFailoverModeManual NsGroupFailoverMode
var pNsGroupFailoverModeAutomatic NsGroupFailoverMode

// NsGroupFailoverModeManual enum exports
var NsGroupFailoverModeManual *NsGroupFailoverMode

// NsGroupFailoverModeAutomatic enum exports
var NsGroupFailoverModeAutomatic *NsGroupFailoverMode

func init() {
	pNsGroupFailoverModeManual = cNsGroupFailoverModeManual
	NsGroupFailoverModeManual = &pNsGroupFailoverModeManual

	pNsGroupFailoverModeAutomatic = cNsGroupFailoverModeAutomatic
	NsGroupFailoverModeAutomatic = &pNsGroupFailoverModeAutomatic

}
