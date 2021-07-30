// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsNIC - Network interface controller.

// Export NsNICFields provides field names to use in filter parameters, for example.
var NsNICFields *NsNICStringFields

func init() {
	fieldName := "name"
	fieldSubnetLabel := "subnet_label"
	fieldDataIp := "data_ip"
	fieldTagged := "tagged"

	NsNICFields = &NsNICStringFields{
		Name:        &fieldName,
		SubnetLabel: &fieldSubnetLabel,
		DataIp:      &fieldDataIp,
		Tagged:      &fieldTagged,
	}
}

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

// Struct for NsNICFields
type NsNICStringFields struct {
	Name        *string
	SubnetLabel *string
	DataIp      *string
	Tagged      *string
}
