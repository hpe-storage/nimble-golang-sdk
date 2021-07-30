// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsTargetSubnet - List of subnet labels.

// Export NsTargetSubnetFields provides field names to use in filter parameters, for example.
var NsTargetSubnetFields *NsTargetSubnetStringFields

func init() {
	fieldID := "id"
	fieldLabel := "label"

	NsTargetSubnetFields = &NsTargetSubnetStringFields{
		ID:    &fieldID,
		Label: &fieldLabel,
	}
}

type NsTargetSubnet struct {
	// ID - Subnet ID.
	ID *string `json:"id,omitempty"`
	// Label - Subnet label.
	Label *string `json:"label,omitempty"`
}

// Struct for NsTargetSubnetFields
type NsTargetSubnetStringFields struct {
	ID    *string
	Label *string
}
