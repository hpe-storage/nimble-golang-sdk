// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsArrayFcConfigFields provides field names to use in filter parameters, for example.
var NsArrayFcConfigFields *NsArrayFcConfigFieldHandles

func init() {
	NsArrayFcConfigFields = &NsArrayFcConfigFieldHandles{
		Name:           "name",
		ArrayName:      "array_name",
		ID:             "id",
		ArrayId:        "array_id",
		CtrlrAFcConfig: "ctrlr_a_fc_config",
		CtrlrBFcConfig: "ctrlr_b_fc_config",
	}
}

// NsArrayFcConfig - Array Fibre Channel configuration.
type NsArrayFcConfig struct {
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

// NsArrayFcConfigFieldHandles provides a string representation for each AccessControlRecord field.
type NsArrayFcConfigFieldHandles struct {
	Name           string
	ArrayName      string
	ID             string
	ArrayId        string
	CtrlrAFcConfig string
	CtrlrBFcConfig string
}
