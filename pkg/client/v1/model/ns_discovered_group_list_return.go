// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsDiscoveredGroupListReturn - Detailed discovered group information.
// Export NsDiscoveredGroupListReturnFields for advance operations like search filter etc.
var NsDiscoveredGroupListReturnFields *NsDiscoveredGroupListReturn

func init(){
		
	NsDiscoveredGroupListReturnFields= &NsDiscoveredGroupListReturn{
		
	}
}

type NsDiscoveredGroupListReturn struct {
	// DiscoveredGroupList - List of discovered group details.
   	DiscoveredGroupList []*NsDiscoveredGroupInfo `json:"discovered_group_list,omitempty"`
}
