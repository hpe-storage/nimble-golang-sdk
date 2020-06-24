/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// InitiatorGroup - Manage initiator groups for initiator authentication. An initiator group is a set of initiators that can be configured as part of your ACL to access a specific volume through group membership.
// Export InitiatorGroupFields for advance operations like search filter etc.
var InitiatorGroupFields *InitiatorGroup

func init(){
	IDfield:= "id"
	Namefield:= "name"
	FullNamefield:= "full_name"
	SearchNamefield:= "search_name"
	Descriptionfield:= "description"
	HostTypefield:= "host_type"
	AppUuIDfield:= "app_uuid"
		
	InitiatorGroupFields= &InitiatorGroup{
		ID: &IDfield,
		Name: &Namefield,
		FullName: &FullNamefield,
		SearchName: &SearchNamefield,
		Description: &Descriptionfield,
		HostType: &HostTypefield,
		AppUuID: &AppUuIDfield,
		
	}
}

type InitiatorGroup struct {
   
    // Identifier for initiator group.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of initiator group.
    
 	Name *string `json:"name,omitempty"`
   
    // Initiator group's full name.
    
 	FullName *string `json:"full_name,omitempty"`
   
    // Initiator group name used for search.
    
 	SearchName *string `json:"search_name,omitempty"`
   
    // Text description of initiator group.
    
 	Description *string `json:"description,omitempty"`
   
    // Initiator group access protocol.
    
   	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
   
    // Initiator group host type. Available options are auto and hpux. The default option is auto. This attribute will be applied to all the initiators in the initiator group. Initiators with different host OSes should not be kept in the same initiator group having a non-default host type attribute.
    
 	HostType *string `json:"host_type,omitempty"`
   
    // List of target Fibre Channel ports with Target Driven Zoning configured on this initiator group.
    
   	FcTdzPorts []*NsFcTdzPort `json:"fc_tdz_ports,omitempty"`
   
    // List of target subnet labels. If specified, discovery and access to volumes will be restricted to the specified subnets.
    
   	TargetSubnets []*NsTargetSubnet `json:"target_subnets,omitempty"`
   
    // List of iSCSI initiators. When create/update iscsi_initiators, either iqn or ip_address is always required with label.
    
   	IscsiInitiators []*NsISCSIInitiator `json:"iscsi_initiators,omitempty"`
   
    // List of FC initiators. When create/update fc_initiators, wwpn is required.
    
   	FcInitiators []*NsFCInitiator `json:"fc_initiators,omitempty"`
   
    // Time when this initiator group was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this initiator group was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Flag to allow modifying VP created initiator groups. When set to true, user can add this initiator to a VP created initiator group.
    
 	VpOverrIDe *bool `json:"vp_override,omitempty"`
   
    // Application identifier of initiator group.
    
 	AppUuID *string `json:"app_uuid,omitempty"`
   
    // Number of volumes that are accessible by the initiator group.
    
   	VolumeCount *int64 `json:"volume_count,omitempty"`
   
    // List of volumes that are accessible by the initiator group.
    
   	VolumeList []*NsVolumeSummaryWithAppCategory `json:"volume_list,omitempty"`
   
    // Total number of connections from initiators in the initiator group.
    
   	NumConnections *int64 `json:"num_connections,omitempty"`
   
    // Key-value pairs that augment an initiator group's attributes.
    
   	Metadata []*NsKeyValue `json:"metadata,omitempty"`
}
