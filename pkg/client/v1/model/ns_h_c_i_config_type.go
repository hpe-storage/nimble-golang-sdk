// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsHCIConfigType Enum.

type NsHCIConfigType string

const (
	cNsHCIConfigTypeCluster NsHCIConfigType = "cluster"
	cNsHCIConfigTypeNode    NsHCIConfigType = "node"
	cNsHCIConfigTypeBlock   NsHCIConfigType = "block"
)

var pNsHCIConfigTypeCluster NsHCIConfigType
var pNsHCIConfigTypeNode NsHCIConfigType
var pNsHCIConfigTypeBlock NsHCIConfigType

// Export
var NsHCIConfigTypeCluster *NsHCIConfigType
var NsHCIConfigTypeNode *NsHCIConfigType
var NsHCIConfigTypeBlock *NsHCIConfigType

func init() {
	pNsHCIConfigTypeCluster = cNsHCIConfigTypeCluster
	NsHCIConfigTypeCluster = &pNsHCIConfigTypeCluster

	pNsHCIConfigTypeNode = cNsHCIConfigTypeNode
	NsHCIConfigTypeNode = &pNsHCIConfigTypeNode

	pNsHCIConfigTypeBlock = cNsHCIConfigTypeBlock
	NsHCIConfigTypeBlock = &pNsHCIConfigTypeBlock

}
