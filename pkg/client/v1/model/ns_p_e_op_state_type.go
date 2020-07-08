// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsPEOpStateType Enum.

type NsPEOpStateType string

const (
	cNsPEOpStateTypeNormal     NsPEOpStateType = "normal"
	cNsPEOpStateTypeDeprecated NsPEOpStateType = "deprecated"
)

var pNsPEOpStateTypeNormal NsPEOpStateType
var pNsPEOpStateTypeDeprecated NsPEOpStateType

// Export
var NsPEOpStateTypeNormal *NsPEOpStateType
var NsPEOpStateTypeDeprecated *NsPEOpStateType

func init() {
	pNsPEOpStateTypeNormal = cNsPEOpStateTypeNormal
	NsPEOpStateTypeNormal = &pNsPEOpStateTypeNormal

	pNsPEOpStateTypeDeprecated = cNsPEOpStateTypeDeprecated
	NsPEOpStateTypeDeprecated = &pNsPEOpStateTypeDeprecated

}
