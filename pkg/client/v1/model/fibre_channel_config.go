// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// FibreChannelConfig - Manage group wide Fibre Channel configuration.
// Export FibreChannelConfigFields for advance operations like search filter etc.
var FibreChannelConfigFields *FibreChannelConfig

func init(){
	IDfield:= "id"
	GroupLeaderArrayfield:= "group_leader_array"
		
	FibreChannelConfigFields= &FibreChannelConfig{
	ID: &IDfield,
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
