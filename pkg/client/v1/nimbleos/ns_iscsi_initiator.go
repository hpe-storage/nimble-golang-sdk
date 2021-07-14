// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsISCSIInitiator - ISCSI initiator.
// Export NsISCSIInitiatorFields for advance operations like search filter etc.
var NsISCSIInitiatorFields *NsISCSIInitiator

func init(){
 IDfield:= "id"
 InitiatorIdfield:= "initiator_id"
 Labelfield:= "label"
 Iqnfield:= "iqn"
 IpAddressfield:= "ip_address"

 NsISCSIInitiatorFields= &NsISCSIInitiator{
  ID:          &IDfield,
  InitiatorId: &InitiatorIdfield,
  Label:       &Labelfield,
  Iqn:         &Iqnfield,
  IpAddress:   &IpAddressfield,
 }
}


type NsISCSIInitiator struct {
 // ID - Unique identifier of the iSCSI initiator.
  ID *string `json:"id,omitempty"`
 // InitiatorId - Unique identifier of the iSCSI initiator.
  InitiatorId *string `json:"initiator_id,omitempty"`
 // Label - Unique label of the iSCSI initiator.
  Label *string `json:"label,omitempty"`
 // Iqn - IQN name of the iSCSI initiator.
  Iqn *string `json:"iqn,omitempty"`
 // IpAddress - IP address of the iSCSI initiator.
  IpAddress *string `json:"ip_address,omitempty"`
}
