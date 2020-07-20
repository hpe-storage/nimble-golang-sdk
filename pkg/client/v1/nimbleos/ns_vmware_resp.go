// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVmwareResp - Response from Vmware app server.
// Export NsVmwareRespFields for advance operations like search filter etc.
var NsVmwareRespFields *NsVmwareResp

func init() {

	NsVmwareRespFields = &NsVmwareResp{
		VmwareError:        "vmware_error",
		VmwareErrorMessage: "vmware_error_message",
		ConnMessage:        "conn_message",
	}
}

type NsVmwareResp struct {
	// VmwareError - Error code from Vmware app server.
	VmwareError string `json:"vmware_error,omitempty"`
	// VmwareErrorMessage - Detailed error message from Vmware app server.
	VmwareErrorMessage string `json:"vmware_error_message,omitempty"`
	// ConnStatusOk - Is the connection status OK.
	ConnStatusOk *bool `json:"conn_status_ok,omitempty"`
	// ConnMessage - Detailed connection message.
	ConnMessage string `json:"conn_message,omitempty"`
	// UserRolePermissionOk - Does the user have permission.
	UserRolePermissionOk *bool `json:"user_role_permission_ok,omitempty"`
	// VolcollHasVm - Does this volcoll have a vm.
	VolcollHasVm *bool `json:"volcoll_has_vm,omitempty"`
	// NumVmsForVolcoll - Number of VMs for this volcoll.
	NumVmsForVolcoll int64 `json:"num_vms_for_volcoll,omitempty"`
	// WarningVmToolsStatus - List of status messages one per VM.
	WarningVmToolsStatus []*string `json:"warning_vm_tools_status,omitempty"`
}

// sdk internal struct
type nsVmwareResp struct {
	VmwareError          *string   `json:"vmware_error,omitempty"`
	VmwareErrorMessage   *string   `json:"vmware_error_message,omitempty"`
	ConnStatusOk         *bool     `json:"conn_status_ok,omitempty"`
	ConnMessage          *string   `json:"conn_message,omitempty"`
	UserRolePermissionOk *bool     `json:"user_role_permission_ok,omitempty"`
	VolcollHasVm         *bool     `json:"volcoll_has_vm,omitempty"`
	NumVmsForVolcoll     *int64    `json:"num_vms_for_volcoll,omitempty"`
	WarningVmToolsStatus []*string `json:"warning_vm_tools_status,omitempty"`
}

// EncodeNsVmwareResp - Transform NsVmwareResp to nsVmwareResp type
func EncodeNsVmwareResp(request interface{}) (*nsVmwareResp, error) {
	reqNsVmwareResp := request.(*NsVmwareResp)
	byte, err := json.Marshal(reqNsVmwareResp)
	if err != nil {
		return nil, err
	}
	respNsVmwareRespPtr := &nsVmwareResp{}
	err = json.Unmarshal(byte, respNsVmwareRespPtr)
	return respNsVmwareRespPtr, err
}

// DecodeNsVmwareResp Transform nsVmwareResp to NsVmwareResp type
func DecodeNsVmwareResp(request interface{}) (*NsVmwareResp, error) {
	reqNsVmwareResp := request.(*nsVmwareResp)
	byte, err := json.Marshal(reqNsVmwareResp)
	if err != nil {
		return nil, err
	}
	respNsVmwareResp := &NsVmwareResp{}
	err = json.Unmarshal(byte, respNsVmwareResp)
	return respNsVmwareResp, err
}
