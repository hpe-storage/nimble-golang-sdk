// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsSubnetTypeMgmt *NsSubnetType
var NsSubnetTypeUnconfigured *NsSubnetType
var NsSubnetTypeData *NsSubnetType
var NsSubnetTypeMgmtData *NsSubnetType
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
