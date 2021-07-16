// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

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

// NsAppServerTypeVss enum exports
var NsAppServerTypeVss *NsAppServerType

// NsAppServerTypeVmware enum exports
var NsAppServerTypeVmware *NsAppServerType

// NsAppServerTypeCisco enum exports
var NsAppServerTypeCisco *NsAppServerType

// NsAppServerTypeContainerNode enum exports
var NsAppServerTypeContainerNode *NsAppServerType

// NsAppServerTypeStackVision enum exports
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
