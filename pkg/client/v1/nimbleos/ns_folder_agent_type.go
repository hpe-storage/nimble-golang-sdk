// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsFolderAgentType Enum.

type NsFolderAgentType string

const (
	cNsFolderAgentTypeSmis      NsFolderAgentType = "smis"
	cNsFolderAgentTypeVvol      NsFolderAgentType = "vvol"
	cNsFolderAgentTypeOpenstack NsFolderAgentType = "openstack"
	cNsFolderAgentTypeNone      NsFolderAgentType = "none"
)

var pNsFolderAgentTypeSmis NsFolderAgentType
var pNsFolderAgentTypeVvol NsFolderAgentType
var pNsFolderAgentTypeOpenstack NsFolderAgentType
var pNsFolderAgentTypeNone NsFolderAgentType

// NsFolderAgentTypeSmis enum exports
var NsFolderAgentTypeSmis *NsFolderAgentType

// NsFolderAgentTypeVvol enum exports
var NsFolderAgentTypeVvol *NsFolderAgentType

// NsFolderAgentTypeOpenstack enum exports
var NsFolderAgentTypeOpenstack *NsFolderAgentType

// NsFolderAgentTypeNone enum exports
var NsFolderAgentTypeNone *NsFolderAgentType

func init() {
	pNsFolderAgentTypeSmis = cNsFolderAgentTypeSmis
	NsFolderAgentTypeSmis = &pNsFolderAgentTypeSmis

	pNsFolderAgentTypeVvol = cNsFolderAgentTypeVvol
	NsFolderAgentTypeVvol = &pNsFolderAgentTypeVvol

	pNsFolderAgentTypeOpenstack = cNsFolderAgentTypeOpenstack
	NsFolderAgentTypeOpenstack = &pNsFolderAgentTypeOpenstack

	pNsFolderAgentTypeNone = cNsFolderAgentTypeNone
	NsFolderAgentTypeNone = &pNsFolderAgentTypeNone

}
