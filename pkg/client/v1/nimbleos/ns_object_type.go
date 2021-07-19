// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

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
	cNsObjectTypeTrustedOauthIssuer    NsObjectType = "trusted_oauth_issuer"
	cNsObjectTypeAlarm                 NsObjectType = "alarm"
	cNsObjectTypeFcPort                NsObjectType = "fc_port"
	cNsObjectTypeProtocolEndpoint      NsObjectType = "protocol_endpoint"
	cNsObjectTypeFolset                NsObjectType = "folset"
	cNsObjectTypeAuditLog              NsObjectType = "audit_log"
	cNsObjectTypeHcClusterConfig       NsObjectType = "hc_cluster_config"
	cNsObjectTypeEncryptConfig         NsObjectType = "encrypt_config"
	cNsObjectTypeClientCredential      NsObjectType = "client_credential"
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
var pNsObjectTypeTrustedOauthIssuer NsObjectType
var pNsObjectTypeAlarm NsObjectType
var pNsObjectTypeFcPort NsObjectType
var pNsObjectTypeProtocolEndpoint NsObjectType
var pNsObjectTypeFolset NsObjectType
var pNsObjectTypeAuditLog NsObjectType
var pNsObjectTypeHcClusterConfig NsObjectType
var pNsObjectTypeEncryptConfig NsObjectType
var pNsObjectTypeClientCredential NsObjectType
var pNsObjectTypeWitness NsObjectType
var pNsObjectTypePartner NsObjectType
var pNsObjectTypeSnapshotLun NsObjectType
var pNsObjectTypeEventDipatcher NsObjectType
var pNsObjectTypeVolacl NsObjectType
var pNsObjectTypeUser NsObjectType

// NsObjectTypeArrayNetconfig enum exports
var NsObjectTypeArrayNetconfig *NsObjectType

// NsObjectTypeUserPolicy enum exports
var NsObjectTypeUserPolicy *NsObjectType

// NsObjectTypeSubnet enum exports
var NsObjectTypeSubnet *NsObjectType

// NsObjectTypeEncryptKey enum exports
var NsObjectTypeEncryptKey *NsObjectType

// NsObjectTypeInitiator enum exports
var NsObjectTypeInitiator *NsObjectType

// NsObjectTypeKeymanager enum exports
var NsObjectTypeKeymanager *NsObjectType

// NsObjectTypeNic enum exports
var NsObjectTypeNic *NsObjectType

// NsObjectTypeBranch enum exports
var NsObjectTypeBranch *NsObjectType

// NsObjectTypeFcTargetPortGroup enum exports
var NsObjectTypeFcTargetPortGroup *NsObjectType

// NsObjectTypeProttmpl enum exports
var NsObjectTypeProttmpl *NsObjectType

// NsObjectTypeProtpol enum exports
var NsObjectTypeProtpol *NsObjectType

// NsObjectTypeSshkey enum exports
var NsObjectTypeSshkey *NsObjectType

// NsObjectTypeFcInterfaceCollection enum exports
var NsObjectTypeFcInterfaceCollection *NsObjectType

// NsObjectTypeVolcoll enum exports
var NsObjectTypeVolcoll *NsObjectType

// NsObjectTypeInitiatorgrpSubnet enum exports
var NsObjectTypeInitiatorgrpSubnet *NsObjectType

// NsObjectTypePeAcl enum exports
var NsObjectTypePeAcl *NsObjectType

// NsObjectTypeVvolAcl enum exports
var NsObjectTypeVvolAcl *NsObjectType

// NsObjectTypeChapuser enum exports
var NsObjectTypeChapuser *NsObjectType

// NsObjectTypeEvents enum exports
var NsObjectTypeEvents *NsObjectType

// NsObjectTypeApplicationServer enum exports
var NsObjectTypeApplicationServer *NsObjectType

// NsObjectTypeGroup enum exports
var NsObjectTypeGroup *NsObjectType

// NsObjectTypePool enum exports
var NsObjectTypePool *NsObjectType

// NsObjectTypeVvol enum exports
var NsObjectTypeVvol *NsObjectType

// NsObjectTypeActiveDirectory enum exports
var NsObjectTypeActiveDirectory *NsObjectType

// NsObjectTypeShelf enum exports
var NsObjectTypeShelf *NsObjectType

// NsObjectTypeDisk enum exports
var NsObjectTypeDisk *NsObjectType

// NsObjectTypeRoute enum exports
var NsObjectTypeRoute *NsObjectType

// NsObjectTypeFolder enum exports
var NsObjectTypeFolder *NsObjectType

// NsObjectTypeIpAddress enum exports
var NsObjectTypeIpAddress *NsObjectType

// NsObjectTypeFc enum exports
var NsObjectTypeFc *NsObjectType

// NsObjectTypeSupport enum exports
var NsObjectTypeSupport *NsObjectType

// NsObjectTypeSnapshot enum exports
var NsObjectTypeSnapshot *NsObjectType

// NsObjectTypeThrottle enum exports
var NsObjectTypeThrottle *NsObjectType

// NsObjectTypeRole enum exports
var NsObjectTypeRole *NsObjectType

// NsObjectTypeSnapcoll enum exports
var NsObjectTypeSnapcoll *NsObjectType

// NsObjectTypeSession enum exports
var NsObjectTypeSession *NsObjectType

// NsObjectTypeAsyncJob enum exports
var NsObjectTypeAsyncJob *NsObjectType

// NsObjectTypeInitiatorgrp enum exports
var NsObjectTypeInitiatorgrp *NsObjectType

// NsObjectTypePerfpolicy enum exports
var NsObjectTypePerfpolicy *NsObjectType

// NsObjectTypePrivilege enum exports
var NsObjectTypePrivilege *NsObjectType

// NsObjectTypeSyslog enum exports
var NsObjectTypeSyslog *NsObjectType

// NsObjectTypeUserGroup enum exports
var NsObjectTypeUserGroup *NsObjectType

// NsObjectTypeProtsched enum exports
var NsObjectTypeProtsched *NsObjectType

// NsObjectTypeNetconfig enum exports
var NsObjectTypeNetconfig *NsObjectType

// NsObjectTypeVol enum exports
var NsObjectTypeVol *NsObjectType

// NsObjectTypeFcInitiatorAlias enum exports
var NsObjectTypeFcInitiatorAlias *NsObjectType

// NsObjectTypeArray enum exports
var NsObjectTypeArray *NsObjectType

// NsObjectTypeTrustedOauthIssuer enum exports
var NsObjectTypeTrustedOauthIssuer *NsObjectType

// NsObjectTypeAlarm enum exports
var NsObjectTypeAlarm *NsObjectType

// NsObjectTypeFcPort enum exports
var NsObjectTypeFcPort *NsObjectType

// NsObjectTypeProtocolEndpoint enum exports
var NsObjectTypeProtocolEndpoint *NsObjectType

// NsObjectTypeFolset enum exports
var NsObjectTypeFolset *NsObjectType

// NsObjectTypeAuditLog enum exports
var NsObjectTypeAuditLog *NsObjectType

// NsObjectTypeHcClusterConfig enum exports
var NsObjectTypeHcClusterConfig *NsObjectType

// NsObjectTypeEncryptConfig enum exports
var NsObjectTypeEncryptConfig *NsObjectType

// NsObjectTypeClientCredential enum exports
var NsObjectTypeClientCredential *NsObjectType

// NsObjectTypeWitness enum exports
var NsObjectTypeWitness *NsObjectType

// NsObjectTypePartner enum exports
var NsObjectTypePartner *NsObjectType

// NsObjectTypeSnapshotLun enum exports
var NsObjectTypeSnapshotLun *NsObjectType

// NsObjectTypeEventDipatcher enum exports
var NsObjectTypeEventDipatcher *NsObjectType

// NsObjectTypeVolacl enum exports
var NsObjectTypeVolacl *NsObjectType

// NsObjectTypeUser enum exports
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

	pNsObjectTypeTrustedOauthIssuer = cNsObjectTypeTrustedOauthIssuer
	NsObjectTypeTrustedOauthIssuer = &pNsObjectTypeTrustedOauthIssuer

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

	pNsObjectTypeClientCredential = cNsObjectTypeClientCredential
	NsObjectTypeClientCredential = &pNsObjectTypeClientCredential

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
