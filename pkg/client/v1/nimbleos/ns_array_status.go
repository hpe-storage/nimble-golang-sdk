// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsArrayStatus Enum.

type NsArrayStatus string

const (
	cNsArrayStatusUnreachable NsArrayStatus = "unreachable"
	cNsArrayStatusReachable   NsArrayStatus = "reachable"
)

var pNsArrayStatusUnreachable NsArrayStatus
var pNsArrayStatusReachable NsArrayStatus

// NsArrayStatusUnreachable enum exports
var NsArrayStatusUnreachable *NsArrayStatus

// NsArrayStatusReachable enum exports
var NsArrayStatusReachable *NsArrayStatus

func init() {
	pNsArrayStatusUnreachable = cNsArrayStatusUnreachable
	NsArrayStatusUnreachable = &pNsArrayStatusUnreachable

	pNsArrayStatusReachable = cNsArrayStatusReachable
	NsArrayStatusReachable = &pNsArrayStatusReachable

}
