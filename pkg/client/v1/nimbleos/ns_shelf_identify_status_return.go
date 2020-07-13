// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsShelfIdentifyStatusReturn - Status of the shelf identifier.
// Export NsShelfIdentifyStatusReturnFields for advance operations like search filter etc.
var NsShelfIdentifyStatusReturnFields *NsShelfIdentifyStatusReturn

func init() {

	NsShelfIdentifyStatusReturnFields = &NsShelfIdentifyStatusReturn{}
}

type NsShelfIdentifyStatusReturn struct {
	// Enabled - Shelf identifier is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// sdk internal struct
type nsShelfIdentifyStatusReturn struct {
	// Enabled - Shelf identifier is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// EncodeNsShelfIdentifyStatusReturn - Transform NsShelfIdentifyStatusReturn to nsShelfIdentifyStatusReturn type
func EncodeNsShelfIdentifyStatusReturn(request interface{}) (*nsShelfIdentifyStatusReturn, error) {
	reqNsShelfIdentifyStatusReturn := request.(*NsShelfIdentifyStatusReturn)
	byte, err := json.Marshal(reqNsShelfIdentifyStatusReturn)
	objPtr := &nsShelfIdentifyStatusReturn{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsShelfIdentifyStatusReturn Transform nsShelfIdentifyStatusReturn to NsShelfIdentifyStatusReturn type
func DecodeNsShelfIdentifyStatusReturn(request interface{}) (*NsShelfIdentifyStatusReturn, error) {
	reqNsShelfIdentifyStatusReturn := request.(*nsShelfIdentifyStatusReturn)
	byte, err := json.Marshal(reqNsShelfIdentifyStatusReturn)
	obj := &NsShelfIdentifyStatusReturn{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
