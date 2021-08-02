// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCtrlrNvmeCardFields provides field names to use in filter parameters, for example.
var NsCtrlrNvmeCardFields *NsCtrlrNvmeCardFieldHandles

func init() {
	NsCtrlrNvmeCardFields = &NsCtrlrNvmeCardFieldHandles{
		SerialNumber: "serial_number",
		Size:         "size",
		State:        "state",
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
	SerialNumber string
	Size         string
	State        string
}
