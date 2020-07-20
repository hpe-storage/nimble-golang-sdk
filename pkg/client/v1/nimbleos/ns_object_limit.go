// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsObjectLimit - Limits (Max total number of objects) for object of a given type.
// Export NsObjectLimitFields for advance operations like search filter etc.
var NsObjectLimitFields *NsObjectLimit

func init() {

	NsObjectLimitFields = &NsObjectLimit{}
}

type NsObjectLimit struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ObjLimit - Limit of the objects.
	ObjLimit int64 `json:"obj_limit,omitempty"`
	// ObjNum - Number of objects after group merge.
	ObjNum int64 `json:"obj_num,omitempty"`
}

// sdk internal struct
type nsObjectLimit struct {
	ObjType  *NsObjectType `json:"obj_type,omitempty"`
	ObjLimit *int64        `json:"obj_limit,omitempty"`
	ObjNum   *int64        `json:"obj_num,omitempty"`
}

// EncodeNsObjectLimit - Transform NsObjectLimit to nsObjectLimit type
func EncodeNsObjectLimit(request interface{}) (*nsObjectLimit, error) {
	reqNsObjectLimit := request.(*NsObjectLimit)
	byte, err := json.Marshal(reqNsObjectLimit)
	if err != nil {
		return nil, err
	}
	respNsObjectLimitPtr := &nsObjectLimit{}
	err = json.Unmarshal(byte, respNsObjectLimitPtr)
	return respNsObjectLimitPtr, err
}

// DecodeNsObjectLimit Transform nsObjectLimit to NsObjectLimit type
func DecodeNsObjectLimit(request interface{}) (*NsObjectLimit, error) {
	reqNsObjectLimit := request.(*nsObjectLimit)
	byte, err := json.Marshal(reqNsObjectLimit)
	if err != nil {
		return nil, err
	}
	respNsObjectLimit := &NsObjectLimit{}
	err = json.Unmarshal(byte, respNsObjectLimit)
	return respNsObjectLimit, err
}
