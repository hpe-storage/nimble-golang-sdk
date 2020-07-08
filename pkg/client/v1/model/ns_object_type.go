// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsObjectType Enum.

type NsObjectType string

const (
	cNsObjectTypeArrayNetconfig        NsObjectType = "array_netconfig"
	cNsObjectTypeUserPolicy            NsObjectType = "user_policy"
	cNsObjectTypeSubnet                NsObjectType = "subnet"
	cNsObjectTypeEncryptKey            NsObjectType = "encrypt_key"
	cNsObjectTypeInitiator             NsObjectType = "initiator"
	cNsObjectTypeKeymanager            NsObjectType = "keymanager"
	cNsObjectTypeNic                   NsObjectType = "nic"
	cNsObjectTypeBranch                NsObjectType = "branch"
	cNsObjectTypeFcTargetPortGroup     NsObjectType = "fc_target_port_group"
	cNsObjectTypeProttmpl              NsObjectType = "prottmpl"
	cNsObjectTypeProtpol               NsObjectType = "protpol"
	cNsObjectTypeSshkey                NsObjectType = "sshkey"
	cNsObjectTypeFcInterfaceCollection NsObjectType = "fc_interface_collection"
	cNsObjectTypeVolcoll               NsObjectType = "volcoll"
	cNsObjectTypeInitiatorgrpSubnet    NsObjectType = "initiatorgrp_subnet"
	cNsObjectTypePeAcl                 NsObjectType = "pe_acl"
	cNsObjectTypeVvolAcl               NsObjectType = "vvol_acl"
	cNsObjectTypeChapuser              NsObjectType = "chapuser"
	cNsObjectTypeEvents                NsObjectType = "events"
	cNsObjectTypeApplicationServer     NsObjectType = "application_server"
	cNsObjectTypeGroup                 NsObjectType = "group"
	cNsObjectTypePool                  NsObjectType = "pool"
	cNsObjectTypeVvol                  NsObjectType = "vvol"
	cNsObjectTypeActiveDirectory       NsObjectType = "active_directory"
	cNsObjectTypeShelf                 NsObjectType = "shelf"
	cNsObjectTypeDisk                  NsObjectType = "disk"
	cNsObjectTypeRoute                 NsObjectType = "route"
	cNsObjectTypeFolder                NsObjectType = "folder"
	cNsObjectTypeIpAddress             NsObjectType = "ip address"
	cNsObjectTypeFc                    NsObjectType = "fc"
	cNsObjectTypeSupport               NsObjectType = "support"
	cNsObjectTypeSnapshot              NsObjectType = "snapshot"
	cNsObjectTypeThrottle              NsObjectType = "throttle"
	cNsObjectTypeRole                  NsObjectType = "role"
	cNsObjectTypeSnapcoll              NsObjectType = "snapcoll"
	cNsObjectTypeSession               NsObjectType = "session"
	cNsObjectTypeAsyncJob              NsObjectType = "async_job"
	cNsObjectTypeInitiatorgrp          NsObjectType = "initiatorgrp"
	cNsObjectTypePerfpolicy            NsObjectType = "perfpolicy"
	cNsObjectTypePrivilege             NsObjectType = "privilege"
	cNsObjectTypeSyslog                NsObjectType = "syslog"
	cNsObjectTypeUserGroup             NsObjectType = "user group"
	cNsObjectTypeProtsched             NsObjectType = "protsched"
	cNsObjectTypeNetconfig             NsObjectType = "netconfig"
	cNsObjectTypeVol                   NsObjectType = "vol"
	cNsObjectTypeFcInitiatorAlias      NsObjectType = "fc_initiator_alias"
	cNsObjectTypeArray                 NsObjectType = "array"
	cNsObjectTypeAlarm                 NsObjectType = "alarm"
	cNsObjectTypeFcPort                NsObjectType = "fc_port"
	cNsObjectTypeProtocolEndpoint      NsObjectType = "protocol_endpoint"
	cNsObjectTypeFolset                NsObjectType = "folset"
	cNsObjectTypeAuditLog              NsObjectType = "audit_log"
	cNsObjectTypeHcClusterConfig       NsObjectType = "hc_cluster_config"
	cNsObjectTypeEncryptConfig         NsObjectType = "encrypt_config"
	cNsObjectTypeWitness               NsObjectType = "witness"
	cNsObjectTypePartner               NsObjectType = "partner"
	cNsObjectTypeSnapshotLun           NsObjectType = "snapshot_lun"
	cNsObjectTypeEventDipatcher        NsObjectType = "event_dipatcher"
	cNsObjectTypeVolacl                NsObjectType = "volacl"
	cNsObjectTypeUser                  NsObjectType = "user"
)

var pNsObjectTypeArrayNetconfig NsObjectType
var pNsObjectTypeUserPolicy NsObjectType
var pNsObjectTypeSubnet NsObjectType
var pNsObjectTypeEncryptKey NsObjectType
var pNsObjectTypeInitiator NsObjectType
var pNsObjectTypeKeymanager NsObjectType
var pNsObjectTypeNic NsObjectType
var pNsObjectTypeBranch NsObjectType
var pNsObjectTypeFcTargetPortGroup NsObjectType
var pNsObjectTypeProttmpl NsObjectType
var pNsObjectTypeProtpol NsObjectType
var pNsObjectTypeSshkey NsObjectType
var pNsObjectTypeFcInterfaceCollection NsObjectType
var pNsObjectTypeVolcoll NsObjectType
var pNsObjectTypeInitiatorgrpSubnet NsObjectType
var pNsObjectTypePeAcl NsObjectType
var pNsObjectTypeVvolAcl NsObjectType
var pNsObjectTypeChapuser NsObjectType
var pNsObjectTypeEvents NsObjectType
var pNsObjectTypeApplicationServer NsObjectType
var pNsObjectTypeGroup NsObjectType
var pNsObjectTypePool NsObjectType
var pNsObjectTypeVvol NsObjectType
var pNsObjectTypeActiveDirectory NsObjectType
var pNsObjectTypeShelf NsObjectType
var pNsObjectTypeDisk NsObjectType
var pNsObjectTypeRoute NsObjectType
var pNsObjectTypeFolder NsObjectType
var pNsObjectTypeIpAddress NsObjectType
var pNsObjectTypeFc NsObjectType
var pNsObjectTypeSupport NsObjectType
var pNsObjectTypeSnapshot NsObjectType
var pNsObjectTypeThrottle NsObjectType
var pNsObjectTypeRole NsObjectType
var pNsObjectTypeSnapcoll NsObjectType
var pNsObjectTypeSession NsObjectType
var pNsObjectTypeAsyncJob NsObjectType
var pNsObjectTypeInitiatorgrp NsObjectType
var pNsObjectTypePerfpolicy NsObjectType
var pNsObjectTypePrivilege NsObjectType
var pNsObjectTypeSyslog NsObjectType
var pNsObjectTypeUserGroup NsObjectType
var pNsObjectTypeProtsched NsObjectType
var pNsObjectTypeNetconfig NsObjectType
var pNsObjectTypeVol NsObjectType
var pNsObjectTypeFcInitiatorAlias NsObjectType
var pNsObjectTypeArray NsObjectType
var pNsObjectTypeAlarm NsObjectType
var pNsObjectTypeFcPort NsObjectType
var pNsObjectTypeProtocolEndpoint NsObjectType
var pNsObjectTypeFolset NsObjectType
var pNsObjectTypeAuditLog NsObjectType
var pNsObjectTypeHcClusterConfig NsObjectType
var pNsObjectTypeEncryptConfig NsObjectType
var pNsObjectTypeWitness NsObjectType
var pNsObjectTypePartner NsObjectType
var pNsObjectTypeSnapshotLun NsObjectType
var pNsObjectTypeEventDipatcher NsObjectType
var pNsObjectTypeVolacl NsObjectType
var pNsObjectTypeUser NsObjectType

// Export
var NsObjectTypeArrayNetconfig *NsObjectType
var NsObjectTypeUserPolicy *NsObjectType
var NsObjectTypeSubnet *NsObjectType
var NsObjectTypeEncryptKey *NsObjectType
var NsObjectTypeInitiator *NsObjectType
var NsObjectTypeKeymanager *NsObjectType
var NsObjectTypeNic *NsObjectType
var NsObjectTypeBranch *NsObjectType
var NsObjectTypeFcTargetPortGroup *NsObjectType
var NsObjectTypeProttmpl *NsObjectType
var NsObjectTypeProtpol *NsObjectType
var NsObjectTypeSshkey *NsObjectType
var NsObjectTypeFcInterfaceCollection *NsObjectType
var NsObjectTypeVolcoll *NsObjectType
var NsObjectTypeInitiatorgrpSubnet *NsObjectType
var NsObjectTypePeAcl *NsObjectType
var NsObjectTypeVvolAcl *NsObjectType
var NsObjectTypeChapuser *NsObjectType
var NsObjectTypeEvents *NsObjectType
var NsObjectTypeApplicationServer *NsObjectType
var NsObjectTypeGroup *NsObjectType
var NsObjectTypePool *NsObjectType
var NsObjectTypeVvol *NsObjectType
var NsObjectTypeActiveDirectory *NsObjectType
var NsObjectTypeShelf *NsObjectType
var NsObjectTypeDisk *NsObjectType
var NsObjectTypeRoute *NsObjectType
var NsObjectTypeFolder *NsObjectType
var NsObjectTypeIpAddress *NsObjectType
var NsObjectTypeFc *NsObjectType
var NsObjectTypeSupport *NsObjectType
var NsObjectTypeSnapshot *NsObjectType
var NsObjectTypeThrottle *NsObjectType
var NsObjectTypeRole *NsObjectType
var NsObjectTypeSnapcoll *NsObjectType
var NsObjectTypeSession *NsObjectType
var NsObjectTypeAsyncJob *NsObjectType
var NsObjectTypeInitiatorgrp *NsObjectType
var NsObjectTypePerfpolicy *NsObjectType
var NsObjectTypePrivilege *NsObjectType
var NsObjectTypeSyslog *NsObjectType
var NsObjectTypeUserGroup *NsObjectType
var NsObjectTypeProtsched *NsObjectType
var NsObjectTypeNetconfig *NsObjectType
var NsObjectTypeVol *NsObjectType
var NsObjectTypeFcInitiatorAlias *NsObjectType
var NsObjectTypeArray *NsObjectType
var NsObjectTypeAlarm *NsObjectType
var NsObjectTypeFcPort *NsObjectType
var NsObjectTypeProtocolEndpoint *NsObjectType
var NsObjectTypeFolset *NsObjectType
var NsObjectTypeAuditLog *NsObjectType
var NsObjectTypeHcClusterConfig *NsObjectType
var NsObjectTypeEncryptConfig *NsObjectType
var NsObjectTypeWitness *NsObjectType
var NsObjectTypePartner *NsObjectType
var NsObjectTypeSnapshotLun *NsObjectType
var NsObjectTypeEventDipatcher *NsObjectType
var NsObjectTypeVolacl *NsObjectType
var NsObjectTypeUser *NsObjectType

func init() {
	pNsObjectTypeArrayNetconfig = cNsObjectTypeArrayNetconfig
	NsObjectTypeArrayNetconfig = &pNsObjectTypeArrayNetconfig

	pNsObjectTypeUserPolicy = cNsObjectTypeUserPolicy
	NsObjectTypeUserPolicy = &pNsObjectTypeUserPolicy

	pNsObjectTypeSubnet = cNsObjectTypeSubnet
	NsObjectTypeSubnet = &pNsObjectTypeSubnet

	pNsObjectTypeEncryptKey = cNsObjectTypeEncryptKey
	NsObjectTypeEncryptKey = &pNsObjectTypeEncryptKey

	pNsObjectTypeInitiator = cNsObjectTypeInitiator
	NsObjectTypeInitiator = &pNsObjectTypeInitiator

	pNsObjectTypeKeymanager = cNsObjectTypeKeymanager
	NsObjectTypeKeymanager = &pNsObjectTypeKeymanager

	pNsObjectTypeNic = cNsObjectTypeNic
	NsObjectTypeNic = &pNsObjectTypeNic

	pNsObjectTypeBranch = cNsObjectTypeBranch
	NsObjectTypeBranch = &pNsObjectTypeBranch

	pNsObjectTypeFcTargetPortGroup = cNsObjectTypeFcTargetPortGroup
	NsObjectTypeFcTargetPortGroup = &pNsObjectTypeFcTargetPortGroup

	pNsObjectTypeProttmpl = cNsObjectTypeProttmpl
	NsObjectTypeProttmpl = &pNsObjectTypeProttmpl

	pNsObjectTypeProtpol = cNsObjectTypeProtpol
	NsObjectTypeProtpol = &pNsObjectTypeProtpol

	pNsObjectTypeSshkey = cNsObjectTypeSshkey
	NsObjectTypeSshkey = &pNsObjectTypeSshkey

	pNsObjectTypeFcInterfaceCollection = cNsObjectTypeFcInterfaceCollection
	NsObjectTypeFcInterfaceCollection = &pNsObjectTypeFcInterfaceCollection

	pNsObjectTypeVolcoll = cNsObjectTypeVolcoll
	NsObjectTypeVolcoll = &pNsObjectTypeVolcoll

	pNsObjectTypeInitiatorgrpSubnet = cNsObjectTypeInitiatorgrpSubnet
	NsObjectTypeInitiatorgrpSubnet = &pNsObjectTypeInitiatorgrpSubnet

	pNsObjectTypePeAcl = cNsObjectTypePeAcl
	NsObjectTypePeAcl = &pNsObjectTypePeAcl

	pNsObjectTypeVvolAcl = cNsObjectTypeVvolAcl
	NsObjectTypeVvolAcl = &pNsObjectTypeVvolAcl

	pNsObjectTypeChapuser = cNsObjectTypeChapuser
	NsObjectTypeChapuser = &pNsObjectTypeChapuser

	pNsObjectTypeEvents = cNsObjectTypeEvents
	NsObjectTypeEvents = &pNsObjectTypeEvents

	pNsObjectTypeApplicationServer = cNsObjectTypeApplicationServer
	NsObjectTypeApplicationServer = &pNsObjectTypeApplicationServer

	pNsObjectTypeGroup = cNsObjectTypeGroup
	NsObjectTypeGroup = &pNsObjectTypeGroup

	pNsObjectTypePool = cNsObjectTypePool
	NsObjectTypePool = &pNsObjectTypePool

	pNsObjectTypeVvol = cNsObjectTypeVvol
	NsObjectTypeVvol = &pNsObjectTypeVvol

	pNsObjectTypeActiveDirectory = cNsObjectTypeActiveDirectory
	NsObjectTypeActiveDirectory = &pNsObjectTypeActiveDirectory

	pNsObjectTypeShelf = cNsObjectTypeShelf
	NsObjectTypeShelf = &pNsObjectTypeShelf

	pNsObjectTypeDisk = cNsObjectTypeDisk
	NsObjectTypeDisk = &pNsObjectTypeDisk

	pNsObjectTypeRoute = cNsObjectTypeRoute
	NsObjectTypeRoute = &pNsObjectTypeRoute

	pNsObjectTypeFolder = cNsObjectTypeFolder
	NsObjectTypeFolder = &pNsObjectTypeFolder

	pNsObjectTypeIpAddress = cNsObjectTypeIpAddress
	NsObjectTypeIpAddress = &pNsObjectTypeIpAddress

	pNsObjectTypeFc = cNsObjectTypeFc
	NsObjectTypeFc = &pNsObjectTypeFc

	pNsObjectTypeSupport = cNsObjectTypeSupport
	NsObjectTypeSupport = &pNsObjectTypeSupport

	pNsObjectTypeSnapshot = cNsObjectTypeSnapshot
	NsObjectTypeSnapshot = &pNsObjectTypeSnapshot

	pNsObjectTypeThrottle = cNsObjectTypeThrottle
	NsObjectTypeThrottle = &pNsObjectTypeThrottle

	pNsObjectTypeRole = cNsObjectTypeRole
	NsObjectTypeRole = &pNsObjectTypeRole

	pNsObjectTypeSnapcoll = cNsObjectTypeSnapcoll
	NsObjectTypeSnapcoll = &pNsObjectTypeSnapcoll

	pNsObjectTypeSession = cNsObjectTypeSession
	NsObjectTypeSession = &pNsObjectTypeSession

	pNsObjectTypeAsyncJob = cNsObjectTypeAsyncJob
	NsObjectTypeAsyncJob = &pNsObjectTypeAsyncJob

	pNsObjectTypeInitiatorgrp = cNsObjectTypeInitiatorgrp
	NsObjectTypeInitiatorgrp = &pNsObjectTypeInitiatorgrp

	pNsObjectTypePerfpolicy = cNsObjectTypePerfpolicy
	NsObjectTypePerfpolicy = &pNsObjectTypePerfpolicy

	pNsObjectTypePrivilege = cNsObjectTypePrivilege
	NsObjectTypePrivilege = &pNsObjectTypePrivilege

	pNsObjectTypeSyslog = cNsObjectTypeSyslog
	NsObjectTypeSyslog = &pNsObjectTypeSyslog

	pNsObjectTypeUserGroup = cNsObjectTypeUserGroup
	NsObjectTypeUserGroup = &pNsObjectTypeUserGroup

	pNsObjectTypeProtsched = cNsObjectTypeProtsched
	NsObjectTypeProtsched = &pNsObjectTypeProtsched

	pNsObjectTypeNetconfig = cNsObjectTypeNetconfig
	NsObjectTypeNetconfig = &pNsObjectTypeNetconfig

	pNsObjectTypeVol = cNsObjectTypeVol
	NsObjectTypeVol = &pNsObjectTypeVol

	pNsObjectTypeFcInitiatorAlias = cNsObjectTypeFcInitiatorAlias
	NsObjectTypeFcInitiatorAlias = &pNsObjectTypeFcInitiatorAlias

	pNsObjectTypeArray = cNsObjectTypeArray
	NsObjectTypeArray = &pNsObjectTypeArray

	pNsObjectTypeAlarm = cNsObjectTypeAlarm
	NsObjectTypeAlarm = &pNsObjectTypeAlarm

	pNsObjectTypeFcPort = cNsObjectTypeFcPort
	NsObjectTypeFcPort = &pNsObjectTypeFcPort

	pNsObjectTypeProtocolEndpoint = cNsObjectTypeProtocolEndpoint
	NsObjectTypeProtocolEndpoint = &pNsObjectTypeProtocolEndpoint

	pNsObjectTypeFolset = cNsObjectTypeFolset
	NsObjectTypeFolset = &pNsObjectTypeFolset

	pNsObjectTypeAuditLog = cNsObjectTypeAuditLog
	NsObjectTypeAuditLog = &pNsObjectTypeAuditLog

	pNsObjectTypeHcClusterConfig = cNsObjectTypeHcClusterConfig
	NsObjectTypeHcClusterConfig = &pNsObjectTypeHcClusterConfig

	pNsObjectTypeEncryptConfig = cNsObjectTypeEncryptConfig
	NsObjectTypeEncryptConfig = &pNsObjectTypeEncryptConfig

	pNsObjectTypeWitness = cNsObjectTypeWitness
	NsObjectTypeWitness = &pNsObjectTypeWitness

	pNsObjectTypePartner = cNsObjectTypePartner
	NsObjectTypePartner = &pNsObjectTypePartner

	pNsObjectTypeSnapshotLun = cNsObjectTypeSnapshotLun
	NsObjectTypeSnapshotLun = &pNsObjectTypeSnapshotLun

	pNsObjectTypeEventDipatcher = cNsObjectTypeEventDipatcher
	NsObjectTypeEventDipatcher = &pNsObjectTypeEventDipatcher

	pNsObjectTypeVolacl = cNsObjectTypeVolacl
	NsObjectTypeVolacl = &pNsObjectTypeVolacl

	pNsObjectTypeUser = cNsObjectTypeUser
	NsObjectTypeUser = &pNsObjectTypeUser

}
