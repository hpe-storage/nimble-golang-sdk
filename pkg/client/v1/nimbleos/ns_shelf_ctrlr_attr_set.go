// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsShelfCtrlrAttrSetFields provides field names to use in filter parameters, for example.
var NsShelfCtrlrAttrSetFields *NsShelfCtrlrAttrSetFieldHandles

func init() {
	NsShelfCtrlrAttrSetFields = &NsShelfCtrlrAttrSetFieldHandles{
		SessionSerial: "session_serial",
		CachedSerial:  "cached_serial",
		HwState:       "hw_state",
		SwType:        "sw_type",
		DiskSerials:   "disk_serials",
		DiskTypes:     "disk_types",
	}
}

// NsShelfCtrlrAttrSet - A shelf logical controller attributes.
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

// NsShelfCtrlrAttrSetFieldHandles provides a string representation for each AccessControlRecord field.
type NsShelfCtrlrAttrSetFieldHandles struct {
	SessionSerial string
	CachedSerial  string
	HwState       string
	SwType        string
	DiskSerials   string
	DiskTypes     string
}
