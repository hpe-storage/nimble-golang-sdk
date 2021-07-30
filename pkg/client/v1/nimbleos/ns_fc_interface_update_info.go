// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsFcInterfaceUpdateInfoFields provides field names to use in filter parameters, for example.
var NsFcInterfaceUpdateInfoFields *NsFcInterfaceUpdateInfoFieldHandles

func init() {
	NsFcInterfaceUpdateInfoFields = &NsFcInterfaceUpdateInfoFieldHandles{
		ID:     "id",
		Online: "online",
	}
}

// NsFcInterfaceUpdateInfo - Fibre Channel interface information to update.
type NsFcInterfaceUpdateInfo struct {
	// ID - ID of Fibre Channel interface.
	ID *string `json:"id,omitempty"`
	// Online - Identify whether the Fibre Channel interface is online.
	Online *bool `json:"online,omitempty"`
}

// NsFcInterfaceUpdateInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsFcInterfaceUpdateInfoFieldHandles struct {
	ID     string
	Online string
}
