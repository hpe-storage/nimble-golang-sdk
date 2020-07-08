// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsAppServerType Enum.

type NsAppServerType string

const (
	cNsAppServerTypeVss           NsAppServerType = "vss"
	cNsAppServerTypeVmware        NsAppServerType = "vmware"
	cNsAppServerTypeCisco         NsAppServerType = "cisco"
	cNsAppServerTypeContainerNode NsAppServerType = "container_node"
	cNsAppServerTypeStackVision   NsAppServerType = "stack_vision"
)

var pNsAppServerTypeVss NsAppServerType
var pNsAppServerTypeVmware NsAppServerType
var pNsAppServerTypeCisco NsAppServerType
var pNsAppServerTypeContainerNode NsAppServerType
var pNsAppServerTypeStackVision NsAppServerType

// Export
var NsAppServerTypeVss *NsAppServerType
var NsAppServerTypeVmware *NsAppServerType
var NsAppServerTypeCisco *NsAppServerType
var NsAppServerTypeContainerNode *NsAppServerType
var NsAppServerTypeStackVision *NsAppServerType

func init() {
	pNsAppServerTypeVss = cNsAppServerTypeVss
	NsAppServerTypeVss = &pNsAppServerTypeVss

	pNsAppServerTypeVmware = cNsAppServerTypeVmware
	NsAppServerTypeVmware = &pNsAppServerTypeVmware

	pNsAppServerTypeCisco = cNsAppServerTypeCisco
	NsAppServerTypeCisco = &pNsAppServerTypeCisco

	pNsAppServerTypeContainerNode = cNsAppServerTypeContainerNode
	NsAppServerTypeContainerNode = &pNsAppServerTypeContainerNode

	pNsAppServerTypeStackVision = cNsAppServerTypeStackVision
	NsAppServerTypeStackVision = &pNsAppServerTypeStackVision

}
