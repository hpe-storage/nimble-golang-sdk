// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFolderCreateAttr - Attributes for folder creation.
// Export NsFolderCreateAttrFields for advance operations like search filter etc.
var NsFolderCreateAttrFields *NsFolderCreateAttr

func init() {
	Namefield := "name"
	PoolIdfield := "pool_id"

	NsFolderCreateAttrFields = &NsFolderCreateAttr{
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
