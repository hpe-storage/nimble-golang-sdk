// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsFcConfigRegenerateReturn - Return values of Fibre channel config regeneration.
// Export NsFcConfigRegenerateReturnFields for advance operations like search filter etc.
var NsFcConfigRegenerateReturnFields *NsFcConfigRegenerateReturn

func init() {

	NsFcConfigRegenerateReturnFields = &NsFcConfigRegenerateReturn{
		GroupLeaderArray: "group_leader_array",
		ID:               "id",
	}
}

type NsFcConfigRegenerateReturn struct {
	// ArrayList - List of array Fibre Channel configs.
	ArrayList []*NsArrayFcConfig `json:"array_list,omitempty"`
	// GroupLeaderArray - Name of the group leader array.
	GroupLeaderArray string `json:"group_leader_array,omitempty"`
	// ID - Identifier for Fibre Channel configuration.
	ID string `json:"id,omitempty"`
}

// sdk internal struct
type nsFcConfigRegenerateReturn struct {
	ArrayList        []*NsArrayFcConfig `json:"array_list,omitempty"`
	GroupLeaderArray *string            `json:"group_leader_array,omitempty"`
	ID               *string            `json:"id,omitempty"`
}

// EncodeNsFcConfigRegenerateReturn - Transform NsFcConfigRegenerateReturn to nsFcConfigRegenerateReturn type
func EncodeNsFcConfigRegenerateReturn(request interface{}) (*nsFcConfigRegenerateReturn, error) {
	reqNsFcConfigRegenerateReturn := request.(*NsFcConfigRegenerateReturn)
	byte, err := json.Marshal(reqNsFcConfigRegenerateReturn)
	if err != nil {
		return nil, err
	}
	respNsFcConfigRegenerateReturnPtr := &nsFcConfigRegenerateReturn{}
	err = json.Unmarshal(byte, respNsFcConfigRegenerateReturnPtr)
	return respNsFcConfigRegenerateReturnPtr, err
}

// DecodeNsFcConfigRegenerateReturn Transform nsFcConfigRegenerateReturn to NsFcConfigRegenerateReturn type
func DecodeNsFcConfigRegenerateReturn(request interface{}) (*NsFcConfigRegenerateReturn, error) {
	reqNsFcConfigRegenerateReturn := request.(*nsFcConfigRegenerateReturn)
	byte, err := json.Marshal(reqNsFcConfigRegenerateReturn)
	if err != nil {
		return nil, err
	}
	respNsFcConfigRegenerateReturn := &NsFcConfigRegenerateReturn{}
	err = json.Unmarshal(byte, respNsFcConfigRegenerateReturn)
	return respNsFcConfigRegenerateReturn, err
}
