// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsObjectNameListWithType - List of objects of a given type.
// Export NsObjectNameListWithTypeFields for advance operations like search filter etc.
var NsObjectNameListWithTypeFields *NsObjectNameListWithType

func init() {

	NsObjectNameListWithTypeFields = &NsObjectNameListWithType{}
}

type NsObjectNameListWithType struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ObjNameList - List of object names.
	ObjNameList []*string `json:"obj_name_list,omitempty"`
}

// sdk internal struct
type nsObjectNameListWithType struct {
	ObjType     *NsObjectType `json:"obj_type,omitempty"`
	ObjNameList []*string     `json:"obj_name_list,omitempty"`
}

// EncodeNsObjectNameListWithType - Transform NsObjectNameListWithType to nsObjectNameListWithType type
func EncodeNsObjectNameListWithType(request interface{}) (*nsObjectNameListWithType, error) {
	reqNsObjectNameListWithType := request.(*NsObjectNameListWithType)
	byte, err := json.Marshal(reqNsObjectNameListWithType)
	if err != nil {
		return nil, err
	}
	respNsObjectNameListWithTypePtr := &nsObjectNameListWithType{}
	err = json.Unmarshal(byte, respNsObjectNameListWithTypePtr)
	return respNsObjectNameListWithTypePtr, err
}

// DecodeNsObjectNameListWithType Transform nsObjectNameListWithType to NsObjectNameListWithType type
func DecodeNsObjectNameListWithType(request interface{}) (*NsObjectNameListWithType, error) {
	reqNsObjectNameListWithType := request.(*nsObjectNameListWithType)
	byte, err := json.Marshal(reqNsObjectNameListWithType)
	if err != nil {
		return nil, err
	}
	respNsObjectNameListWithType := &NsObjectNameListWithType{}
	err = json.Unmarshal(byte, respNsObjectNameListWithType)
	return respNsObjectNameListWithType, err
}
