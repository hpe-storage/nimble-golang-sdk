// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcConfigRegenerateReturn - Return values of Fibre channel config regeneration.
// Export NsFcConfigRegenerateReturnFields for advance operations like search filter etc.
var NsFcConfigRegenerateReturnFields *NsFcConfigRegenerateReturn

func init() {
	GroupLeaderArrayfield := "group_leader_array"
	IDfield := "id"

	NsFcConfigRegenerateReturnFields = &NsFcConfigRegenerateReturn{
		GroupLeaderArray: &GroupLeaderArrayfield,
		ID:               &IDfield,
	}
}

type NsFcConfigRegenerateReturn struct {
	// ArrayList - List of array Fibre Channel configs.
	ArrayList []*NsArrayFcConfig `json:"array_list,omitempty"`
	// GroupLeaderArray - Name of the group leader array.
	GroupLeaderArray *string `json:"group_leader_array,omitempty"`
	// ID - Identifier for Fibre Channel configuration.
	ID *string `json:"id,omitempty"`
}
