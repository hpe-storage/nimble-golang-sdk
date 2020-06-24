/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsFcPortInfo - Fibre Channel port information.
// Export NsFcPortInfoFields for advance operations like search filter etc.
var NsFcPortInfoFields *NsFcPortInfo

func init(){
	Namefield:= "name"
	BusLocationfield:= "bus_location"
		
	NsFcPortInfoFields= &NsFcPortInfo{
		Name: &Namefield,
		BusLocation: &BusLocationfield,
		
	}
}

type NsFcPortInfo struct {
   
    // Name of Fibre Channel port.
    
 	Name *string `json:"name,omitempty"`
   
    // PCI bus location of the HBA for this Fibre Channel port.
    
 	BusLocation *string `json:"bus_location,omitempty"`
   
    // HBA port number for this Fibre Channel port.
    
   	Port *int64 `json:"port,omitempty"`
   
    // HBA slot number for this Fibre Channel port.
    
   	Slot *int64 `json:"slot,omitempty"`
}
