// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsDiskRemoveType Enum.

type NsDiskRemoveType string

const (
	cNsDiskRemoveTypeTempFail   NsDiskRemoveType = "temp_fail"
	cNsDiskRemoveTypeLdrRemove  NsDiskRemoveType = "ldr_remove"
	cNsDiskRemoveTypePermFail   NsDiskRemoveType = "perm_fail"
	cNsDiskRemoveTypeRemove     NsDiskRemoveType = "remove"
	cNsDiskRemoveTypeDiagRemove NsDiskRemoveType = "diag_remove"
)

var pNsDiskRemoveTypeTempFail NsDiskRemoveType
var pNsDiskRemoveTypeLdrRemove NsDiskRemoveType
var pNsDiskRemoveTypePermFail NsDiskRemoveType
var pNsDiskRemoveTypeRemove NsDiskRemoveType
var pNsDiskRemoveTypeDiagRemove NsDiskRemoveType

// Export
var NsDiskRemoveTypeTempFail *NsDiskRemoveType
var NsDiskRemoveTypeLdrRemove *NsDiskRemoveType
var NsDiskRemoveTypePermFail *NsDiskRemoveType
var NsDiskRemoveTypeRemove *NsDiskRemoveType
var NsDiskRemoveTypeDiagRemove *NsDiskRemoveType

func init() {
	pNsDiskRemoveTypeTempFail = cNsDiskRemoveTypeTempFail
	NsDiskRemoveTypeTempFail = &pNsDiskRemoveTypeTempFail

	pNsDiskRemoveTypeLdrRemove = cNsDiskRemoveTypeLdrRemove
	NsDiskRemoveTypeLdrRemove = &pNsDiskRemoveTypeLdrRemove

	pNsDiskRemoveTypePermFail = cNsDiskRemoveTypePermFail
	NsDiskRemoveTypePermFail = &pNsDiskRemoveTypePermFail

	pNsDiskRemoveTypeRemove = cNsDiskRemoveTypeRemove
	NsDiskRemoveTypeRemove = &pNsDiskRemoveTypeRemove

	pNsDiskRemoveTypeDiagRemove = cNsDiskRemoveTypeDiagRemove
	NsDiskRemoveTypeDiagRemove = &pNsDiskRemoveTypeDiagRemove

}
