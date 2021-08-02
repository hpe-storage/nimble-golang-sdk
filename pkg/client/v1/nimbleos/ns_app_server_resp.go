// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAppServerRespFields provides field names to use in filter parameters, for example.
var NsAppServerRespFields *NsAppServerRespFieldHandles

func init() {
	NsAppServerRespFields = &NsAppServerRespFieldHandles{
		GeneralError:    "general_error",
		AppSync:         "app_sync",
		HasAssocVols:    "has_assoc_vols",
		VssResponse:     "vss_response",
		VmwResponse:     "vmw_response",
		GenericResponse: "generic_response",
	}
}

// NsAppServerResp - Response from app server.
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

// NsAppServerRespFieldHandles provides a string representation for each AccessControlRecord field.
type NsAppServerRespFieldHandles struct {
	GeneralError    string
	AppSync         string
	HasAssocVols    string
	VssResponse     string
	VmwResponse     string
	GenericResponse string
}
