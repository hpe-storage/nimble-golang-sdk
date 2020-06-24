/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsShelfCtrlr - A shelf controller information.
// Export NsShelfCtrlrFields for advance operations like search filter etc.
var NsShelfCtrlrFields *NsShelfCtrlr

func init(){
	ExpSasAddrfield:= "exp_sas_addr"
	CachedSerialfield:= "cached_serial"
		
	NsShelfCtrlrFields= &NsShelfCtrlr{
		ExpSasAddr: &ExpSasAddrfield,
		CachedSerial: &CachedSerialfield,
		
	}
}

type NsShelfCtrlr struct {
   
    // Expander SAS address.
    
 	ExpSasAddr *string `json:"exp_sas_addr,omitempty"`
   
    // Position of this controller in the chassis.
    
   	CtrlrSIDe *NsShelfCtrlrSIDe `json:"ctrlr_side,omitempty"`
   
    // Location ID based on SAS topology.
    
  	EncLocID  *int64 `json:"enc_loc_id,omitempty"`
   
    // Cached serial.
    
 	CachedSerial *string `json:"cached_serial,omitempty"`
   
    // The time of last valid sensor reading, in epoch seconds.
    
   	CtrlrSensorLastRun *int64 `json:"ctrlr_sensor_last_run,omitempty"`
   
    // SES device hardware mastership failure.
    
 	HwMshipFailure *bool `json:"hw_mship_failure,omitempty"`
   
    // SES device hardware mastership state.
    
   	HwMasterState *NsShelfSesMasterHwState `json:"hw_master_state,omitempty"`
   
    // SES device software mastership state.
    
   	SwMasterState *NsShelfSesMasterSwState `json:"sw_master_state,omitempty"`
   
    // The list of individual sensor reading in this controller.
    
   	CtrlrSensors []*NsShelfSensor `json:"ctrlr_sensors,omitempty"`
   
    // The overall status for the fans on this controller.
    
   	FanOverallStatus *NsShelfSensorState `json:"fan_overall_status,omitempty"`
   
    // The overall status for the temperature of this controller.
    
   	TempOverallStatus *NsShelfSensorState `json:"temp_overall_status,omitempty"`
   
    // Controller hardware model.
    
   	CtrlrHwModel *NsShelfCtrlrModel `json:"ctrlr_hw_model,omitempty"`
   
    // Port info for each SAS port.
    
   	PortInfo []*NsShelfPortInfo `json:"port_info,omitempty"`
   
    // Extra attributes as attribute value pairs.
    
	ExtraAttributes []*string `json:"extra_attributes,omitempty"`
   
    // List of ctrlr attribute set for each logical controller.
    
   	CtrlrAttrsetList []*NsShelfCtrlrAttrSet `json:"ctrlr_attrset_list,omitempty"`
   
    // Status of chassis identifier.
    
 	IDentifyStatus *bool `json:"identify_status,omitempty"`
}
