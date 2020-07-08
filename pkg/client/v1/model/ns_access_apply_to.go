// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsAccessApplyToVolume *NsAccessApplyTo
var NsAccessApplyToPe *NsAccessApplyTo
var NsAccessApplyToVvolVolume *NsAccessApplyTo
var NsAccessApplyToVvolSnapshot *NsAccessApplyTo
var NsAccessApplyToSnapshot *NsAccessApplyTo
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
