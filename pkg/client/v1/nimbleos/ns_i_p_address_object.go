// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsIPAddressObject - Object wrapper of IP Address.

// Export NsIPAddressObjectFields provides field names to use in filter parameters, for example.
var NsIPAddressObjectFields *NsIPAddressObjectStringFields

func init() {
	fieldIpAddr := "ip_addr"

	NsIPAddressObjectFields = &NsIPAddressObjectStringFields{
		IpAddr: &fieldIpAddr,
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
