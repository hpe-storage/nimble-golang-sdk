// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// HcClusterConfig - Configuration information for virtual appliance that provides highly available storage and compute.
// Export HcClusterConfigFields for advance operations like search filter etc.
var HcClusterConfigFields *HcClusterConfig

func init() {
	IDfield := "id"
	UniqueIdfield := "unique_id"
	Namefield := "name"
	Descriptionfield := "description"
	Usernamefield := "username"
	Passwordfield := "password"

	HcClusterConfigFields = &HcClusterConfig{
		ID:          &IDfield,
		UniqueId:    &UniqueIdfield,
		Name:        &Namefield,
		Description: &Descriptionfield,
		Username:    &Usernamefield,
		Password:    &Passwordfield,
	}
}

type HcClusterConfig struct {
	// ID - Identifier for the hc cluster config.
	ID *string `json:"id,omitempty"`
	// UniqueId - Unique identifier for the HC component.
	UniqueId *string `json:"unique_id,omitempty"`
	// Name - Name for the HC component.
	Name *string `json:"name,omitempty"`
	// Description - Text description of HC component.
	Description *string `json:"description,omitempty"`
	// Username - HC component username.
	Username *string `json:"username,omitempty"`
	// Password - HC component password.
	Password *string `json:"password,omitempty"`
	// Type - HCI config type ({node|block|cluster|switch}).
	Type *NsHCIConfigType `json:"type,omitempty"`
	// Metadata - Key-value pairs that augment a HC cluster config's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// CreationTime - Time when this HC cluster configuration was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this HC cluster configuration was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
}
