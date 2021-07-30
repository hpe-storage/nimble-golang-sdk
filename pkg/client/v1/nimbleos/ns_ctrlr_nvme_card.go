// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsCtrlrNvmeCardFields provides field names to use in filter parameters, for example.
var NsCtrlrNvmeCardFields *NsCtrlrNvmeCardFieldHandles

func init() {
	fieldSerialNumber := "serial_number"
	fieldSize := "size"
	fieldState := "state"

	NsCtrlrNvmeCardFields = &NsCtrlrNvmeCardFieldHandles{
		SerialNumber: &fieldSerialNumber,
		Size:         &fieldSize,
		State:        &fieldState,
	}
}

// NsCtrlrNvmeCard - NVMe accelerator card.
type NsCtrlrNvmeCard struct {
	// SerialNumber - Serial number.
	SerialNumber *string `json:"serial_number,omitempty"`
	// Size - NVMe card cache size in bytes.
	Size *int64 `json:"size,omitempty"`
	// State - Online state.
	State *NsCtrlrNvmeCardState `json:"state,omitempty"`
}

// NsCtrlrNvmeCardFieldHandles provides a string representation for each AccessControlRecord field.
type NsCtrlrNvmeCardFieldHandles struct {
	SerialNumber *string
	Size         *string
	State        *string
}
