// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsShelfCtrlrFields provides field names to use in filter parameters, for example.
var NsShelfCtrlrFields *NsShelfCtrlrFieldHandles

func init() {
	NsShelfCtrlrFields = &NsShelfCtrlrFieldHandles{
		ExpSasAddr:         "exp_sas_addr",
		CtrlrSide:          "ctrlr_side",
		EncLocId:           "enc_loc_id",
		CachedSerial:       "cached_serial",
		CtrlrSensorLastRun: "ctrlr_sensor_last_run",
		HwMshipFailure:     "hw_mship_failure",
		HwMasterState:      "hw_master_state",
		SwMasterState:      "sw_master_state",
		CtrlrSensors:       "ctrlr_sensors",
		FanOverallStatus:   "fan_overall_status",
		TempOverallStatus:  "temp_overall_status",
		PsuOverallStatus:   "psu_overall_status",
		CtrlrHwModel:       "ctrlr_hw_model",
		PortInfo:           "port_info",
		ExtraAttributes:    "extra_attributes",
		CtrlrAttrsetList:   "ctrlr_attrset_list",
		IDentifyStatus:     "identify_status",
	}
}

// NsShelfCtrlr - A shelf controller information.
type NsShelfCtrlr struct {
	// ExpSasAddr - Expander SAS address.
	ExpSasAddr *string `json:"exp_sas_addr,omitempty"`
	// CtrlrSide - Position of this controller in the chassis.
	CtrlrSide *NsShelfCtrlrSide `json:"ctrlr_side,omitempty"`
	// EncLocId - Location ID based on SAS topology.
	EncLocId *int64 `json:"enc_loc_id,omitempty"`
	// CachedSerial - Cached serial.
	CachedSerial *string `json:"cached_serial,omitempty"`
	// CtrlrSensorLastRun - The time of last valid sensor reading, in epoch seconds.
	CtrlrSensorLastRun *int64 `json:"ctrlr_sensor_last_run,omitempty"`
	// HwMshipFailure - SES device hardware mastership failure.
	HwMshipFailure *bool `json:"hw_mship_failure,omitempty"`
	// HwMasterState - SES device hardware mastership state.
	HwMasterState *NsShelfSesMasterHwState `json:"hw_master_state,omitempty"`
	// SwMasterState - SES device software mastership state.
	SwMasterState *NsShelfSesMasterSwState `json:"sw_master_state,omitempty"`
	// CtrlrSensors - The list of individual sensor reading in this controller.
	CtrlrSensors []*NsShelfSensor `json:"ctrlr_sensors,omitempty"`
	// FanOverallStatus - The overall status for the fans on this controller.
	FanOverallStatus *NsShelfSensorState `json:"fan_overall_status,omitempty"`
	// TempOverallStatus - The overall status for the temperature of this controller.
	TempOverallStatus *NsShelfSensorState `json:"temp_overall_status,omitempty"`
	// PsuOverallStatus - The overall status for the PSU on this controller.
	PsuOverallStatus *NsShelfSensorState `json:"psu_overall_status,omitempty"`
	// CtrlrHwModel - Controller hardware model.
	CtrlrHwModel *NsShelfCtrlrModel `json:"ctrlr_hw_model,omitempty"`
	// PortInfo - Port info for each SAS port.
	PortInfo []*NsShelfPortInfo `json:"port_info,omitempty"`
	// ExtraAttributes - Extra attributes as attribute value pairs.
	ExtraAttributes []*string `json:"extra_attributes,omitempty"`
	// CtrlrAttrsetList - List of ctrlr attribute set for each logical controller.
	CtrlrAttrsetList []*NsShelfCtrlrAttrSet `json:"ctrlr_attrset_list,omitempty"`
	// IDentifyStatus - Status of chassis identifier.
	IDentifyStatus *bool `json:"identify_status,omitempty"`
}

// NsShelfCtrlrFieldHandles provides a string representation for each NsShelfCtrlr field.
type NsShelfCtrlrFieldHandles struct {
	ExpSasAddr         string
	CtrlrSide          string
	EncLocId           string
	CachedSerial       string
	CtrlrSensorLastRun string
	HwMshipFailure     string
	HwMasterState      string
	SwMasterState      string
	CtrlrSensors       string
	FanOverallStatus   string
	TempOverallStatus  string
	PsuOverallStatus   string
	CtrlrHwModel       string
	PortInfo           string
	ExtraAttributes    string
	CtrlrAttrsetList   string
	IDentifyStatus     string
}
