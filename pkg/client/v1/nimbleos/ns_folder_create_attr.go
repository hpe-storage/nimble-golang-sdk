// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFolderCreateAttr - Attributes for folder creation.

// Export NsFolderCreateAttrFields provides field names to use in filter parameters, for example.
var NsFolderCreateAttrFields *NsFolderCreateAttrStringFields

func init() {
	fieldName := "name"
	fieldPoolId := "pool_id"

	NsFolderCreateAttrFields = &NsFolderCreateAttrStringFields{
		Name:   &fieldName,
		PoolId: &fieldPoolId,
	}
}

type NsFolderCreateAttr struct {
	// Name - Name of folder.
	Name *string `json:"name,omitempty"`
	// PoolId - ID of pool to create the folder in.
	PoolId *string `json:"pool_id,omitempty"`
}

// Struct for NsFolderCreateAttrFields
type NsFolderCreateAttrStringFields struct {
	Name   *string
	PoolId *string
}
