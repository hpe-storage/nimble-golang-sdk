// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// FibreChannelConfigFields provides field names to use in filter parameters, for example.
var FibreChannelConfigFields *FibreChannelConfigFieldHandles

func init() {
	FibreChannelConfigFields = &FibreChannelConfigFieldHandles{
		ID:               "id",
		ArrayList:        "array_list",
		GroupLeaderArray: "group_leader_array",
	}
}

// FibreChannelConfig - Manage group wide Fibre Channel configuration.
type FibreChannelConfig struct {
	// ID - Identifier for Fibre Channel configuration.
	ID *string `json:"id,omitempty"`
	// ArrayList - List of array Fibre Channel configs.
	ArrayList []*NsArrayFcConfig `json:"array_list,omitempty"`
	// GroupLeaderArray - Name of the group leader array.
	GroupLeaderArray *string `json:"group_leader_array,omitempty"`
}

// FibreChannelConfigFieldHandles provides a string representation for each FibreChannelConfig field.
type FibreChannelConfigFieldHandles struct {
	ID               string
	ArrayList        string
	GroupLeaderArray string
}
