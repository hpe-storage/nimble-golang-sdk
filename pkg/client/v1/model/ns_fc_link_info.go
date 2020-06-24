/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsFcLinkInfo - Fibre Channel link information.
// Export NsFcLinkInfoFields for advance operations like search filter etc.
var NsFcLinkInfoFields *NsFcLinkInfo

func init(){
		
	NsFcLinkInfoFields= &NsFcLinkInfo{
		
	}
}

type NsFcLinkInfo struct {
   
    // Speed of the Fibre Channel link.
    
   	LinkSpeed *NsPlatFcLinkSpeed `json:"link_speed,omitempty"`
   
    // Maximum speed of the Fibre Channel link.
    
   	MaxLinkSpeed *NsPlatFcLinkSpeed `json:"max_link_speed,omitempty"`
   
    // Status of the Fibre Channel link.
    
   	LinkStatus *NsPlatFcLinkStatus `json:"link_status,omitempty"`
   
    // Operational status of the Fibre Channel link.
    
   	OperationalStatus *NsPlatFcOperationalStatus `json:"operational_status,omitempty"`
}
