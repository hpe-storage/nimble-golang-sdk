// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsPEOpStateType Enum.

type NsPEOpStateType string

const (
	cNsPEOpStateTypeNormal     NsPEOpStateType = "normal"
	cNsPEOpStateTypeDeprecated NsPEOpStateType = "deprecated"
)

var pNsPEOpStateTypeNormal NsPEOpStateType
var pNsPEOpStateTypeDeprecated NsPEOpStateType

// NsPEOpStateTypeNormal enum exports
var NsPEOpStateTypeNormal *NsPEOpStateType

// NsPEOpStateTypeDeprecated enum exports
var NsPEOpStateTypeDeprecated *NsPEOpStateType

func init() {
	pNsPEOpStateTypeNormal = cNsPEOpStateTypeNormal
	NsPEOpStateTypeNormal = &pNsPEOpStateTypeNormal

	pNsPEOpStateTypeDeprecated = cNsPEOpStateTypeDeprecated
	NsPEOpStateTypeDeprecated = &pNsPEOpStateTypeDeprecated

}
