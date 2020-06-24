/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Initiator - Manage initiators in initiator groups. An initiator group has a set of initiators that can be configured as part of your ACL to access a specific volume through group membership.
// Export InitiatorFields for advance operations like search filter etc.
var InitiatorFields *Initiator

func init(){
	IDfield:= "id"
	InitiatorGroupIDfield:= "initiator_group_id"
	InitiatorGroupNamefield:= "initiator_group_name"
	Labelfield:= "label"
	Iqnfield:= "iqn"
	IpAddressfield:= "ip_address"
	Aliasfield:= "alias"
	ChapuserIDfield:= "chapuser_id"
	Wwpnfield:= "wwpn"
		
	InitiatorFields= &Initiator{
		ID: &IDfield,
		InitiatorGroupID: &InitiatorGroupIDfield,
		InitiatorGroupName: &InitiatorGroupNamefield,
		Label: &Labelfield,
		Iqn: &Iqnfield,
		IpAddress: &IpAddressfield,
		Alias: &Aliasfield,
		ChapuserID: &ChapuserIDfield,
		Wwpn: &Wwpnfield,
		
	}
}

type Initiator struct {
   
    // Identifier for initiator.
    
 	ID *string `json:"id,omitempty"`
   
    // Access protocol used by the initiator. Valid values are 'iscsi' and 'fc'.
    
   	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
   
    // Identifier of the initiator group that this initiator is assigned to.
    
 	InitiatorGroupID *string `json:"initiator_group_id,omitempty"`
   
    // Name of the initiator group that this initiator is assigned to.
    
 	InitiatorGroupName *string `json:"initiator_group_name,omitempty"`
   
    // Unique Identifier of the iSCSI initiator. Label is required when creating iSCSI initiator.
    
 	Label *string `json:"label,omitempty"`
   
    // IQN name of the iSCSI initiator. Each initiator IQN name must have an associated IP address specified using the 'ip_address' attribute. You can choose not to enter the IP address for an initiator if you prefer not to authenticate using both name and IP address, in this case the IP address will be returned as '*'.
    
 	Iqn *string `json:"iqn,omitempty"`
   
    // IP address of the iSCSI initiator. Each initiator IP address must have an associated name specified using 'name' attribute. You can choose not to enter the name for an initiator if you prefer not to authenticate using both name and IP address, in this case the IQN name will be returned as '*'.
    
 	IpAddress *string `json:"ip_address,omitempty"`
   
    // Alias of the Fibre Channel initiator. Maximum alias length is 32 characters. Each initiator alias must have an associated WWPN specified using the 'wwpn' attribute. You can choose not to enter the WWPN for an initiator when using previously saved initiator alias.
    
 	Alias *string `json:"alias,omitempty"`
   
    // Identifier for the CHAP user.
    
 	ChapuserID *string `json:"chapuser_id,omitempty"`
   
    // WWPN (World Wide Port Name) of the Fibre Channel initiator. WWPN is required when creating a Fibre Channel initiator. Each initiator WWPN can have an associated alias specified using the 'alias' attribute. You can choose not to enter the alias for an initiator if you prefer not to assign an initiator alias.
    
 	Wwpn *string `json:"wwpn,omitempty"`
   
    // Flag to allow modifying VP created initiator groups. When set to true, user can add this initiator to a VP created initiator group.
    
 	VpOverrIDe *bool `json:"vp_override,omitempty"`
   
    // Time when this initiator group was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this initiator group was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Forcibly add Fibre Channel initiator to initiator group by updating or removing conflicting Fibre Channel initiator aliases.
    
 	OverrIDeExistingAlias *bool `json:"override_existing_alias,omitempty"`
}
