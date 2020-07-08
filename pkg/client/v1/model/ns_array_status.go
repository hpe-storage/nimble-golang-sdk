// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsArrayStatus Enum.

type NsArrayStatus string

const (
	cNsArrayStatusUnreachable NsArrayStatus = "unreachable"
	cNsArrayStatusReachable   NsArrayStatus = "reachable"
)

var pNsArrayStatusUnreachable NsArrayStatus
var pNsArrayStatusReachable NsArrayStatus

// Export
var NsArrayStatusUnreachable *NsArrayStatus
var NsArrayStatusReachable *NsArrayStatus

func init() {
	pNsArrayStatusUnreachable = cNsArrayStatusUnreachable
	NsArrayStatusUnreachable = &pNsArrayStatusUnreachable

	pNsArrayStatusReachable = cNsArrayStatusReachable
	NsArrayStatusReachable = &pNsArrayStatusReachable

}
