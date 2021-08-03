// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsZeroConfIPAddrFields provides field names to use in filter parameters, for example.
var NsZeroConfIPAddrFields *NsZeroConfIPAddrFieldHandles

func init() {
	NsZeroConfIPAddrFields = &NsZeroConfIPAddrFieldHandles{
		Nic:          "nic",
		LocalIpaddr:  "local_ipaddr",
		RemoteIpaddr: "remote_ipaddr",
	}
}

// NsZeroConfIPAddr - Zero Conf of array.
type NsZeroConfIPAddr struct {
	// Nic - Nic of array.
	Nic *string `json:"nic,omitempty"`
	// LocalIpaddr - Local IP address of array.
	LocalIpaddr *string `json:"local_ipaddr,omitempty"`
	// RemoteIpaddr - Remote IP address of array.
	RemoteIpaddr *string `json:"remote_ipaddr,omitempty"`
}

// NsZeroConfIPAddrFieldHandles provides a string representation for each NsZeroConfIPAddr field.
type NsZeroConfIPAddrFieldHandles struct {
	Nic          string
	LocalIpaddr  string
	RemoteIpaddr string
}
