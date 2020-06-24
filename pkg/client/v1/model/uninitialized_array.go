/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// UninitializedArray - Lists discovered arrays that are not members of any group and are in the same subnet.
// Export UninitializedArrayFields for advance operations like search filter etc.
var UninitializedArrayFields *UninitializedArray

func init(){
	IDfield:= "id"
	Serialfield:= "serial"
	ArrayNamefield:= "array_name"
	Modelfield:= "model"
	ModelStrfield:= "model_str"
	Versionfield:= "version"
	IpAddressfield:= "ip_address"
		
	UninitializedArrayFields= &UninitializedArray{
		ID: &IDfield,
		Serial: &Serialfield,
		ArrayName: &ArrayNamefield,
		Model: &Modelfield,
		ModelStr: &ModelStrfield,
		Version: &Versionfield,
		IpAddress: &IpAddressfield,
		
	}
}

type UninitializedArray struct {
   
    // Identifier for the interface.
    
 	ID *string `json:"id,omitempty"`
   
    // Serial Number of the uninitialized array.
    
 	Serial *string `json:"serial,omitempty"`
   
    // Name of the uninitialized array.
    
 	ArrayName *string `json:"array_name,omitempty"`
   
    // Model of the uninitialized array.
    
 	Model *string `json:"model,omitempty"`
   
    // Model description of the uninitialized array.
    
 	ModelStr *string `json:"model_str,omitempty"`
   
    // Version of the uninitialized array.
    
 	Version *string `json:"version,omitempty"`
   
    // Link local zero conf address of the uninitialized array.
    
 	IpAddress *string `json:"ip_address,omitempty"`
   
    // List of link local zero conf address of the uninitialized array.
    
   	ZconfIpaddrs []*NsIPAddressObject `json:"zconf_ipaddrs,omitempty"`
   
    // Number of Fibre Channel ports of the uninitialized array.
    
   	CountOfFcPorts *int64 `json:"count_of_fc_ports,omitempty"`
   
    // True if it is an All-Flash array, False otherwise.
    
 	AllFlash *bool `json:"all_flash,omitempty"`
}
