// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsRoute - Route configuration.

// Export NsRouteFields provides field names to use in filter parameters, for example.
var NsRouteFields *NsRouteStringFields

func init() {
	fieldTgtNetwork := "tgt_network"
	fieldTgtNetmask := "tgt_netmask"
	fieldGateway := "gateway"

	NsRouteFields = &NsRouteStringFields{
		TgtNetwork: &fieldTgtNetwork,
		TgtNetmask: &fieldTgtNetmask,
		Gateway:    &fieldGateway,
	}
}

type NsRoute struct {
	// TgtNetwork - Target network address.
	TgtNetwork *string `json:"tgt_network,omitempty"`
	// TgtNetmask - Target network mask.
	TgtNetmask *string `json:"tgt_netmask,omitempty"`
	// Gateway - Gateway IP address.
	Gateway *string `json:"gateway,omitempty"`
}

// Struct for NsRouteFields
type NsRouteStringFields struct {
	TgtNetwork *string
	TgtNetmask *string
	Gateway    *string
}
