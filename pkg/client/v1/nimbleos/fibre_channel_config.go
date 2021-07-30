// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// FibreChannelConfig - Manage group wide Fibre Channel configuration.
// Export FibreChannelConfigFields for advance operations like search filter etc.
var FibreChannelConfigFields *FibreChannelConfigStringFields

func init() {
	IDfield := "id"
	ArrayListfield := "array_list"
	GroupLeaderArrayfield := "group_leader_array"

	FibreChannelConfigFields = &FibreChannelConfigStringFields{
		ID:               &IDfield,
		ArrayList:        &ArrayListfield,
		GroupLeaderArray: &GroupLeaderArrayfield,
	}
}

type FibreChannelConfig struct {
	// ID - Identifier for Fibre Channel configuration.
	ID *string `json:"id,omitempty"`
	// ArrayList - List of array Fibre Channel configs.
	ArrayList []*NsArrayFcConfig `json:"array_list,omitempty"`
	// GroupLeaderArray - Name of the group leader array.
	GroupLeaderArray *string `json:"group_leader_array,omitempty"`
}

// Struct for FibreChannelConfigFields
type FibreChannelConfigStringFields struct {
	ID               *string
	ArrayList        *string
	GroupLeaderArray *string
}
