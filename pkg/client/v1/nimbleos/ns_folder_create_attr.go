// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFolderCreateAttr - Attributes for folder creation.
// Export NsFolderCreateAttrFields for advance operations like search filter etc.
var NsFolderCreateAttrFields *NsFolderCreateAttrStringFields

func init() {
	Namefield := "name"
	PoolIdfield := "pool_id"

	NsFolderCreateAttrFields = &NsFolderCreateAttrStringFields{
		Name:   &Namefield,
		PoolId: &PoolIdfield,
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
