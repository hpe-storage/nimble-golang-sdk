// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsRouteFields provides field names to use in filter parameters, for example.
var NsRouteFields *NsRouteFieldHandles

func init() {
	NsRouteFields = &NsRouteFieldHandles{
		TgtNetwork: "tgt_network",
		TgtNetmask: "tgt_netmask",
		Gateway:    "gateway",
	}
}

// NsRoute - Route configuration.
type NsRoute struct {
	// TgtNetwork - Target network address.
	TgtNetwork *string `json:"tgt_network,omitempty"`
	// TgtNetmask - Target network mask.
	TgtNetmask *string `json:"tgt_netmask,omitempty"`
	// Gateway - Gateway IP address.
	Gateway *string `json:"gateway,omitempty"`
}

// NsRouteFieldHandles provides a string representation for each AccessControlRecord field.
type NsRouteFieldHandles struct {
	TgtNetwork string
	TgtNetmask string
	Gateway    string
}
