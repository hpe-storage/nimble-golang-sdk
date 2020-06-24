/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsNIC - Network interface controller.
// Export NsNICFields for advance operations like search filter etc.
var NsNICFields *NsNIC

func init(){
	Namefield:= "name"
	SubnetLabelfield:= "subnet_label"
	DataIpfield:= "data_ip"
		
	NsNICFields= &NsNIC{
		Name: &Namefield,
		SubnetLabel: &SubnetLabelfield,
		DataIp: &DataIpfield,
		
	}
}

type NsNIC struct {
   
    // Name of NIC.
    
 	Name *string `json:"name,omitempty"`
   
    // Subnet label for this NIC.
    
 	SubnetLabel *string `json:"subnet_label,omitempty"`
   
    // Data IP address.
    
 	DataIp *string `json:"data_ip,omitempty"`
   
    // Identify whether the NIC is tagged.
    
 	Tagged *bool `json:"tagged,omitempty"`
}
