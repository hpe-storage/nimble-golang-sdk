// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsIPAddressObjectFields provides field names to use in filter parameters, for example.
var NsIPAddressObjectFields *NsIPAddressObjectFieldHandles

func init() {
	fieldIpAddr := "ip_addr"

	NsIPAddressObjectFields = &NsIPAddressObjectFieldHandles{
		IpAddr: &fieldIpAddr,
	}
}

// NsIPAddressObject - Object wrapper of IP Address.
type NsIPAddressObject struct {
	// IpAddr - An IP Address.
	IpAddr *string `json:"ip_addr,omitempty"`
}

// NsIPAddressObjectFieldHandles provides a string representation for each AccessControlRecord field.
type NsIPAddressObjectFieldHandles struct {
	IpAddr *string
}
