// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Autosupport - Get status of autosupport.
// Export AutosupportFields for advance operations like search filter etc.
var AutosupportFields *Autosupport

func init() {

	AutosupportFields = &Autosupport{
		ID:        "id",
		GroupId:   "group_id",
		GroupName: "group_name",
	}
}

type Autosupport struct {
	// ID - Identifier of the autosupport.
	ID string `json:"id,omitempty"`
	// ArrayList - List of arrays in the group with autosupport information.
	ArrayList []*NsArrayAsupDetail `json:"array_list,omitempty"`
	// ArrayCount - Number of arrays in the group.
	ArrayCount int64 `json:"array_count,omitempty"`
	// GroupId - Identifier for the group.
	GroupId string `json:"group_id,omitempty"`
	// GroupName - Name of the group.
	GroupName string `json:"group_name,omitempty"`
}

// sdk internal struct
type autosupport struct {
	ID         *string              `json:"id,omitempty"`
	ArrayList  []*NsArrayAsupDetail `json:"array_list,omitempty"`
	ArrayCount *int64               `json:"array_count,omitempty"`
	GroupId    *string              `json:"group_id,omitempty"`
	GroupName  *string              `json:"group_name,omitempty"`
}

// EncodeAutosupport - Transform Autosupport to autosupport type
func EncodeAutosupport(request interface{}) (*autosupport, error) {
	reqAutosupport := request.(*Autosupport)
	byte, err := json.Marshal(reqAutosupport)
	if err != nil {
		return nil, err
	}
	respAutosupportPtr := &autosupport{}
	err = json.Unmarshal(byte, respAutosupportPtr)
	return respAutosupportPtr, err
}

// DecodeAutosupport Transform autosupport to Autosupport type
func DecodeAutosupport(request interface{}) (*Autosupport, error) {
	reqAutosupport := request.(*autosupport)
	byte, err := json.Marshal(reqAutosupport)
	if err != nil {
		return nil, err
	}
	respAutosupport := &Autosupport{}
	err = json.Unmarshal(byte, respAutosupport)
	return respAutosupport, err
}
