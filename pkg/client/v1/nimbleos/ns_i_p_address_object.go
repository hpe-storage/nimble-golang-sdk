// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsIPAddressObject - Object wrapper of IP Address.
// Export NsIPAddressObjectFields for advance operations like search filter etc.
var NsIPAddressObjectFields *NsIPAddressObject

func init() {

	NsIPAddressObjectFields = &NsIPAddressObject{
		IpAddr: "ip_addr",
	}
}

type NsIPAddressObject struct {
	// IpAddr - An IP Address.
	IpAddr string `json:"ip_addr,omitempty"`
}

// sdk internal struct
type nsIPAddressObject struct {
	IpAddr *string `json:"ip_addr,omitempty"`
}

// EncodeNsIPAddressObject - Transform NsIPAddressObject to nsIPAddressObject type
func EncodeNsIPAddressObject(request interface{}) (*nsIPAddressObject, error) {
	reqNsIPAddressObject := request.(*NsIPAddressObject)
	byte, err := json.Marshal(reqNsIPAddressObject)
	if err != nil {
		return nil, err
	}
	respNsIPAddressObjectPtr := &nsIPAddressObject{}
	err = json.Unmarshal(byte, respNsIPAddressObjectPtr)
	return respNsIPAddressObjectPtr, err
}

// DecodeNsIPAddressObject Transform nsIPAddressObject to NsIPAddressObject type
func DecodeNsIPAddressObject(request interface{}) (*NsIPAddressObject, error) {
	reqNsIPAddressObject := request.(*nsIPAddressObject)
	byte, err := json.Marshal(reqNsIPAddressObject)
	if err != nil {
		return nil, err
	}
	respNsIPAddressObject := &NsIPAddressObject{}
	err = json.Unmarshal(byte, respNsIPAddressObject)
	return respNsIPAddressObject, err
}
