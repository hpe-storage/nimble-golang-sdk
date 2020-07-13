// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsSubnetType Enum.

type NsSubnetType string

const (
	cNsSubnetTypeMgmt         NsSubnetType = "mgmt"
	cNsSubnetTypeUnconfigured NsSubnetType = "unconfigured"
	cNsSubnetTypeData         NsSubnetType = "data"
	cNsSubnetTypeMgmtData     NsSubnetType = "mgmt_data"
	cNsSubnetTypeInvalid      NsSubnetType = "invalid"
)

var pNsSubnetTypeMgmt NsSubnetType
var pNsSubnetTypeUnconfigured NsSubnetType
var pNsSubnetTypeData NsSubnetType
var pNsSubnetTypeMgmtData NsSubnetType
var pNsSubnetTypeInvalid NsSubnetType

// NsSubnetTypeMgmt enum exports
var NsSubnetTypeMgmt *NsSubnetType

// NsSubnetTypeUnconfigured enum exports
var NsSubnetTypeUnconfigured *NsSubnetType

// NsSubnetTypeData enum exports
var NsSubnetTypeData *NsSubnetType

// NsSubnetTypeMgmtData enum exports
var NsSubnetTypeMgmtData *NsSubnetType

// NsSubnetTypeInvalid enum exports
var NsSubnetTypeInvalid *NsSubnetType

func init() {
	pNsSubnetTypeMgmt = cNsSubnetTypeMgmt
	NsSubnetTypeMgmt = &pNsSubnetTypeMgmt

	pNsSubnetTypeUnconfigured = cNsSubnetTypeUnconfigured
	NsSubnetTypeUnconfigured = &pNsSubnetTypeUnconfigured

	pNsSubnetTypeData = cNsSubnetTypeData
	NsSubnetTypeData = &pNsSubnetTypeData

	pNsSubnetTypeMgmtData = cNsSubnetTypeMgmtData
	NsSubnetTypeMgmtData = &pNsSubnetTypeMgmtData

	pNsSubnetTypeInvalid = cNsSubnetTypeInvalid
	NsSubnetTypeInvalid = &pNsSubnetTypeInvalid

}
