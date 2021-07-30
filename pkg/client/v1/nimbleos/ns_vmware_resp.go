// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVmwareResp - Response from Vmware app server.
// Export NsVmwareRespFields for advance operations like search filter etc.
var NsVmwareRespFields *NsVmwareRespStringFields

func init() {
	VmwareErrorfield := "vmware_error"
	VmwareErrorMessagefield := "vmware_error_message"
	ConnStatusOkfield := "conn_status_ok"
	ConnMessagefield := "conn_message"
	UserRolePermissionOkfield := "user_role_permission_ok"
	VolcollHasVmfield := "volcoll_has_vm"
	NumVmsForVolcollfield := "num_vms_for_volcoll"
	WarningVmToolsStatusfield := "warning_vm_tools_status"

	NsVmwareRespFields = &NsVmwareRespStringFields{
		VmwareError:          &VmwareErrorfield,
		VmwareErrorMessage:   &VmwareErrorMessagefield,
		ConnStatusOk:         &ConnStatusOkfield,
		ConnMessage:          &ConnMessagefield,
		UserRolePermissionOk: &UserRolePermissionOkfield,
		VolcollHasVm:         &VolcollHasVmfield,
		NumVmsForVolcoll:     &NumVmsForVolcollfield,
		WarningVmToolsStatus: &WarningVmToolsStatusfield,
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

// Struct for NsVmwareRespFields
type NsVmwareRespStringFields struct {
	VmwareError          *string
	VmwareErrorMessage   *string
	ConnStatusOk         *string
	ConnMessage          *string
	UserRolePermissionOk *string
	VolcollHasVm         *string
	NumVmsForVolcoll     *string
	WarningVmToolsStatus *string
}
