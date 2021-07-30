// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsIPAddressObject - Object wrapper of IP Address.
// Export NsIPAddressObjectFields for advance operations like search filter etc.
var NsIPAddressObjectFields *NsIPAddressObjectStringFields

func init() {
	IpAddrfield := "ip_addr"

	NsIPAddressObjectFields = &NsIPAddressObjectStringFields{
		IpAddr: &IpAddrfield,
	}
}

type NsIPAddressObject struct {
	// IpAddr - An IP Address.
	IpAddr *string `json:"ip_addr,omitempty"`
}

// Struct for NsIPAddressObjectFields
type NsIPAddressObjectStringFields struct {
	IpAddr *string
}
