// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsVmwareRespFields provides field names to use in filter parameters, for example.
var NsVmwareRespFields *NsVmwareRespFieldHandles

func init() {
	NsVmwareRespFields = &NsVmwareRespFieldHandles{
		VmwareError:          "vmware_error",
		VmwareErrorMessage:   "vmware_error_message",
		ConnStatusOk:         "conn_status_ok",
		ConnMessage:          "conn_message",
		UserRolePermissionOk: "user_role_permission_ok",
		VolcollHasVm:         "volcoll_has_vm",
		NumVmsForVolcoll:     "num_vms_for_volcoll",
		WarningVmToolsStatus: "warning_vm_tools_status",
	}
}

// NsVmwareResp - Response from Vmware app server.
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

// NsVmwareRespFieldHandles provides a string representation for each AccessControlRecord field.
type NsVmwareRespFieldHandles struct {
	VmwareError          string
	VmwareErrorMessage   string
	ConnStatusOk         string
	ConnMessage          string
	UserRolePermissionOk string
	VolcollHasVm         string
	NumVmsForVolcoll     string
	WarningVmToolsStatus string
}
