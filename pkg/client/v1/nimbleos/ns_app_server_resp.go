// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAppServerResp - Response from app server.

// Export NsAppServerRespFields provides field names to use in filter parameters, for example.
var NsAppServerRespFields *NsAppServerRespStringFields

func init() {
	fieldGeneralError := "general_error"
	fieldAppSync := "app_sync"
	fieldHasAssocVols := "has_assoc_vols"
	fieldVssResponse := "vss_response"
	fieldVmwResponse := "vmw_response"
	fieldGenericResponse := "generic_response"

	NsAppServerRespFields = &NsAppServerRespStringFields{
		GeneralError:    &fieldGeneralError,
		AppSync:         &fieldAppSync,
		HasAssocVols:    &fieldHasAssocVols,
		VssResponse:     &fieldVssResponse,
		VmwResponse:     &fieldVmwResponse,
		GenericResponse: &fieldGenericResponse,
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
