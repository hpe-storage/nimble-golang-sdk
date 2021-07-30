// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsDiscoveredGroupListReturnFields provides field names to use in filter parameters, for example.
var NsDiscoveredGroupListReturnFields *NsDiscoveredGroupListReturnFieldHandles

func init() {
	NsDiscoveredGroupListReturnFields = &NsDiscoveredGroupListReturnFieldHandles{
		DiscoveredGroupList: "discovered_group_list",
	}
}

// NsDiscoveredGroupListReturn - Detailed discovered group information.
type NsDiscoveredGroupListReturn struct {
	// DiscoveredGroupList - List of discovered group details.
	DiscoveredGroupList []*NsDiscoveredGroupInfo `json:"discovered_group_list,omitempty"`
}

// NsDiscoveredGroupListReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsDiscoveredGroupListReturnFieldHandles struct {
	DiscoveredGroupList string
}
