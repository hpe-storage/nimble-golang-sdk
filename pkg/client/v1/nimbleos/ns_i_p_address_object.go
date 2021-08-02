// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsIPAddressObjectFields provides field names to use in filter parameters, for example.
var NsIPAddressObjectFields *NsIPAddressObjectFieldHandles

func init() {
	NsIPAddressObjectFields = &NsIPAddressObjectFieldHandles{
		IpAddr: "ip_addr",
	}
}

// NsIPAddressObject - Object wrapper of IP Address.
type NsIPAddressObject struct {
	// IpAddr - An IP Address.
	IpAddr *string `json:"ip_addr,omitempty"`
}

// NsIPAddressObjectFieldHandles provides a string representation for each AccessControlRecord field.
type NsIPAddressObjectFieldHandles struct {
	IpAddr string
}
