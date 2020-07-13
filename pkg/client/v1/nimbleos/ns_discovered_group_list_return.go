// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsDiscoveredGroupListReturn - Detailed discovered group information.
// Export NsDiscoveredGroupListReturnFields for advance operations like search filter etc.
var NsDiscoveredGroupListReturnFields *NsDiscoveredGroupListReturn

func init() {

	NsDiscoveredGroupListReturnFields = &NsDiscoveredGroupListReturn{}
}

type NsDiscoveredGroupListReturn struct {
	// DiscoveredGroupList - List of discovered group details.
	DiscoveredGroupList []*NsDiscoveredGroupInfo `json:"discovered_group_list,omitempty"`
}

// sdk internal struct
type nsDiscoveredGroupListReturn struct {
	// DiscoveredGroupList - List of discovered group details.
	DiscoveredGroupList []*NsDiscoveredGroupInfo `json:"discovered_group_list,omitempty"`
}

// EncodeNsDiscoveredGroupListReturn - Transform NsDiscoveredGroupListReturn to nsDiscoveredGroupListReturn type
func EncodeNsDiscoveredGroupListReturn(request interface{}) (*nsDiscoveredGroupListReturn, error) {
	reqNsDiscoveredGroupListReturn := request.(*NsDiscoveredGroupListReturn)
	byte, err := json.Marshal(reqNsDiscoveredGroupListReturn)
	objPtr := &nsDiscoveredGroupListReturn{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsDiscoveredGroupListReturn Transform nsDiscoveredGroupListReturn to NsDiscoveredGroupListReturn type
func DecodeNsDiscoveredGroupListReturn(request interface{}) (*NsDiscoveredGroupListReturn, error) {
	reqNsDiscoveredGroupListReturn := request.(*nsDiscoveredGroupListReturn)
	byte, err := json.Marshal(reqNsDiscoveredGroupListReturn)
	obj := &NsDiscoveredGroupListReturn{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
