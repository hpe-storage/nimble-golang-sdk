/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Controller - Controller is a redundant collection of hardware capable of running the array software.
// Export ControllerFields for advance operations like search filter etc.
var ControllerFields *Controller

func init(){
	IDfield:= "id"
	Namefield:= "name"
	ArrayNamefield:= "array_name"
	ArrayIDfield:= "array_id"
	Serialfield:= "serial"
	Hostnamefield:= "hostname"
	SupportAddressfield:= "support_address"
	SupportNetmaskfield:= "support_netmask"
	SupportNicfield:= "support_nic"
		
	ControllerFields= &Controller{
		ID: &IDfield,
		Name: &Namefield,
		ArrayName: &ArrayNamefield,
		ArrayID: &ArrayIDfield,
		Serial: &Serialfield,
		Hostname: &Hostnamefield,
		SupportAddress: &SupportAddressfield,
		SupportNetmask: &SupportNetmaskfield,
		SupportNic: &SupportNicfield,
		
	}
}

type Controller struct {
   
    // Identifier of the controller.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of the controller.
    
 	Name *string `json:"name,omitempty"`
   
    // Name of the array containing this controller.
    
 	ArrayName *string `json:"array_name,omitempty"`
   
    // Rest ID of the array containing this controller.
    
 	ArrayID *string `json:"array_id,omitempty"`
   
    // Indicate that it is ok to provide partially available response.
    
 	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
   
    // Serial number for this controller.
    
 	Serial *string `json:"serial,omitempty"`
   
    // Host name for the controller.
    
 	Hostname *string `json:"hostname,omitempty"`
   
    // IP address used for support.
    
 	SupportAddress *string `json:"support_address,omitempty"`
   
    // IP netmask used for support.
    
 	SupportNetmask *string `json:"support_netmask,omitempty"`
   
    // Network card used for support.
    
 	SupportNic *string `json:"support_nic,omitempty"`
   
    // Overall power supply status for the controller.
    
   	PowerStatus *NsPowerSupplyStatus `json:"power_status,omitempty"`
   
    // Overall fan status for the controller.
    
   	FanStatus *NsFanStatus `json:"fan_status,omitempty"`
   
    // Overall temperature status for the controller.
    
   	TemperatureStatus *NsTemperatureStatus `json:"temperature_status,omitempty"`
   
    // Status for each power supply in the controller.
    
   	PowerSupplies []*NsCtrlrHwSensorInfo `json:"power_supplies,omitempty"`
   
    // Status for each fan in the controller.
    
   	Fans []*NsCtrlrHwSensorInfo `json:"fans,omitempty"`
   
    // Status for temperature sensor in the controller.
    
   	TemperatureSensors []*NsCtrlrHwSensorInfo `json:"temperature_sensors,omitempty"`
   
    // Status of the system's raid partitions.
    
   	PartitionStatus []*NsCtrlrRaIDInfo `json:"partition_status,omitempty"`
   
    // Identifies which controller this is on its array.
    
   	CtrlrSIDe *NsControllerSIDe `json:"ctrlr_side,omitempty"`
   
    // Indicates whether this controller is active or not.
    
   	State *NsControllerState `json:"state,omitempty"`
   
    // Indicates if the NVMe accelerator card is enabled.
    
   	NvmeCardsEnabled *int64 `json:"nvme_cards_enabled,omitempty"`
   
    // List of NVMe accelerator cards.
    
   	NvmeCards []*NsCtrlrNvmeCard `json:"nvme_cards,omitempty"`
   
    // Time of the last autosupport by the controller.
    
   	AsupTime *int64 `json:"asup_time,omitempty"`
}
