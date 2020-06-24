/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsGroupMergeReturn - Response of group merge validation.
// Export NsGroupMergeReturnFields for advance operations like search filter etc.
var NsGroupMergeReturnFields *NsGroupMergeReturn

func init(){
	SrcSIDfield:= "src_sid"
	SrcGroupNamefield:= "src_group_name"
	DstGroupNamefield:= "dst_group_name"
	DstGroupSwversionfield:= "dst_group_swversion"
	ValIDationErrorMsgfield:= "validation_error_msg"
		
	NsGroupMergeReturnFields= &NsGroupMergeReturn{
		SrcSID: &SrcSIDfield,
		SrcGroupName: &SrcGroupNamefield,
		DstGroupName: &DstGroupNamefield,
		DstGroupSwversion: &DstGroupSwversionfield,
		ValIDationErrorMsg: &ValIDationErrorMsgfield,
		
	}
}

type NsGroupMergeReturn struct {
   
    // List of volumes which are online.
    
	OnlineVols []*string `json:"online_vols,omitempty"`
   
    // List of snapshots which are online.
    
   	OnlineSnaps []*NsVolAndSnapName `json:"online_snaps,omitempty"`
   
    // List of partners on groupB which are not paused/stopped.
    
	ActivePartners []*string `json:"active_partners,omitempty"`
   
    // List of partners on groupA that think we are behind a NAT.
    
	DstNatPartners []*string `json:"dst_nat_partners,omitempty"`
   
    // List of partner throttles on groupB.
    
	SrcThrottles []*string `json:"src_throttles,omitempty"`
   
    // List of partner throttles on groupA.
    
	DstThrottles []*string `json:"dst_throttles,omitempty"`
   
    // List of volumes which have the same UID (replicated).
    
   	ReplObjs []*NsReplPairListWithObjectType `json:"repl_objs,omitempty"`
   
    // List of objects with conflicting names.
    
   	NameConflicts []*NsObjectNameListWithType `json:"name_conflicts,omitempty"`
   
    // List of objects with conflicting names. These object need to be resolved manually.
    
   	NameConflictsManualResolve []*NsObjectNameListWithType `json:"name_conflicts_manual_resolve,omitempty"`
   
    // List of objects with conflicting serial numbers.
    
   	SerialConflicts []*NsObjectNameListWithType `json:"serial_conflicts,omitempty"`
   
    // List of objects with conflicting app_uuid.
    
   	AppUuIDConflicts []*NsObjectNameListWithType `json:"app_uuid_conflicts,omitempty"`
   
    // List of objects with conflicting names and their owners.
    
   	NameConflictsAndOwners []*NsObjectOwnerPairWithType `json:"name_conflicts_and_owners,omitempty"`
   
    // List of object types whose total number limit would be violated.
    
   	LimitViolations []*NsObjectLimit `json:"limit_violations,omitempty"`
   
    // List of snapshot retainment param limit violations.
    
   	SnapRetainLimitViolations []*NsSnapRetainLimit `json:"snap_retain_limit_violations,omitempty"`
   
    // List of validation errors with details.
    
   	NetworkErrorList []*NsErrorWithArguments `json:"network_error_list,omitempty"`
   
    // List of validation errors for automatic switchover of Group Management.
    
   	AutoSwitchoverConflicts []*NsErrorWithArguments `json:"auto_switchover_conflicts,omitempty"`
   
    // List of validation errors.
    
   	ErrorList []*NsErrorWithArguments `json:"error_list,omitempty"`
   
    // List of WWPN/alias conflicts (same initiator WWPN, but different aliases).
    
   	AliasConflicts []*NsAliasConflictPair `json:"alias_conflicts,omitempty"`
   
    // List of Fibre Channel interfaces which are online.
    
   	OnlineFcIntfs []*NsFibreChannelInterfaceFullName `json:"online_fc_intfs,omitempty"`
   
    // List of LUN conflicts.
    
   	LunConflicts []*NsLunConflictPair `json:"lun_conflicts,omitempty"`
   
    // Session ID for source group.
    
 	SrcSID *string `json:"src_sid,omitempty"`
   
    // Name of the group we are merging with.
    
 	SrcGroupName *string `json:"src_group_name,omitempty"`
   
    // Name of the destination group.
    
 	DstGroupName *string `json:"dst_group_name,omitempty"`
   
    // Software version of the destination group.
    
 	DstGroupSwversion *string `json:"dst_group_swversion,omitempty"`
   
    // List host type with different bit mask.
    
	HtypeNameConflicts []*string `json:"htype_name_conflicts,omitempty"`
   
    // List of initiators and igroups with different host types.
    
   	HostTypeConflicts []*NsHostType `json:"host_type_conflicts,omitempty"`
   
    // Umbrella error code for group merge validation.
    
   	ValIDationError []*NsErrorWithArguments `json:"validation_error,omitempty"`
   
    // Detailed error message.
    
 	ValIDationErrorMsg *string `json:"validation_error_msg,omitempty"`
}
