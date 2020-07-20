// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

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
	if err != nil {
		return nil, err
	}
	respNsArgMapPtr := &nsArgMap{}
	err = json.Unmarshal(byte, respNsArgMapPtr)
	return respNsArgMapPtr, err
}

// DecodeNsArgMap Transform nsArgMap to NsArgMap type
func DecodeNsArgMap(request interface{}) (*NsArgMap, error) {
	reqNsArgMap := request.(*nsArgMap)
	byte, err := json.Marshal(reqNsArgMap)
	if err != nil {
		return nil, err
	}
	respNsArgMap := &NsArgMap{}
	err = json.Unmarshal(byte, respNsArgMap)
	return respNsArgMap, err
}
