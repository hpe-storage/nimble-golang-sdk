// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapLunInfoFields provides field names to use in filter parameters, for example.
var NsSnapLunInfoFields *NsSnapLunInfoFieldHandles

func init() {
	NsSnapLunInfoFields = &NsSnapLunInfoFieldHandles{
		ID:   "id",
		Name: "name",
		Lun:  "lun",
	}
}

// NsSnapLunInfo - Snapshot LUN information as presented in the access_control_records object set. This information is only available for Fibre Channel.
type NsSnapLunInfo struct {
	// ID - Snapshot ID.
	ID *string `json:"id,omitempty"`
	// Name - Snapshot name.
	Name *string `json:"name,omitempty"`
	// Lun - Snapshot LUN.
	Lun *int64 `json:"lun,omitempty"`
}

// NsSnapLunInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsSnapLunInfoFieldHandles struct {
	ID   string
	Name string
	Lun  string
}
