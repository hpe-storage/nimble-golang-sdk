// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsPlatNicType Enum.

type NsPlatNicType string

const (
	cNsPlatNicTypeNicTypeUnknown NsPlatNicType = "nic_type_unknown"
	cNsPlatNicTypeNicTypeTp      NsPlatNicType = "nic_type_tp"
	cNsPlatNicTypeNicTypeSfp     NsPlatNicType = "nic_type_sfp"
)

var pNsPlatNicTypeNicTypeUnknown NsPlatNicType
var pNsPlatNicTypeNicTypeTp NsPlatNicType
var pNsPlatNicTypeNicTypeSfp NsPlatNicType

// Export
var NsPlatNicTypeNicTypeUnknown *NsPlatNicType
var NsPlatNicTypeNicTypeTp *NsPlatNicType
var NsPlatNicTypeNicTypeSfp *NsPlatNicType

func init() {
	pNsPlatNicTypeNicTypeUnknown = cNsPlatNicTypeNicTypeUnknown
	NsPlatNicTypeNicTypeUnknown = &pNsPlatNicTypeNicTypeUnknown

	pNsPlatNicTypeNicTypeTp = cNsPlatNicTypeNicTypeTp
	NsPlatNicTypeNicTypeTp = &pNsPlatNicTypeNicTypeTp

	pNsPlatNicTypeNicTypeSfp = cNsPlatNicTypeNicTypeSfp
	NsPlatNicTypeNicTypeSfp = &pNsPlatNicTypeNicTypeSfp

}
