// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsTargetSubnetFields provides field names to use in filter parameters, for example.
var NsTargetSubnetFields *NsTargetSubnetFieldHandles

func init() {
	fieldID := "id"
	fieldLabel := "label"

	NsTargetSubnetFields = &NsTargetSubnetFieldHandles{
		ID:    &fieldID,
		Label: &fieldLabel,
	}
}

// NsTargetSubnet - List of subnet labels.
type NsTargetSubnet struct {
	// ID - Subnet ID.
	ID *string `json:"id,omitempty"`
	// Label - Subnet label.
	Label *string `json:"label,omitempty"`
}

// NsTargetSubnetFieldHandles provides a string representation for each AccessControlRecord field.
type NsTargetSubnetFieldHandles struct {
	ID    *string
	Label *string
}
