// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsReplPairListWithObjectType - Replicated objects of the specified type.
// Export NsReplPairListWithObjectTypeFields for advance operations like search filter etc.
var NsReplPairListWithObjectTypeFields *NsReplPairListWithObjectType

func init() {

	NsReplPairListWithObjectTypeFields = &NsReplPairListWithObjectType{}
}

type NsReplPairListWithObjectType struct {
	// ObjType - Type of the replicated object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ReplList - List of replicated objects of this type.
	ReplList []*NsReplPair `json:"repl_list,omitempty"`
}

// sdk internal struct
type nsReplPairListWithObjectType struct {
	ObjType  *NsObjectType `json:"obj_type,omitempty"`
	ReplList []*NsReplPair `json:"repl_list,omitempty"`
}

// EncodeNsReplPairListWithObjectType - Transform NsReplPairListWithObjectType to nsReplPairListWithObjectType type
func EncodeNsReplPairListWithObjectType(request interface{}) (*nsReplPairListWithObjectType, error) {
	reqNsReplPairListWithObjectType := request.(*NsReplPairListWithObjectType)
	byte, err := json.Marshal(reqNsReplPairListWithObjectType)
	if err != nil {
		return nil, err
	}
	respNsReplPairListWithObjectTypePtr := &nsReplPairListWithObjectType{}
	err = json.Unmarshal(byte, respNsReplPairListWithObjectTypePtr)
	return respNsReplPairListWithObjectTypePtr, err
}

// DecodeNsReplPairListWithObjectType Transform nsReplPairListWithObjectType to NsReplPairListWithObjectType type
func DecodeNsReplPairListWithObjectType(request interface{}) (*NsReplPairListWithObjectType, error) {
	reqNsReplPairListWithObjectType := request.(*nsReplPairListWithObjectType)
	byte, err := json.Marshal(reqNsReplPairListWithObjectType)
	if err != nil {
		return nil, err
	}
	respNsReplPairListWithObjectType := &NsReplPairListWithObjectType{}
	err = json.Unmarshal(byte, respNsReplPairListWithObjectType)
	return respNsReplPairListWithObjectType, err
}
