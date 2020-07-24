// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsIPAddressObject - Object wrapper of IP Address.
// Export NsIPAddressObjectFields for advance operations like search filter etc.
var NsIPAddressObjectFields *NsIPAddressObject

func init() {
	IpAddrfield := "ip_addr"

	NsIPAddressObjectFields = &NsIPAddressObject{
		IpAddr: &IpAddrfield,
	}
}

type NsIPAddressObject struct {
	// IpAddr - An IP Address.
	IpAddr *string `json:"ip_addr,omitempty"`
}
