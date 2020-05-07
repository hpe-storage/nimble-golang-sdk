/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsVmwareResp


// NsVmwareResp :
type NsVmwareResp struct {
   // VmwareError
   VmwareError string `json:"vmware_error,omitempty"`
   // VmwareErrorMessage
   VmwareErrorMessage string `json:"vmware_error_message,omitempty"`
   // ConnStatusOk
   ConnStatusOk bool `json:"conn_status_ok,omitempty"`
   // ConnMessage
   ConnMessage string `json:"conn_message,omitempty"`
   // UserRolePermissionOk
   UserRolePermissionOk bool `json:"user_role_permission_ok,omitempty"`
   // VolcollHasVm
   VolcollHasVm bool `json:"volcoll_has_vm,omitempty"`
   // NumVmsForVolcoll
   NumVmsForVolcoll float64 `json:"num_vms_for_volcoll,omitempty"`
}
