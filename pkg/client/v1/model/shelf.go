/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Shelf - Disk shelf and head unit houses disks and controller.
// Export ShelfFields for advance operations like search filter etc.
var ShelfFields *Shelf

func init(){
	IDfield:= "id"
	ArrayNamefield:= "array_name"
	ArrayIDfield:= "array_id"
	Serialfield:= "serial"
	Modelfield:= "model"
	ModelExtfield:= "model_ext"
		
	ShelfFields= &Shelf{
		ID: &IDfield,
		ArrayName: &ArrayNamefield,
		ArrayID: &ArrayIDfield,
		Serial: &Serialfield,
		Model: &Modelfield,
		ModelExt: &ModelExtfield,
		
	}
}

type Shelf struct {
   
    // ID of shelf.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of array the shelf belongs to.
    
 	ArrayName *string `json:"array_name,omitempty"`
   
    // ID of array the shelf belongs to.
    
 	ArrayID *string `json:"array_id,omitempty"`
   
    // Indicate that it is okay to provide partially available response.
    
 	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
   
    // Chassis type.
    
   	ChassisType *NsChassisType `json:"chassis_type,omitempty"`
   
    // List of ctrlr info.
    
   	Ctrlrs []*NsShelfCtrlr `json:"ctrlrs,omitempty"`
   
    // The serial number of the chassis.
    
 	Serial *string `json:"serial,omitempty"`
   
    // Model of the shelf or head unit.
    
 	Model *string `json:"model,omitempty"`
   
    // Extended model of the shelf or head unit.
    
 	ModelExt *string `json:"model_ext,omitempty"`
   
    // List of chassis sensor readings.
    
   	ChassisSensors []*NsShelfSensor `json:"chassis_sensors,omitempty"`
   
    // The overall status for the PSUs.
    
   	PsuOverallStatus *NsShelfSensorState `json:"psu_overall_status,omitempty"`
   
    // The overall status for the fans on both controllers.
    
   	FanOverallStatus *NsShelfSensorState `json:"fan_overall_status,omitempty"`
   
    // The overall status for the temperature on both controllers.
    
   	TempOverallStatus *NsShelfSensorState `json:"temp_overall_status,omitempty"`
   
    // Attributes for the disk sets in this shelf.
    
   	DiskSets []*NsDiskSetAttr `json:"disk_sets,omitempty"`
   
    // Activated state for shelf or disk set means it is available to store date on. An activated shelf may not be deactivated.
    
 	Activated *bool `json:"activated,omitempty"`
   
    // Driveset to activate.
    
   	Driveset *int64 `json:"driveset,omitempty"`
   
    // Forcibly activate shelf.
    
 	Force *bool `json:"force,omitempty"`
   
    // Accept the removal of data on the shelf disks and activate foreign shelf.
    
 	AcceptForeign *bool `json:"accept_foreign,omitempty"`
   
    // Accept the reduction or elimination of deduplication capability on the system as a result of activating a shelf that does not meet the necessary deduplication requirements.
    
 	AcceptDedupeImpact *bool `json:"accept_dedupe_impact,omitempty"`
   
    // Indicates this is the last request in a series of shelf add requests.
    
 	LastRequest *bool `json:"last_request,omitempty"`
}
