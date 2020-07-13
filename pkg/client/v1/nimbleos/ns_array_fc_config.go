// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsArrayFcConfig - Array Fibre Channel configuration.
// Export NsArrayFcConfigFields for advance operations like search filter etc.
var NsArrayFcConfigFields *NsArrayFcConfig

func init() {

	NsArrayFcConfigFields = &NsArrayFcConfig{
		Name:      "name",
		ArrayName: "array_name",
		ID:        "id",
		ArrayId:   "array_id",
	}
}

type NsArrayFcConfig struct {
	// Name - Name of the array.
	Name string `json:"name,omitempty"`
	// ArrayName - Name of the array.
	ArrayName string `json:"array_name,omitempty"`
	// ID - ID of the array.
	ID string `json:"id,omitempty"`
	// ArrayId - ID of the array.
	ArrayId string `json:"array_id,omitempty"`
	// CtrlrAFcConfig - Controller A Fibre Channel configuration.
	CtrlrAFcConfig *NsCtrlrFcConfig `json:"ctrlr_a_fc_config,omitempty"`
	// CtrlrBFcConfig - Controller B Fibre Channel configuration.
	CtrlrBFcConfig *NsCtrlrFcConfig `json:"ctrlr_b_fc_config,omitempty"`
}

// sdk internal struct
type nsArrayFcConfig struct {
	// Name - Name of the array.
	Name *string `json:"name,omitempty"`
	// ArrayName - Name of the array.
	ArrayName *string `json:"array_name,omitempty"`
	// ID - ID of the array.
	ID *string `json:"id,omitempty"`
	// ArrayId - ID of the array.
	ArrayId *string `json:"array_id,omitempty"`
	// CtrlrAFcConfig - Controller A Fibre Channel configuration.
	CtrlrAFcConfig *NsCtrlrFcConfig `json:"ctrlr_a_fc_config,omitempty"`
	// CtrlrBFcConfig - Controller B Fibre Channel configuration.
	CtrlrBFcConfig *NsCtrlrFcConfig `json:"ctrlr_b_fc_config,omitempty"`
}

// EncodeNsArrayFcConfig - Transform NsArrayFcConfig to nsArrayFcConfig type
func EncodeNsArrayFcConfig(request interface{}) (*nsArrayFcConfig, error) {
	reqNsArrayFcConfig := request.(*NsArrayFcConfig)
	byte, err := json.Marshal(reqNsArrayFcConfig)
	objPtr := &nsArrayFcConfig{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsArrayFcConfig Transform nsArrayFcConfig to NsArrayFcConfig type
func DecodeNsArrayFcConfig(request interface{}) (*NsArrayFcConfig, error) {
	reqNsArrayFcConfig := request.(*nsArrayFcConfig)
	byte, err := json.Marshal(reqNsArrayFcConfig)
	obj := &NsArrayFcConfig{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
