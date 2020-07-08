// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsAgentType Enum.

type NsAgentType string

const (
	cNsAgentTypeSmis        NsAgentType = "smis"
	cNsAgentTypeVvol        NsAgentType = "vvol"
	cNsAgentTypeOpenstack   NsAgentType = "openstack"
	cNsAgentTypeOpenstackv2 NsAgentType = "openstackv2"
	cNsAgentTypeNone        NsAgentType = "none"
)

var pNsAgentTypeSmis NsAgentType
var pNsAgentTypeVvol NsAgentType
var pNsAgentTypeOpenstack NsAgentType
var pNsAgentTypeOpenstackv2 NsAgentType
var pNsAgentTypeNone NsAgentType

// Export
var NsAgentTypeSmis *NsAgentType
var NsAgentTypeVvol *NsAgentType
var NsAgentTypeOpenstack *NsAgentType
var NsAgentTypeOpenstackv2 *NsAgentType
var NsAgentTypeNone *NsAgentType

func init() {
	pNsAgentTypeSmis = cNsAgentTypeSmis
	NsAgentTypeSmis = &pNsAgentTypeSmis

	pNsAgentTypeVvol = cNsAgentTypeVvol
	NsAgentTypeVvol = &pNsAgentTypeVvol

	pNsAgentTypeOpenstack = cNsAgentTypeOpenstack
	NsAgentTypeOpenstack = &pNsAgentTypeOpenstack

	pNsAgentTypeOpenstackv2 = cNsAgentTypeOpenstackv2
	NsAgentTypeOpenstackv2 = &pNsAgentTypeOpenstackv2

	pNsAgentTypeNone = cNsAgentTypeNone
	NsAgentTypeNone = &pNsAgentTypeNone

}
