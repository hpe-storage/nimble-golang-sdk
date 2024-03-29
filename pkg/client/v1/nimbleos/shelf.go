// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// ShelfFields provides field names to use in filter parameters, for example.
var ShelfFields *ShelfFieldHandles

func init() {
	ShelfFields = &ShelfFieldHandles{
		ID:                 "id",
		ArrayName:          "array_name",
		ArrayId:            "array_id",
		PartialResponseOk:  "partial_response_ok",
		ChassisType:        "chassis_type",
		Ctrlrs:             "ctrlrs",
		Serial:             "serial",
		Model:              "model",
		ModelExt:           "model_ext",
		ChassisSensors:     "chassis_sensors",
		PsuOverallStatus:   "psu_overall_status",
		FanOverallStatus:   "fan_overall_status",
		TempOverallStatus:  "temp_overall_status",
		DiskSets:           "disk_sets",
		Activated:          "activated",
		Driveset:           "driveset",
		Force:              "force",
		AcceptForeign:      "accept_foreign",
		AcceptDedupeImpact: "accept_dedupe_impact",
		LastRequest:        "last_request",
	}
}

// Shelf - Disk shelf and head unit houses disks and controller.
type Shelf struct {
	// ID - ID of shelf.
	ID *string `json:"id,omitempty"`
	// ArrayName - Name of array the shelf belongs to.
	ArrayName *string `json:"array_name,omitempty"`
	// ArrayId - ID of array the shelf belongs to.
	ArrayId *string `json:"array_id,omitempty"`
	// PartialResponseOk - Indicate that it is okay to provide partially available response.
	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
	// ChassisType - Chassis type.
	ChassisType *NsChassisType `json:"chassis_type,omitempty"`
	// Ctrlrs - List of ctrlr info.
	Ctrlrs []*NsShelfCtrlr `json:"ctrlrs,omitempty"`
	// Serial - The serial number of the chassis.
	Serial *string `json:"serial,omitempty"`
	// Model - Model of the shelf or head unit.
	Model *string `json:"model,omitempty"`
	// ModelExt - Extended model of the shelf or head unit.
	ModelExt *string `json:"model_ext,omitempty"`
	// ChassisSensors - List of chassis sensor readings.
	ChassisSensors []*NsShelfSensor `json:"chassis_sensors,omitempty"`
	// PsuOverallStatus - The overall status for the PSUs.
	PsuOverallStatus *NsShelfSensorState `json:"psu_overall_status,omitempty"`
	// FanOverallStatus - The overall status for the fans on both controllers.
	FanOverallStatus *NsShelfSensorState `json:"fan_overall_status,omitempty"`
	// TempOverallStatus - The overall status for the temperature on both controllers.
	TempOverallStatus *NsShelfSensorState `json:"temp_overall_status,omitempty"`
	// DiskSets - Attributes for the disk sets in this shelf.
	DiskSets []*NsDiskSetAttr `json:"disk_sets,omitempty"`
	// Activated - Activated state for shelf or disk set means it is available to store date on. An activated shelf may not be deactivated.
	Activated *bool `json:"activated,omitempty"`
	// Driveset - Driveset to activate.
	Driveset *int64 `json:"driveset,omitempty"`
	// Force - Forcibly activate shelf.
	Force *bool `json:"force,omitempty"`
	// AcceptForeign - Accept the removal of data on the shelf disks and activate foreign shelf.
	AcceptForeign *bool `json:"accept_foreign,omitempty"`
	// AcceptDedupeImpact - Accept the reduction or elimination of deduplication capability on the system as a result of activating a shelf that does not meet the necessary deduplication requirements.
	AcceptDedupeImpact *bool `json:"accept_dedupe_impact,omitempty"`
	// LastRequest - Indicates this is the last request in a series of shelf add requests.
	LastRequest *bool `json:"last_request,omitempty"`
}

// ShelfFieldHandles provides a string representation for each Shelf field.
type ShelfFieldHandles struct {
	ID                 string
	ArrayName          string
	ArrayId            string
	PartialResponseOk  string
	ChassisType        string
	Ctrlrs             string
	Serial             string
	Model              string
	ModelExt           string
	ChassisSensors     string
	PsuOverallStatus   string
	FanOverallStatus   string
	TempOverallStatus  string
	DiskSets           string
	Activated          string
	Driveset           string
	Force              string
	AcceptForeign      string
	AcceptDedupeImpact string
	LastRequest        string
}
