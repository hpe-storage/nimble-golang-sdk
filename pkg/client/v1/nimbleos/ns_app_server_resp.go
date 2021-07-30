// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAppServerResp - Response from app server.
// Export NsAppServerRespFields for advance operations like search filter etc.
var NsAppServerRespFields *NsAppServerRespStringFields

func init() {
	GeneralErrorfield := "general_error"
	AppSyncfield := "app_sync"
	HasAssocVolsfield := "has_assoc_vols"
	VssResponsefield := "vss_response"
	VmwResponsefield := "vmw_response"
	GenericResponsefield := "generic_response"

	NsAppServerRespFields = &NsAppServerRespStringFields{
		GeneralError:    &GeneralErrorfield,
		AppSync:         &AppSyncfield,
		HasAssocVols:    &HasAssocVolsfield,
		VssResponse:     &VssResponsefield,
		VmwResponse:     &VmwResponsefield,
		GenericResponse: &GenericResponsefield,
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

// Struct for NsAppServerRespFields
type NsAppServerRespStringFields struct {
	GeneralError    *string
	AppSync         *string
	HasAssocVols    *string
	VssResponse     *string
	VmwResponse     *string
	GenericResponse *string
}
