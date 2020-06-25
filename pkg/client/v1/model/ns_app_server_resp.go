// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsAppServerResp - Response from app server.
// Export NsAppServerRespFields for advance operations like search filter etc.
var NsAppServerRespFields *NsAppServerResp

func init(){
	GeneralErrorfield:= "general_error"
		
	NsAppServerRespFields= &NsAppServerResp{
		GeneralError:    &GeneralErrorfield,
	}
}

type NsAppServerResp struct {
	// GeneralError - Error code from app server.
 	GeneralError *string `json:"general_error,omitempty"`
	// AppSync - Type of app server.
   	AppSync *NsAppSyncType `json:"app_sync,omitempty"`
	// HasAssocVols - Indicates if there are associated volumes.
    HasAssocVols *bool `json:"has_assoc_vols,omitempty"`
	// VssResponse - Response from vss app server.
	VssResponse *NsVssResp `json:"vss_response,omitempty"`
	// VmwResponse - Response from vmware app server.
	VmwResponse *NsVmwareResp `json:"vmw_response,omitempty"`
	// GenericResponse - Response from generic app server.
	GenericResponse *NsGenericResp `json:"generic_response,omitempty"`
}
