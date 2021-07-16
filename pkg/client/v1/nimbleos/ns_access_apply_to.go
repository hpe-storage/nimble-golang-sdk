// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsAccessApplyTo Enum.

type NsAccessApplyTo string

const (
	cNsAccessApplyToVolume       NsAccessApplyTo = "volume"
	cNsAccessApplyToPe           NsAccessApplyTo = "pe"
	cNsAccessApplyToVvolVolume   NsAccessApplyTo = "vvol_volume"
	cNsAccessApplyToVvolSnapshot NsAccessApplyTo = "vvol_snapshot"
	cNsAccessApplyToSnapshot     NsAccessApplyTo = "snapshot"
	cNsAccessApplyToBoth         NsAccessApplyTo = "both"
)

var pNsAccessApplyToVolume NsAccessApplyTo
var pNsAccessApplyToPe NsAccessApplyTo
var pNsAccessApplyToVvolVolume NsAccessApplyTo
var pNsAccessApplyToVvolSnapshot NsAccessApplyTo
var pNsAccessApplyToSnapshot NsAccessApplyTo
var pNsAccessApplyToBoth NsAccessApplyTo

// NsAccessApplyToVolume enum exports
var NsAccessApplyToVolume *NsAccessApplyTo

// NsAccessApplyToPe enum exports
var NsAccessApplyToPe *NsAccessApplyTo

// NsAccessApplyToVvolVolume enum exports
var NsAccessApplyToVvolVolume *NsAccessApplyTo

// NsAccessApplyToVvolSnapshot enum exports
var NsAccessApplyToVvolSnapshot *NsAccessApplyTo

// NsAccessApplyToSnapshot enum exports
var NsAccessApplyToSnapshot *NsAccessApplyTo

// NsAccessApplyToBoth enum exports
var NsAccessApplyToBoth *NsAccessApplyTo

func init() {
	pNsAccessApplyToVolume = cNsAccessApplyToVolume
	NsAccessApplyToVolume = &pNsAccessApplyToVolume

	pNsAccessApplyToPe = cNsAccessApplyToPe
	NsAccessApplyToPe = &pNsAccessApplyToPe

	pNsAccessApplyToVvolVolume = cNsAccessApplyToVvolVolume
	NsAccessApplyToVvolVolume = &pNsAccessApplyToVvolVolume

	pNsAccessApplyToVvolSnapshot = cNsAccessApplyToVvolSnapshot
	NsAccessApplyToVvolSnapshot = &pNsAccessApplyToVvolSnapshot

	pNsAccessApplyToSnapshot = cNsAccessApplyToSnapshot
	NsAccessApplyToSnapshot = &pNsAccessApplyToSnapshot

	pNsAccessApplyToBoth = cNsAccessApplyToBoth
	NsAccessApplyToBoth = &pNsAccessApplyToBoth

}
