// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsVmwareResp - Response from Vmware app server.
// Export NsVmwareRespFields for advance operations like search filter etc.
var NsVmwareRespFields *NsVmwareResp

func init(){
	VmwareErrorfield:= "vmware_error"
	VmwareErrorMessagefield:= "vmware_error_message"
	ConnMessagefield:= "conn_message"
		
	NsVmwareRespFields= &NsVmwareResp{
		VmwareError:            &VmwareErrorfield,
		VmwareErrorMessage:     &VmwareErrorMessagefield,
		ConnMessage:            &ConnMessagefield,
	}
}

type NsVmwareResp struct {
	// VmwareError - Error code from Vmware app server.
 	VmwareError *string `json:"vmware_error,omitempty"`
	// VmwareErrorMessage - Detailed error message from Vmware app server.
 	VmwareErrorMessage *string `json:"vmware_error_message,omitempty"`
	// ConnStatusOk - Is the connection status OK.
    ConnStatusOk *bool `json:"conn_status_ok,omitempty"`
	// ConnMessage - Detailed connection message.
 	ConnMessage *string `json:"conn_message,omitempty"`
	// UserRolePermissionOk - Does the user have permission.
    UserRolePermissionOk *bool `json:"user_role_permission_ok,omitempty"`
	// VolcollHasVm - Does this volcoll have a vm.
    VolcollHasVm *bool `json:"volcoll_has_vm,omitempty"`
	// NumVmsForVolcoll - Number of VMs for this volcoll.
   	NumVmsForVolcoll *int64 `json:"num_vms_for_volcoll,omitempty"`
	// WarningVmToolsStatus - List of status messages one per VM.
	WarningVmToolsStatus []*string `json:"warning_vm_tools_status,omitempty"`
}
