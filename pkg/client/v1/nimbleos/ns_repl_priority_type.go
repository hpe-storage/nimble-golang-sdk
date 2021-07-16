// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsReplPriorityType Enum.

type NsReplPriorityType string

const (
	cNsReplPriorityTypeNormal NsReplPriorityType = "normal"
	cNsReplPriorityTypeHigh   NsReplPriorityType = "high"
)

var pNsReplPriorityTypeNormal NsReplPriorityType
var pNsReplPriorityTypeHigh NsReplPriorityType

// NsReplPriorityTypeNormal enum exports
var NsReplPriorityTypeNormal *NsReplPriorityType

// NsReplPriorityTypeHigh enum exports
var NsReplPriorityTypeHigh *NsReplPriorityType

func init() {
	pNsReplPriorityTypeNormal = cNsReplPriorityTypeNormal
	NsReplPriorityTypeNormal = &pNsReplPriorityTypeNormal

	pNsReplPriorityTypeHigh = cNsReplPriorityTypeHigh
	NsReplPriorityTypeHigh = &pNsReplPriorityTypeHigh

}
