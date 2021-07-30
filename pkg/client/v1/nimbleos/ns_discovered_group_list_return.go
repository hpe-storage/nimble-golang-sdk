// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsDiscoveredGroupListReturn - Detailed discovered group information.

// Export NsDiscoveredGroupListReturnFields provides field names to use in filter parameters, for example.
var NsDiscoveredGroupListReturnFields *NsDiscoveredGroupListReturnStringFields

func init() {
	fieldDiscoveredGroupList := "discovered_group_list"

	NsDiscoveredGroupListReturnFields = &NsDiscoveredGroupListReturnStringFields{
		DiscoveredGroupList: &fieldDiscoveredGroupList,
	}
}

type NsDiscoveredGroupListReturn struct {
	// DiscoveredGroupList - List of discovered group details.
	DiscoveredGroupList []*NsDiscoveredGroupInfo `json:"discovered_group_list,omitempty"`
}

// Struct for NsDiscoveredGroupListReturnFields
type NsDiscoveredGroupListReturnStringFields struct {
	DiscoveredGroupList *string
}
