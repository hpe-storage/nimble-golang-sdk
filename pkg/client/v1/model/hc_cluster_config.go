// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// HcClusterConfig - Configuration information for virtual appliance that provides highly available storage and compute.
// Export HcClusterConfigFields for advance operations like search filter etc.
var HcClusterConfigFields *HcClusterConfig

func init() {

	HcClusterConfigFields = &HcClusterConfig{
		ID:          "id",
		UniqueId:    "unique_id",
		Name:        "name",
		Description: "description",
		Username:    "username",
		Password:    "password",
	}
}

type HcClusterConfig struct {
	// ID - Identifier for the hc cluster config.
	ID string `json:"id,omitempty"`
	// UniqueId - Unique identifier for the HC component.
	UniqueId string `json:"unique_id,omitempty"`
	// Name - Name for the HC component.
	Name string `json:"name,omitempty"`
	// Description - Text description of HC component.
	Description string `json:"description,omitempty"`
	// Username - HC component username.
	Username string `json:"username,omitempty"`
	// Password - HC component password.
	Password string `json:"password,omitempty"`
	// Type - HCI config type ({invalid|node|block|cluster}).
	Type *NsHCIConfigType `json:"type,omitempty"`
	// Metadata - Key-value pairs that augment a HC cluster config's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// CreationTime - Time when this HC cluster configuration was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this HC cluster configuration was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
}

// sdk internal struct
type hcClusterConfig struct {
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
	// Type - HCI config type ({invalid|node|block|cluster}).
	Type *NsHCIConfigType `json:"type,omitempty"`
	// Metadata - Key-value pairs that augment a HC cluster config's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// CreationTime - Time when this HC cluster configuration was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this HC cluster configuration was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
}

// EncodeHcClusterConfig - Transform HcClusterConfig to hcClusterConfig type
func EncodeHcClusterConfig(request interface{}) (*hcClusterConfig, error) {
	reqHcClusterConfig := request.(*HcClusterConfig)
	byte, err := json.Marshal(reqHcClusterConfig)
	objPtr := &hcClusterConfig{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeHcClusterConfig Transform hcClusterConfig to HcClusterConfig type
func DecodeHcClusterConfig(request interface{}) (*HcClusterConfig, error) {
	reqHcClusterConfig := request.(*hcClusterConfig)
	byte, err := json.Marshal(reqHcClusterConfig)
	obj := &HcClusterConfig{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
