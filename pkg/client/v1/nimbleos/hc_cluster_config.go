// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export HcClusterConfigFields provides field names to use in filter parameters, for example.
var HcClusterConfigFields *HcClusterConfigFieldHandles

func init() {
	fieldID := "id"
	fieldUniqueId := "unique_id"
	fieldName := "name"
	fieldDescription := "description"
	fieldUsername := "username"
	fieldPassword := "password"
	fieldType := "type"
	fieldMetadata := "metadata"
	fieldCreationTime := "creation_time"
	fieldLastModified := "last_modified"

	HcClusterConfigFields = &HcClusterConfigFieldHandles{
		ID:           &fieldID,
		UniqueId:     &fieldUniqueId,
		Name:         &fieldName,
		Description:  &fieldDescription,
		Username:     &fieldUsername,
		Password:     &fieldPassword,
		Type:         &fieldType,
		Metadata:     &fieldMetadata,
		CreationTime: &fieldCreationTime,
		LastModified: &fieldLastModified,
	}
}

// HcClusterConfig - Configuration information for virtual appliance that provides highly available storage and compute.
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

// HcClusterConfigFieldHandles provides a string representation for each AccessControlRecord field.
type HcClusterConfigFieldHandles struct {
	ID           *string
	UniqueId     *string
	Name         *string
	Description  *string
	Username     *string
	Password     *string
	Type         *string
	Metadata     *string
	CreationTime *string
	LastModified *string
}
