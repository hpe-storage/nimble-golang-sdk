// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsArgMap - Just a string to string map.
// Export NsArgMapFields for advance operations like search filter etc.
var NsArgMapFields *NsArgMap

func init() {

	NsArgMapFields = &NsArgMap{}
}

type NsArgMap struct {
}

// sdk internal struct
type nsArgMap struct {
}

// EncodeNsArgMap - Transform NsArgMap to nsArgMap type
func EncodeNsArgMap(request interface{}) (*nsArgMap, error) {
	reqNsArgMap := request.(*NsArgMap)
	byte, err := json.Marshal(reqNsArgMap)
	objPtr := &nsArgMap{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsArgMap Transform nsArgMap to NsArgMap type
func DecodeNsArgMap(request interface{}) (*NsArgMap, error) {
	reqNsArgMap := request.(*nsArgMap)
	byte, err := json.Marshal(reqNsArgMap)
	obj := &NsArgMap{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
