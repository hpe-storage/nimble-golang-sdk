// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsShelfCtrlrAttrSet - A shelf logical controller attributes.
// Export NsShelfCtrlrAttrSetFields for advance operations like search filter etc.
var NsShelfCtrlrAttrSetFields *NsShelfCtrlrAttrSet

func init(){
 SessionSerialfield:= "session_serial"
 CachedSerialfield:= "cached_serial"
 DiskSerialsfield:= "disk_serials"
 DiskTypesfield:= "disk_types"

 NsShelfCtrlrAttrSetFields= &NsShelfCtrlrAttrSet{
  SessionSerial: &SessionSerialfield,
  CachedSerial:  &CachedSerialfield,
  DiskSerials:   &DiskSerialsfield,
  DiskTypes:     &DiskTypesfield,
 }
}


type NsShelfCtrlrAttrSet struct {
 // SessionSerial - Session serial.
  SessionSerial *string `json:"session_serial,omitempty"`
 // CachedSerial - Cached serial.
  CachedSerial *string `json:"cached_serial,omitempty"`
 // HwState - The hardware state for this logical controller.
    HwState *NsShelfHwState `json:"hw_state,omitempty"`
 // SwType - The software type of this logical controller.
    SwType *NsShelfSwType `json:"sw_type,omitempty"`
 // DiskSerials - Comma separated list of disk serials connected to this logical controller.
  DiskSerials *string `json:"disk_serials,omitempty"`
 // DiskTypes - Comma separated list of disk types (H for HDD, S for SSD).
  DiskTypes *string `json:"disk_types,omitempty"`
}
