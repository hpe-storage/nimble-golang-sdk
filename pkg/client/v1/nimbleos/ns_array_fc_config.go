// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArrayFcConfig - Array Fibre Channel configuration.
// Export NsArrayFcConfigFields for advance operations like search filter etc.
var NsArrayFcConfigFields *NsArrayFcConfigStringFields

func init() {
	Namefield := "name"
	ArrayNamefield := "array_name"
	IDfield := "id"
	ArrayIdfield := "array_id"
	CtrlrAFcConfigfield := "ctrlr_a_fc_config"
	CtrlrBFcConfigfield := "ctrlr_b_fc_config"

	NsArrayFcConfigFields = &NsArrayFcConfigStringFields{
		Name:           &Namefield,
		ArrayName:      &ArrayNamefield,
		ID:             &IDfield,
		ArrayId:        &ArrayIdfield,
		CtrlrAFcConfig: &CtrlrAFcConfigfield,
		CtrlrBFcConfig: &CtrlrBFcConfigfield,
	}
}

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

// Struct for NsArrayFcConfigFields
type NsArrayFcConfigStringFields struct {
	Name           *string
	ArrayName      *string
	ID             *string
	ArrayId        *string
	CtrlrAFcConfig *string
	CtrlrBFcConfig *string
}
