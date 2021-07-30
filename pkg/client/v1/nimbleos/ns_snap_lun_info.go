// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapLunInfo - Snapshot LUN information as presented in the access_control_records object set. This information is only available for Fibre Channel.

// Export NsSnapLunInfoFields provides field names to use in filter parameters, for example.
var NsSnapLunInfoFields *NsSnapLunInfoStringFields

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldLun := "lun"

	NsSnapLunInfoFields = &NsSnapLunInfoStringFields{
		ID:   &fieldID,
		Name: &fieldName,
		Lun:  &fieldLun,
	}
}

type NsSnapLunInfo struct {
	// ID - Snapshot ID.
	ID *string `json:"id,omitempty"`
	// Name - Snapshot name.
	Name *string `json:"name,omitempty"`
	// Lun - Snapshot LUN.
	Lun *int64 `json:"lun,omitempty"`
}

// Struct for NsSnapLunInfoFields
type NsSnapLunInfoStringFields struct {
	ID   *string
	Name *string
	Lun  *string
}
