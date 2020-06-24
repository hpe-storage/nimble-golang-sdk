/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsArraySummaryInfo - Array summary information, including version, model, and IP configurations.
// Export NsArraySummaryInfoFields for advance operations like search filter etc.
var NsArraySummaryInfoFields *NsArraySummaryInfo

func init(){
	Namefield:= "name"
	Versionfield:= "version"
	Serialfield:= "serial"
	Modelfield:= "model"
		
	NsArraySummaryInfoFields= &NsArraySummaryInfo{
		Name: &Namefield,
		Version: &Versionfield,
		Serial: &Serialfield,
		Model: &Modelfield,
		
	}
}

type NsArraySummaryInfo struct {
   
    // Unique name of array.
    
 	Name *string `json:"name,omitempty"`
   
    // Version of array.
    
 	Version *string `json:"version,omitempty"`
   
    // Array serial number.
    
 	Serial *string `json:"serial,omitempty"`
   
    // Array hardware model.
    
 	Model *string `json:"model,omitempty"`
   
    // Zero Conf IP of array, including Nic, local and remote IP addresses.
    
   	ZconfIpaddrs []*NsZeroConfIPAddr `json:"zconf_ipaddrs,omitempty"`
   
    // Status of array.
    
   	Status *NsArrayStatus `json:"status,omitempty"`
   
    // Count of Fibre Channel ports per controller.
    
   	CountOfFcPorts *int64 `json:"count_of_fc_ports,omitempty"`
   
    // Whether it is an all-flash array.
    
 	AllFlash *bool `json:"all_flash,omitempty"`
}
