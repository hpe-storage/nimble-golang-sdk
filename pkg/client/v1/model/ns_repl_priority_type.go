// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsReplPriorityType Enum.

type NsReplPriorityType string

const (
	cNsReplPriorityTypeNormal NsReplPriorityType = "normal"
	cNsReplPriorityTypeHigh   NsReplPriorityType = "high"
)

var pNsReplPriorityTypeNormal NsReplPriorityType
var pNsReplPriorityTypeHigh NsReplPriorityType

// Export
var NsReplPriorityTypeNormal *NsReplPriorityType
var NsReplPriorityTypeHigh *NsReplPriorityType

func init() {
	pNsReplPriorityTypeNormal = cNsReplPriorityTypeNormal
	NsReplPriorityTypeNormal = &pNsReplPriorityTypeNormal

	pNsReplPriorityTypeHigh = cNsReplPriorityTypeHigh
	NsReplPriorityTypeHigh = &pNsReplPriorityTypeHigh

}
