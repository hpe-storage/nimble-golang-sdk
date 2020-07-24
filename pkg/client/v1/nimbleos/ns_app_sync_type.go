// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsAppSyncType Enum.

type NsAppSyncType string

const (
	cNsAppSyncTypeVss     NsAppSyncType = "vss"
	cNsAppSyncTypeVmware  NsAppSyncType = "vmware"
	cNsAppSyncTypeNone    NsAppSyncType = "none"
	cNsAppSyncTypeGeneric NsAppSyncType = "generic"
)

var pNsAppSyncTypeVss NsAppSyncType
var pNsAppSyncTypeVmware NsAppSyncType
var pNsAppSyncTypeNone NsAppSyncType
var pNsAppSyncTypeGeneric NsAppSyncType

// NsAppSyncTypeVss enum exports
var NsAppSyncTypeVss *NsAppSyncType

// NsAppSyncTypeVmware enum exports
var NsAppSyncTypeVmware *NsAppSyncType

// NsAppSyncTypeNone enum exports
var NsAppSyncTypeNone *NsAppSyncType

// NsAppSyncTypeGeneric enum exports
var NsAppSyncTypeGeneric *NsAppSyncType

func init() {
	pNsAppSyncTypeVss = cNsAppSyncTypeVss
	NsAppSyncTypeVss = &pNsAppSyncTypeVss

	pNsAppSyncTypeVmware = cNsAppSyncTypeVmware
	NsAppSyncTypeVmware = &pNsAppSyncTypeVmware

	pNsAppSyncTypeNone = cNsAppSyncTypeNone
	NsAppSyncTypeNone = &pNsAppSyncTypeNone

	pNsAppSyncTypeGeneric = cNsAppSyncTypeGeneric
	NsAppSyncTypeGeneric = &pNsAppSyncTypeGeneric

}
