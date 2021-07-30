// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsFolderCreateAttrFields provides field names to use in filter parameters, for example.
var NsFolderCreateAttrFields *NsFolderCreateAttrFieldHandles

func init() {
	fieldName := "name"
	fieldPoolId := "pool_id"

	NsFolderCreateAttrFields = &NsFolderCreateAttrFieldHandles{
		Name:   &fieldName,
		PoolId: &fieldPoolId,
	}
}

// NsFolderCreateAttr - Attributes for folder creation.
type NsFolderCreateAttr struct {
	// Name - Name of folder.
	Name *string `json:"name,omitempty"`
	// PoolId - ID of pool to create the folder in.
	PoolId *string `json:"pool_id,omitempty"`
}

// NsFolderCreateAttrFieldHandles provides a string representation for each AccessControlRecord field.
type NsFolderCreateAttrFieldHandles struct {
	Name   *string
	PoolId *string
}
