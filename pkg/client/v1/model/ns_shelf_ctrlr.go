/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsShelfCtrlr


// NsShelfCtrlr :
type NsShelfCtrlr struct {
   // ExpSasAddr
   ExpSasAddr string `json:"exp_sas_addr,omitempty"`
   // EncLocID
   EncLocID  int32 `json:"enc_loc_id,omitempty"`
   // CachedSerial
   CachedSerial string `json:"cached_serial,omitempty"`
   // CtrlrSensorLastRun
   CtrlrSensorLastRun float64 `json:"ctrlr_sensor_last_run,omitempty"`
   // HwMshipFailure
   HwMshipFailure bool `json:"hw_mship_failure,omitempty"`
   // IDentifyStatus
   IDentifyStatus bool `json:"identify_status,omitempty"`
}
