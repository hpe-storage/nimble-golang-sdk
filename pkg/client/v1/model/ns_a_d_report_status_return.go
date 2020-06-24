/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsADReportStatusReturn - Status of the Active Directory domain.
// Export NsADReportStatusReturnFields for advance operations like search filter etc.
var NsADReportStatusReturnFields *NsADReportStatusReturn

func init(){
		
	NsADReportStatusReturnFields= &NsADReportStatusReturn{
		
	}
}

type NsADReportStatusReturn struct {
   
    // Joined the Active Directory group.
    
 	Joined *bool `json:"joined,omitempty"`
   
    // Active Directory group is enabled.
    
 	Enabled *bool `json:"enabled,omitempty"`
   
    // Status of the local service.
    
 	LocalServiceStatus *bool `json:"local_service_status,omitempty"`
   
    // Status of the remote service.
    
 	RemoteServiceStatus *bool `json:"remote_service_status,omitempty"`
   
    // Trust is valid.
    
 	TrustValID *bool `json:"trust_valid,omitempty"`
}
