// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsNIC - Network interface controller.
// Export NsNICFields for advance operations like search filter etc.
var NsNICFields *NsNIC

func init() {

	NsNICFields = &NsNIC{
		Name:        "name",
		SubnetLabel: "subnet_label",
		DataIp:      "data_ip",
	}
}

type NsNIC struct {
	// Name - Name of NIC.
	Name string `json:"name,omitempty"`
	// SubnetLabel - Subnet label for this NIC.
	SubnetLabel string `json:"subnet_label,omitempty"`
	// DataIp - Data IP address.
	DataIp string `json:"data_ip,omitempty"`
	// Tagged - Identify whether the NIC is tagged.
	Tagged *bool `json:"tagged,omitempty"`
}

// sdk internal struct
type nsNIC struct {
	// Name - Name of NIC.
	Name *string `json:"name,omitempty"`
	// SubnetLabel - Subnet label for this NIC.
	SubnetLabel *string `json:"subnet_label,omitempty"`
	// DataIp - Data IP address.
	DataIp *string `json:"data_ip,omitempty"`
	// Tagged - Identify whether the NIC is tagged.
	Tagged *bool `json:"tagged,omitempty"`
}

// EncodeNsNIC - Transform NsNIC to nsNIC type
func EncodeNsNIC(request interface{}) (*nsNIC, error) {
	reqNsNIC := request.(*NsNIC)
	byte, err := json.Marshal(reqNsNIC)
	objPtr := &nsNIC{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsNIC Transform nsNIC to NsNIC type
func DecodeNsNIC(request interface{}) (*NsNIC, error) {
	reqNsNIC := request.(*nsNIC)
	byte, err := json.Marshal(reqNsNIC)
	obj := &NsNIC{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
