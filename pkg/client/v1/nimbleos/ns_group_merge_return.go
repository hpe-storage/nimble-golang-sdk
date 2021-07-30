// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsGroupMergeReturn - Response of group merge validation.

// Export NsGroupMergeReturnFields provides field names to use in filter parameters, for example.
var NsGroupMergeReturnFields *NsGroupMergeReturnStringFields

func init() {
	fieldOnlineVols := "online_vols"
	fieldOnlineSnaps := "online_snaps"
	fieldActivePartners := "active_partners"
	fieldDstNatPartners := "dst_nat_partners"
	fieldSrcThrottles := "src_throttles"
	fieldDstThrottles := "dst_throttles"
	fieldReplObjs := "repl_objs"
	fieldNameConflicts := "name_conflicts"
	fieldNameConflictsManualResolve := "name_conflicts_manual_resolve"
	fieldSerialConflicts := "serial_conflicts"
	fieldAppUuidConflicts := "app_uuid_conflicts"
	fieldNameConflictsAndOwners := "name_conflicts_and_owners"
	fieldLimitViolations := "limit_violations"
	fieldSnapRetainLimitViolations := "snap_retain_limit_violations"
	fieldNetworkErrorList := "network_error_list"
	fieldAutoSwitchoverConflicts := "auto_switchover_conflicts"
	fieldErrorList := "error_list"
	fieldAliasConflicts := "alias_conflicts"
	fieldOnlineFcIntfs := "online_fc_intfs"
	fieldLunConflicts := "lun_conflicts"
	fieldSrcSid := "src_sid"
	fieldSrcGroupName := "src_group_name"
	fieldDstGroupName := "dst_group_name"
	fieldDstGroupSwversion := "dst_group_swversion"
	fieldHtypeNameConflicts := "htype_name_conflicts"
	fieldHostTypeConflicts := "host_type_conflicts"
	fieldValidationError := "validation_error"
	fieldValidationErrorMsg := "validation_error_msg"
	fieldWarningList := "warning_list"

	NsGroupMergeReturnFields = &NsGroupMergeReturnStringFields{
		OnlineVols:                 &fieldOnlineVols,
		OnlineSnaps:                &fieldOnlineSnaps,
		ActivePartners:             &fieldActivePartners,
		DstNatPartners:             &fieldDstNatPartners,
		SrcThrottles:               &fieldSrcThrottles,
		DstThrottles:               &fieldDstThrottles,
		ReplObjs:                   &fieldReplObjs,
		NameConflicts:              &fieldNameConflicts,
		NameConflictsManualResolve: &fieldNameConflictsManualResolve,
		SerialConflicts:            &fieldSerialConflicts,
		AppUuidConflicts:           &fieldAppUuidConflicts,
		NameConflictsAndOwners:     &fieldNameConflictsAndOwners,
		LimitViolations:            &fieldLimitViolations,
		SnapRetainLimitViolations:  &fieldSnapRetainLimitViolations,
		NetworkErrorList:           &fieldNetworkErrorList,
		AutoSwitchoverConflicts:    &fieldAutoSwitchoverConflicts,
		ErrorList:                  &fieldErrorList,
		AliasConflicts:             &fieldAliasConflicts,
		OnlineFcIntfs:              &fieldOnlineFcIntfs,
		LunConflicts:               &fieldLunConflicts,
		SrcSid:                     &fieldSrcSid,
		SrcGroupName:               &fieldSrcGroupName,
		DstGroupName:               &fieldDstGroupName,
		DstGroupSwversion:          &fieldDstGroupSwversion,
		HtypeNameConflicts:         &fieldHtypeNameConflicts,
		HostTypeConflicts:          &fieldHostTypeConflicts,
		ValidationError:            &fieldValidationError,
		ValidationErrorMsg:         &fieldValidationErrorMsg,
		WarningList:                &fieldWarningList,
	}
}

type NsGroupMergeReturn struct {
	// OnlineVols - List of volumes which are online.
	OnlineVols []*string `json:"online_vols,omitempty"`
	// OnlineSnaps - List of snapshots which are online.
	OnlineSnaps []*NsVolAndSnapName `json:"online_snaps,omitempty"`
	// ActivePartners - List of partners on groupB which are not paused/stopped.
	ActivePartners []*string `json:"active_partners,omitempty"`
	// DstNatPartners - List of partners on groupA that think we are behind a NAT.
	DstNatPartners []*string `json:"dst_nat_partners,omitempty"`
	// SrcThrottles - List of partner throttles on groupB.
	SrcThrottles []*string `json:"src_throttles,omitempty"`
	// DstThrottles - List of partner throttles on groupA.
	DstThrottles []*string `json:"dst_throttles,omitempty"`
	// ReplObjs - List of volumes which have the same UID (replicated).
	ReplObjs []*NsReplPairListWithObjectType `json:"repl_objs,omitempty"`
	// NameConflicts - List of objects with conflicting names.
	NameConflicts []*NsObjectNameListWithType `json:"name_conflicts,omitempty"`
	// NameConflictsManualResolve - List of objects with conflicting names. These object need to be resolved manually.
	NameConflictsManualResolve []*NsObjectNameListWithType `json:"name_conflicts_manual_resolve,omitempty"`
	// SerialConflicts - List of objects with conflicting serial numbers.
	SerialConflicts []*NsObjectNameListWithType `json:"serial_conflicts,omitempty"`
	// AppUuidConflicts - List of objects with conflicting app_uuid.
	AppUuidConflicts []*NsObjectNameListWithType `json:"app_uuid_conflicts,omitempty"`
	// NameConflictsAndOwners - List of objects with conflicting names and their owners.
	NameConflictsAndOwners []*NsObjectOwnerPairWithType `json:"name_conflicts_and_owners,omitempty"`
	// LimitViolations - List of object types whose total number limit would be violated.
	LimitViolations []*NsObjectLimit `json:"limit_violations,omitempty"`
	// SnapRetainLimitViolations - List of snapshot retainment param limit violations.
	SnapRetainLimitViolations []*NsSnapRetainLimit `json:"snap_retain_limit_violations,omitempty"`
	// NetworkErrorList - List of validation errors with details.
	NetworkErrorList []*NsErrorWithArguments `json:"network_error_list,omitempty"`
	// AutoSwitchoverConflicts - List of validation errors for automatic switchover of Group Management.
	AutoSwitchoverConflicts []*NsErrorWithArguments `json:"auto_switchover_conflicts,omitempty"`
	// ErrorList - List of validation errors.
	ErrorList []*NsErrorWithArguments `json:"error_list,omitempty"`
	// AliasConflicts - List of WWPN/alias conflicts (same initiator WWPN, but different aliases).
	AliasConflicts []*NsAliasConflictPair `json:"alias_conflicts,omitempty"`
	// OnlineFcIntfs - List of Fibre Channel interfaces which are online.
	OnlineFcIntfs []*NsFibreChannelInterfaceFullName `json:"online_fc_intfs,omitempty"`
	// LunConflicts - List of LUN conflicts.
	LunConflicts []*NsLunConflictPair `json:"lun_conflicts,omitempty"`
	// SrcSid - Session ID for source group.
	SrcSid *string `json:"src_sid,omitempty"`
	// SrcGroupName - Name of the group we are merging with.
	SrcGroupName *string `json:"src_group_name,omitempty"`
	// DstGroupName - Name of the destination group.
	DstGroupName *string `json:"dst_group_name,omitempty"`
	// DstGroupSwversion - Software version of the destination group.
	DstGroupSwversion *string `json:"dst_group_swversion,omitempty"`
	// HtypeNameConflicts - List host type with different bit mask.
	HtypeNameConflicts []*string `json:"htype_name_conflicts,omitempty"`
	// HostTypeConflicts - List of initiators and igroups with different host types.
	HostTypeConflicts []*NsHostType `json:"host_type_conflicts,omitempty"`
	// ValidationError - Umbrella error code for group merge validation.
	ValidationError []*NsErrorWithArguments `json:"validation_error,omitempty"`
	// ValidationErrorMsg - Detailed error message.
	ValidationErrorMsg *string `json:"validation_error_msg,omitempty"`
	// WarningList - List of warning messages.
	WarningList []*string `json:"warning_list,omitempty"`
}

// Struct for NsGroupMergeReturnFields
type NsGroupMergeReturnStringFields struct {
	OnlineVols                 *string
	OnlineSnaps                *string
	ActivePartners             *string
	DstNatPartners             *string
	SrcThrottles               *string
	DstThrottles               *string
	ReplObjs                   *string
	NameConflicts              *string
	NameConflictsManualResolve *string
	SerialConflicts            *string
	AppUuidConflicts           *string
	NameConflictsAndOwners     *string
	LimitViolations            *string
	SnapRetainLimitViolations  *string
	NetworkErrorList           *string
	AutoSwitchoverConflicts    *string
	ErrorList                  *string
	AliasConflicts             *string
	OnlineFcIntfs              *string
	LunConflicts               *string
	SrcSid                     *string
	SrcGroupName               *string
	DstGroupName               *string
	DstGroupSwversion          *string
	HtypeNameConflicts         *string
	HostTypeConflicts          *string
	ValidationError            *string
	ValidationErrorMsg         *string
	WarningList                *string
}
