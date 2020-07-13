// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsHostType - Host type attribute.
// Export NsHostTypeFields for advance operations like search filter etc.
var NsHostTypeFields *NsHostType

func init() {

	NsHostTypeFields = &NsHostType{
		InitiatorName:       "initiator_name",
		SourceHostType:      "source_host_type",
		DestinationHostType: "destination_host_type",
	}
}

type NsHostType struct {
	// InitiatorName - The initiator for which the conflict exists.
	InitiatorName string `json:"initiator_name,omitempty"`
	// SourceInitiatorGroup - Source initiator group.
	SourceInitiatorGroup []*string `json:"source_initiator_group,omitempty"`
	// SourceHostType - Source host type.
	SourceHostType string `json:"source_host_type,omitempty"`
	// DestinationInitiatorGroup - Destination initiator group.
	DestinationInitiatorGroup []*string `json:"destination_initiator_group,omitempty"`
	// DestinationHostType - Destination Host type.
	DestinationHostType string `json:"destination_host_type,omitempty"`
}

// sdk internal struct
type nsHostType struct {
	// InitiatorName - The initiator for which the conflict exists.
	InitiatorName *string `json:"initiator_name,omitempty"`
	// SourceInitiatorGroup - Source initiator group.
	SourceInitiatorGroup []*string `json:"source_initiator_group,omitempty"`
	// SourceHostType - Source host type.
	SourceHostType *string `json:"source_host_type,omitempty"`
	// DestinationInitiatorGroup - Destination initiator group.
	DestinationInitiatorGroup []*string `json:"destination_initiator_group,omitempty"`
	// DestinationHostType - Destination Host type.
	DestinationHostType *string `json:"destination_host_type,omitempty"`
}

// EncodeNsHostType - Transform NsHostType to nsHostType type
func EncodeNsHostType(request interface{}) (*nsHostType, error) {
	reqNsHostType := request.(*NsHostType)
	byte, err := json.Marshal(reqNsHostType)
	objPtr := &nsHostType{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsHostType Transform nsHostType to NsHostType type
func DecodeNsHostType(request interface{}) (*NsHostType, error) {
	reqNsHostType := request.(*nsHostType)
	byte, err := json.Marshal(reqNsHostType)
	obj := &NsHostType{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
