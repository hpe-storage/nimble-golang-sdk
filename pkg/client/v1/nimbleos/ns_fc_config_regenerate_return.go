// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsFcConfigRegenerateReturnFields provides field names to use in filter parameters, for example.
var NsFcConfigRegenerateReturnFields *NsFcConfigRegenerateReturnFieldHandles

func init() {
	fieldArrayList := "array_list"
	fieldGroupLeaderArray := "group_leader_array"
	fieldID := "id"

	NsFcConfigRegenerateReturnFields = &NsFcConfigRegenerateReturnFieldHandles{
		ArrayList:        &fieldArrayList,
		GroupLeaderArray: &fieldGroupLeaderArray,
		ID:               &fieldID,
	}
}

// NsFcConfigRegenerateReturn - Return values of Fibre channel config regeneration.
type NsFcConfigRegenerateReturn struct {
	// ArrayList - List of array Fibre Channel configs.
	ArrayList []*NsArrayFcConfig `json:"array_list,omitempty"`
	// GroupLeaderArray - Name of the group leader array.
	GroupLeaderArray *string `json:"group_leader_array,omitempty"`
	// ID - Identifier for Fibre Channel configuration.
	ID *string `json:"id,omitempty"`
}

// NsFcConfigRegenerateReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsFcConfigRegenerateReturnFieldHandles struct {
	ArrayList        *string
	GroupLeaderArray *string
	ID               *string
}
