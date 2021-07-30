// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsNICFields provides field names to use in filter parameters, for example.
var NsNICFields *NsNICFieldHandles

func init() {
	fieldName := "name"
	fieldSubnetLabel := "subnet_label"
	fieldDataIp := "data_ip"
	fieldTagged := "tagged"

	NsNICFields = &NsNICFieldHandles{
		Name:        &fieldName,
		SubnetLabel: &fieldSubnetLabel,
		DataIp:      &fieldDataIp,
		Tagged:      &fieldTagged,
	}
}

// NsNIC - Network interface controller.
type NsNIC struct {
	// Name - Name of NIC.
	Name *string `json:"name,omitempty"`
	// SubnetLabel - Subnet label for this NIC.
	SubnetLabel *string `json:"subnet_label,omitempty"`
	// DataIp - Data IP address.
	DataIp *string `json:"data_ip,omitempty"`
	// Tagged - Identify whether the NIC is tagged.
	Tagged *bool `json:"tagged,omitempty"`
}

// NsNICFieldHandles provides a string representation for each AccessControlRecord field.
type NsNICFieldHandles struct {
	Name        *string
	SubnetLabel *string
	DataIp      *string
	Tagged      *string
}
