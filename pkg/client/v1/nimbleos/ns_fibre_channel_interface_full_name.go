// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsFibreChannelInterfaceFullName - Fully qualified name information for a Fibre Channel interface (array/controller/interface).
// Export NsFibreChannelInterfaceFullNameFields for advance operations like search filter etc.
var NsFibreChannelInterfaceFullNameFields *NsFibreChannelInterfaceFullName

func init(){
 ArrayNamefield:= "array_name"
 CtrlrNamefield:= "ctrlr_name"
 IntfNamefield:= "intf_name"

 NsFibreChannelInterfaceFullNameFields= &NsFibreChannelInterfaceFullName{
  ArrayName: &ArrayNamefield,
  CtrlrName: &CtrlrNamefield,
  IntfName:  &IntfNamefield,
 }
}


type NsFibreChannelInterfaceFullName struct {
 // ArrayName - Array name.
  ArrayName *string `json:"array_name,omitempty"`
 // CtrlrName - Controller name.
  CtrlrName *string `json:"ctrlr_name,omitempty"`
 // IntfName - Fibre Channel interface name.
  IntfName *string `json:"intf_name,omitempty"`
}
