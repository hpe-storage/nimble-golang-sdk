// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsHCIConfigType Enum.

type NsHCIConfigType string

const (
	cNsHCIConfigTypeNode    NsHCIConfigType = "node"
	cNsHCIConfigTypeCluster NsHCIConfigType = "cluster"
	cNsHCIConfigTypeBlock   NsHCIConfigType = "block"
	cNsHCIConfigTypeSwitch  NsHCIConfigType = "switch"
)

var pNsHCIConfigTypeNode NsHCIConfigType
var pNsHCIConfigTypeCluster NsHCIConfigType
var pNsHCIConfigTypeBlock NsHCIConfigType
var pNsHCIConfigTypeSwitch NsHCIConfigType

// NsHCIConfigTypeNode enum exports
var NsHCIConfigTypeNode *NsHCIConfigType

// NsHCIConfigTypeCluster enum exports
var NsHCIConfigTypeCluster *NsHCIConfigType

// NsHCIConfigTypeBlock enum exports
var NsHCIConfigTypeBlock *NsHCIConfigType

// NsHCIConfigTypeSwitch enum exports
var NsHCIConfigTypeSwitch *NsHCIConfigType

func init() {
	pNsHCIConfigTypeNode = cNsHCIConfigTypeNode
	NsHCIConfigTypeNode = &pNsHCIConfigTypeNode

	pNsHCIConfigTypeCluster = cNsHCIConfigTypeCluster
	NsHCIConfigTypeCluster = &pNsHCIConfigTypeCluster

	pNsHCIConfigTypeBlock = cNsHCIConfigTypeBlock
	NsHCIConfigTypeBlock = &pNsHCIConfigTypeBlock

	pNsHCIConfigTypeSwitch = cNsHCIConfigTypeSwitch
	NsHCIConfigTypeSwitch = &pNsHCIConfigTypeSwitch

}
