// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsArrayFcConfig - Array Fibre Channel configuration.
// Export NsArrayFcConfigFields for advance operations like search filter etc.
var NsArrayFcConfigFields *NsArrayFcConfig

func init(){
	Namefield:= "name"
	ArrayNamefield:= "array_name"
	IDfield:= "id"
	ArrayIdfield:= "array_id"
		
	NsArrayFcConfigFields= &NsArrayFcConfig{
	Name: &Namefield,
	ArrayName: &ArrayNamefield,
	ID: &IDfield,
	ArrayId: &ArrayIdfield,
		
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
