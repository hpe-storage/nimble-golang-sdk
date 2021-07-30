// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcInterfaceUpdateInfo - Fibre Channel interface information to update.

// Export NsFcInterfaceUpdateInfoFields provides field names to use in filter parameters, for example.
var NsFcInterfaceUpdateInfoFields *NsFcInterfaceUpdateInfoStringFields

func init() {
	fieldID := "id"
	fieldOnline := "online"

	NsFcInterfaceUpdateInfoFields = &NsFcInterfaceUpdateInfoStringFields{
		ID:     &fieldID,
		Online: &fieldOnline,
	}
}

type NsFcInterfaceUpdateInfo struct {
	// ID - ID of Fibre Channel interface.
	ID *string `json:"id,omitempty"`
	// Online - Identify whether the Fibre Channel interface is online.
	Online *bool `json:"online,omitempty"`
}

// Struct for NsFcInterfaceUpdateInfoFields
type NsFcInterfaceUpdateInfoStringFields struct {
	ID     *string
	Online *string
}
