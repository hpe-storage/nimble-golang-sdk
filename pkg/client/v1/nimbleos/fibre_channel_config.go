// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export FibreChannelConfigFields provides field names to use in filter parameters, for example.
var FibreChannelConfigFields *FibreChannelConfigFieldHandles

func init() {
	fieldID := "id"
	fieldArrayList := "array_list"
	fieldGroupLeaderArray := "group_leader_array"

	FibreChannelConfigFields = &FibreChannelConfigFieldHandles{
		ID:               &fieldID,
		ArrayList:        &fieldArrayList,
		GroupLeaderArray: &fieldGroupLeaderArray,
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

// FibreChannelConfigFieldHandles provides a string representation for each AccessControlRecord field.
type FibreChannelConfigFieldHandles struct {
	ID               *string
	ArrayList        *string
	GroupLeaderArray *string
}
