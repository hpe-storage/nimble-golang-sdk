// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsAppServerResp - Response from app server.
// Export NsAppServerRespFields for advance operations like search filter etc.
var NsAppServerRespFields *NsAppServerResp

func init() {

	NsAppServerRespFields = &NsAppServerResp{
		GeneralError: "general_error",
	}
}

type NsAppServerResp struct {
	// GeneralError - Error code from app server.
	GeneralError string `json:"general_error,omitempty"`
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

// sdk internal struct
type nsAppServerResp struct {
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

// EncodeNsAppServerResp - Transform NsAppServerResp to nsAppServerResp type
func EncodeNsAppServerResp(request interface{}) (*nsAppServerResp, error) {
	reqNsAppServerResp := request.(*NsAppServerResp)
	byte, err := json.Marshal(reqNsAppServerResp)
	objPtr := &nsAppServerResp{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsAppServerResp Transform nsAppServerResp to NsAppServerResp type
func DecodeNsAppServerResp(request interface{}) (*NsAppServerResp, error) {
	reqNsAppServerResp := request.(*nsAppServerResp)
	byte, err := json.Marshal(reqNsAppServerResp)
	obj := &NsAppServerResp{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
