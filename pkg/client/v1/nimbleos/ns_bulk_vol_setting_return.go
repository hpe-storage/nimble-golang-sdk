// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsBulkVolSettingReturn - Return codes for setting an attribute to a list of items.
// Export NsBulkVolSettingReturnFields for advance operations like search filter etc.
var NsBulkVolSettingReturnFields *NsBulkVolSettingReturn

func init() {

	NsBulkVolSettingReturnFields = &NsBulkVolSettingReturn{}
}

type NsBulkVolSettingReturn struct {
	// ErrorCodes - Error codes for every element in a list of items.
	ErrorCodes []*string `json:"error_codes,omitempty"`
}

// sdk internal struct
type nsBulkVolSettingReturn struct {
	// ErrorCodes - Error codes for every element in a list of items.
	ErrorCodes []*string `json:"error_codes,omitempty"`
}

// EncodeNsBulkVolSettingReturn - Transform NsBulkVolSettingReturn to nsBulkVolSettingReturn type
func EncodeNsBulkVolSettingReturn(request interface{}) (*nsBulkVolSettingReturn, error) {
	reqNsBulkVolSettingReturn := request.(*NsBulkVolSettingReturn)
	byte, err := json.Marshal(reqNsBulkVolSettingReturn)
	objPtr := &nsBulkVolSettingReturn{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsBulkVolSettingReturn Transform nsBulkVolSettingReturn to NsBulkVolSettingReturn type
func DecodeNsBulkVolSettingReturn(request interface{}) (*NsBulkVolSettingReturn, error) {
	reqNsBulkVolSettingReturn := request.(*nsBulkVolSettingReturn)
	byte, err := json.Marshal(reqNsBulkVolSettingReturn)
	obj := &NsBulkVolSettingReturn{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
