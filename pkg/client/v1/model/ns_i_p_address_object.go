/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsIPAddressObject - Object wrapper of IP Address.
// Export NsIPAddressObjectFields for advance operations like search filter etc.
var NsIPAddressObjectFields *NsIPAddressObject

func init(){
	IpAddrfield:= "ip_addr"
		
	NsIPAddressObjectFields= &NsIPAddressObject{
		IpAddr: &IpAddrfield,
		
	}
}

type NsIPAddressObject struct {
   
    // An IP Address.
    
 	IpAddr *string `json:"ip_addr,omitempty"`
}
