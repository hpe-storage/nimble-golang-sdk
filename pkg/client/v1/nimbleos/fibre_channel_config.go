// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// FibreChannelConfig - Manage group wide Fibre Channel configuration.
// Export FibreChannelConfigFields for advance operations like search filter etc.
var FibreChannelConfigFields *FibreChannelConfig

func init() {

	FibreChannelConfigFields = &FibreChannelConfig{
		ID:               "id",
		GroupLeaderArray: "group_leader_array",
	}
}

type FibreChannelConfig struct {
	// ID - Identifier for Fibre Channel configuration.
	ID string `json:"id,omitempty"`
	// ArrayList - List of array Fibre Channel configs.
	ArrayList []*NsArrayFcConfig `json:"array_list,omitempty"`
	// GroupLeaderArray - Name of the group leader array.
	GroupLeaderArray string `json:"group_leader_array,omitempty"`
}

// sdk internal struct
type fibreChannelConfig struct {
	// ID - Identifier for Fibre Channel configuration.
	ID *string `json:"id,omitempty"`
	// ArrayList - List of array Fibre Channel configs.
	ArrayList []*NsArrayFcConfig `json:"array_list,omitempty"`
	// GroupLeaderArray - Name of the group leader array.
	GroupLeaderArray *string `json:"group_leader_array,omitempty"`
}

// EncodeFibreChannelConfig - Transform FibreChannelConfig to fibreChannelConfig type
func EncodeFibreChannelConfig(request interface{}) (*fibreChannelConfig, error) {
	reqFibreChannelConfig := request.(*FibreChannelConfig)
	byte, err := json.Marshal(reqFibreChannelConfig)
	objPtr := &fibreChannelConfig{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeFibreChannelConfig Transform fibreChannelConfig to FibreChannelConfig type
func DecodeFibreChannelConfig(request interface{}) (*FibreChannelConfig, error) {
	reqFibreChannelConfig := request.(*fibreChannelConfig)
	byte, err := json.Marshal(reqFibreChannelConfig)
	obj := &FibreChannelConfig{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
