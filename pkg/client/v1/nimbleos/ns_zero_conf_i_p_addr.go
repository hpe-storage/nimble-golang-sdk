// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsZeroConfIPAddr - Zero Conf of array.

// Export NsZeroConfIPAddrFields provides field names to use in filter parameters, for example.
var NsZeroConfIPAddrFields *NsZeroConfIPAddrStringFields

func init() {
	fieldNic := "nic"
	fieldLocalIpaddr := "local_ipaddr"
	fieldRemoteIpaddr := "remote_ipaddr"

	NsZeroConfIPAddrFields = &NsZeroConfIPAddrStringFields{
		Nic:          &fieldNic,
		LocalIpaddr:  &fieldLocalIpaddr,
		RemoteIpaddr: &fieldRemoteIpaddr,
	}
}

type NsZeroConfIPAddr struct {
	// Nic - Nic of array.
	Nic *string `json:"nic,omitempty"`
	// LocalIpaddr - Local IP address of array.
	LocalIpaddr *string `json:"local_ipaddr,omitempty"`
	// RemoteIpaddr - Remote IP address of array.
	RemoteIpaddr *string `json:"remote_ipaddr,omitempty"`
}

// Struct for NsZeroConfIPAddrFields
type NsZeroConfIPAddrStringFields struct {
	Nic          *string
	LocalIpaddr  *string
	RemoteIpaddr *string
}
