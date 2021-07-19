// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsTargetSubnet - List of subnet labels.
// Export NsTargetSubnetFields for advance operations like search filter etc.
var NsTargetSubnetFields *NsTargetSubnet

func init() {
	IDfield := "id"
	Labelfield := "label"

	NsTargetSubnetFields = &NsTargetSubnet{
		ID:    &IDfield,
		Label: &Labelfield,
	}
}

type NsTargetSubnet struct {
	// ID - Subnet ID.
	ID *string `json:"id,omitempty"`
	// Label - Subnet label.
	Label *string `json:"label,omitempty"`
}
