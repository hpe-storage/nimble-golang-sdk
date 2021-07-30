// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsRoute - Route configuration.
// Export NsRouteFields for advance operations like search filter etc.
var NsRouteFields *NsRouteStringFields

func init() {
	TgtNetworkfield := "tgt_network"
	TgtNetmaskfield := "tgt_netmask"
	Gatewayfield := "gateway"

	NsRouteFields = &NsRouteStringFields{
		TgtNetwork: &TgtNetworkfield,
		TgtNetmask: &TgtNetmaskfield,
		Gateway:    &Gatewayfield,
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
