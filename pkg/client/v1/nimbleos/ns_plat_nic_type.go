// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

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

// NsPlatNicTypeNicTypeUnknown enum exports
var NsPlatNicTypeNicTypeUnknown *NsPlatNicType

// NsPlatNicTypeNicTypeTp enum exports
var NsPlatNicTypeNicTypeTp *NsPlatNicType

// NsPlatNicTypeNicTypeSfp enum exports
var NsPlatNicTypeNicTypeSfp *NsPlatNicType

func init() {
	pNsPlatNicTypeNicTypeUnknown = cNsPlatNicTypeNicTypeUnknown
	NsPlatNicTypeNicTypeUnknown = &pNsPlatNicTypeNicTypeUnknown

	pNsPlatNicTypeNicTypeTp = cNsPlatNicTypeNicTypeTp
	NsPlatNicTypeNicTypeTp = &pNsPlatNicTypeNicTypeTp

	pNsPlatNicTypeNicTypeSfp = cNsPlatNicTypeNicTypeSfp
	NsPlatNicTypeNicTypeSfp = &pNsPlatNicTypeNicTypeSfp

}
