// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsRoute - Route configuration.
// Export NsRouteFields for advance operations like search filter etc.
var NsRouteFields *NsRoute

func init(){
	TgtNetworkfield:= "tgt_network"
	TgtNetmaskfield:= "tgt_netmask"
	Gatewayfield:= "gateway"
		
	NsRouteFields= &NsRoute{
	TgtNetwork: &TgtNetworkfield,
	TgtNetmask: &TgtNetmaskfield,
	Gateway: &Gatewayfield,
		
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
