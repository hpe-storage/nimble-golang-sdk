// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsShelfCtrlrAttrSetFields provides field names to use in filter parameters, for example.
var NsShelfCtrlrAttrSetFields *NsShelfCtrlrAttrSetFieldHandles

func init() {
	fieldSessionSerial := "session_serial"
	fieldCachedSerial := "cached_serial"
	fieldHwState := "hw_state"
	fieldSwType := "sw_type"
	fieldDiskSerials := "disk_serials"
	fieldDiskTypes := "disk_types"

	NsShelfCtrlrAttrSetFields = &NsShelfCtrlrAttrSetFieldHandles{
		SessionSerial: &fieldSessionSerial,
		CachedSerial:  &fieldCachedSerial,
		HwState:       &fieldHwState,
		SwType:        &fieldSwType,
		DiskSerials:   &fieldDiskSerials,
		DiskTypes:     &fieldDiskTypes,
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
	SessionSerial *string
	CachedSerial  *string
	HwState       *string
	SwType        *string
	DiskSerials   *string
	DiskTypes     *string
}
