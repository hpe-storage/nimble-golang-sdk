// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsDiscoveredGroupListReturn - Detailed discovered group information.
// Export NsDiscoveredGroupListReturnFields for advance operations like search filter etc.
var NsDiscoveredGroupListReturnFields *NsDiscoveredGroupListReturnStringFields

func init() {
	DiscoveredGroupListfield := "discovered_group_list"

	NsDiscoveredGroupListReturnFields = &NsDiscoveredGroupListReturnStringFields{
		DiscoveredGroupList: &DiscoveredGroupListfield,
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
