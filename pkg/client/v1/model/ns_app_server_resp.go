/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsAppServerResp - Response from app server.
// Export NsAppServerRespFields for advance operations like search filter etc.
var NsAppServerRespFields *NsAppServerResp

func init(){
	GeneralErrorfield:= "general_error"
		
	NsAppServerRespFields= &NsAppServerResp{
		GeneralError: &GeneralErrorfield,
		
	}
}

type NsAppServerResp struct {
   
    // Error code from app server.
    
 	GeneralError *string `json:"general_error,omitempty"`
   
    // Type of app server.
    
   	AppSync *NsAppSyncType `json:"app_sync,omitempty"`
   
    // Indicates if there are associated volumes.
    
 	HasAssocVols *bool `json:"has_assoc_vols,omitempty"`
   
    // Response from vss app server.
    
	VssResponse *NsVssResp `json:"vss_response,omitempty"`
   
    // Response from vmware app server.
    
	VmwResponse *NsVmwareResp `json:"vmw_response,omitempty"`
   
    // Response from generic app server.
    
	GenericResponse *NsGenericResp `json:"generic_response,omitempty"`
}
