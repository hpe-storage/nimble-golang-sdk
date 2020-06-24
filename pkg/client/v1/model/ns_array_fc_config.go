/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsArrayFcConfig - Array Fibre Channel configuration.
// Export NsArrayFcConfigFields for advance operations like search filter etc.
var NsArrayFcConfigFields *NsArrayFcConfig

func init(){
	Namefield:= "name"
	ArrayNamefield:= "array_name"
	IDfield:= "id"
	ArrayIDfield:= "array_id"
		
	NsArrayFcConfigFields= &NsArrayFcConfig{
		Name: &Namefield,
		ArrayName: &ArrayNamefield,
		ID: &IDfield,
		ArrayID: &ArrayIDfield,
		
	}
}

type NsArrayFcConfig struct {
   
    // Name of the array.
    
 	Name *string `json:"name,omitempty"`
   
    // Name of the array.
    
 	ArrayName *string `json:"array_name,omitempty"`
   
    // ID of the array.
    
 	ID *string `json:"id,omitempty"`
   
    // ID of the array.
    
 	ArrayID *string `json:"array_id,omitempty"`
   
    // Controller A Fibre Channel configuration.
    
	CtrlrAFcConfig *NsCtrlrFcConfig `json:"ctrlr_a_fc_config,omitempty"`
   
    // Controller B Fibre Channel configuration.
    
	CtrlrBFcConfig *NsCtrlrFcConfig `json:"ctrlr_b_fc_config,omitempty"`
}
