// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
)

// NsGroupService type
type NsGroupService struct {
	ip     string
	client *client.GroupMgmtClient

	// Declare all the supported services
	versionService                    *VersionService
	applicationCategoryService        *ApplicationCategoryService
	chapUserService                   *ChapUserService
	masterKeyService                  *MasterKeyService
	alarmService                      *AlarmService
	volumeService                     *VolumeService
	shelfService                      *ShelfService
	keyManagerService                 *KeyManagerService
	protectionTemplateService         *ProtectionTemplateService
	folderService                     *FolderService
	tokenService                      *TokenService
	fibreChannelInterfaceService      *FibreChannelInterfaceService
	networkInterfaceService           *NetworkInterfaceService
	arrayService                      *ArrayService
	fibreChannelConfigService         *FibreChannelConfigService
	initiatorService                  *InitiatorService
	performancePolicyService          *PerformancePolicyService
	spaceDomainService                *SpaceDomainService
	snapshotCollectionService         *SnapshotCollectionService
	replicationPartnerService         *ReplicationPartnerService
	eventService                      *EventService
	snapshotService                   *SnapshotService
	applicationServerService          *ApplicationServerService
	userPolicyService                 *UserPolicyService
	userGroupService                  *UserGroupService
	subnetService                     *SubnetService
	controllerService                 *ControllerService
	fibreChannelSessionService        *FibreChannelSessionService
	userService                       *UserService
	protectionScheduleService         *ProtectionScheduleService
	initiatorGroupService             *InitiatorGroupService
	accessControlRecordService        *AccessControlRecordService
	activeDirectoryMembershipService  *ActiveDirectoryMembershipService
	fibreChannelPortService           *FibreChannelPortService
	protocolEndpointService           *ProtocolEndpointService
	witnessService                    *WitnessService
	jobService                        *JobService
	auditLogService                   *AuditLogService
	poolService                       *PoolService
	volumeCollectionService           *VolumeCollectionService
	diskService                       *DiskService
	fibreChannelInitiatorAliasService *FibreChannelInitiatorAliasService
	groupService                      *GroupService
	softwareVersionService            *SoftwareVersionService
	networkConfigService              *NetworkConfigService
}

// NewNsGroupService - initializes NsGroupService
func NewNsGroupService(ip string, username string, password string) (gs *NsGroupService, err error) {
	client, err := client.NewClient(ip, username, password, "v1")
	if err != nil {
		return nil, err
	}
	return &NsGroupService{ip: ip, client: client}, nil
}

// SetDebug - enable debugging
func (gs *NsGroupService) SetDebug() {
	gs.client.EnableDebug()
}

// GetVersionService - returns service of a type *VersionService
func (gs *NsGroupService) GetVersionService() (vs *VersionService) {
	if gs.versionService == nil {
		gs.versionService = NewVersionService(gs)
	}
	return gs.versionService
}

// GetApplicationCategoryService - returns service of a type *ApplicationCategoryService
func (gs *NsGroupService) GetApplicationCategoryService() (vs *ApplicationCategoryService) {
	if gs.applicationCategoryService == nil {
		gs.applicationCategoryService = NewApplicationCategoryService(gs)
	}
	return gs.applicationCategoryService
}

// GetChapUserService - returns service of a type *ChapUserService
func (gs *NsGroupService) GetChapUserService() (vs *ChapUserService) {
	if gs.chapUserService == nil {
		gs.chapUserService = NewChapUserService(gs)
	}
	return gs.chapUserService
}

// GetMasterKeyService - returns service of a type *MasterKeyService
func (gs *NsGroupService) GetMasterKeyService() (vs *MasterKeyService) {
	if gs.masterKeyService == nil {
		gs.masterKeyService = NewMasterKeyService(gs)
	}
	return gs.masterKeyService
}

// GetAlarmService - returns service of a type *AlarmService
func (gs *NsGroupService) GetAlarmService() (vs *AlarmService) {
	if gs.alarmService == nil {
		gs.alarmService = NewAlarmService(gs)
	}
	return gs.alarmService
}

// GetVolumeService - returns service of a type *VolumeService
func (gs *NsGroupService) GetVolumeService() (vs *VolumeService) {
	if gs.volumeService == nil {
		gs.volumeService = NewVolumeService(gs)
	}
	return gs.volumeService
}

// GetShelfService - returns service of a type *ShelfService
func (gs *NsGroupService) GetShelfService() (vs *ShelfService) {
	if gs.shelfService == nil {
		gs.shelfService = NewShelfService(gs)
	}
	return gs.shelfService
}

// GetKeyManagerService - returns service of a type *KeyManagerService
func (gs *NsGroupService) GetKeyManagerService() (vs *KeyManagerService) {
	if gs.keyManagerService == nil {
		gs.keyManagerService = NewKeyManagerService(gs)
	}
	return gs.keyManagerService
}

// GetProtectionTemplateService - returns service of a type *ProtectionTemplateService
func (gs *NsGroupService) GetProtectionTemplateService() (vs *ProtectionTemplateService) {
	if gs.protectionTemplateService == nil {
		gs.protectionTemplateService = NewProtectionTemplateService(gs)
	}
	return gs.protectionTemplateService
}

// GetFolderService - returns service of a type *FolderService
func (gs *NsGroupService) GetFolderService() (vs *FolderService) {
	if gs.folderService == nil {
		gs.folderService = NewFolderService(gs)
	}
	return gs.folderService
}

// GetTokenService - returns service of a type *TokenService
func (gs *NsGroupService) GetTokenService() (vs *TokenService) {
	if gs.tokenService == nil {
		gs.tokenService = NewTokenService(gs)
	}
	return gs.tokenService
}

// GetFibreChannelInterfaceService - returns service of a type *FibreChannelInterfaceService
func (gs *NsGroupService) GetFibreChannelInterfaceService() (vs *FibreChannelInterfaceService) {
	if gs.fibreChannelInterfaceService == nil {
		gs.fibreChannelInterfaceService = NewFibreChannelInterfaceService(gs)
	}
	return gs.fibreChannelInterfaceService
}

// GetNetworkInterfaceService - returns service of a type *NetworkInterfaceService
func (gs *NsGroupService) GetNetworkInterfaceService() (vs *NetworkInterfaceService) {
	if gs.networkInterfaceService == nil {
		gs.networkInterfaceService = NewNetworkInterfaceService(gs)
	}
	return gs.networkInterfaceService
}

// GetArrayService - returns service of a type *ArrayService
func (gs *NsGroupService) GetArrayService() (vs *ArrayService) {
	if gs.arrayService == nil {
		gs.arrayService = NewArrayService(gs)
	}
	return gs.arrayService
}

// GetFibreChannelConfigService - returns service of a type *FibreChannelConfigService
func (gs *NsGroupService) GetFibreChannelConfigService() (vs *FibreChannelConfigService) {
	if gs.fibreChannelConfigService == nil {
		gs.fibreChannelConfigService = NewFibreChannelConfigService(gs)
	}
	return gs.fibreChannelConfigService
}

// GetInitiatorService - returns service of a type *InitiatorService
func (gs *NsGroupService) GetInitiatorService() (vs *InitiatorService) {
	if gs.initiatorService == nil {
		gs.initiatorService = NewInitiatorService(gs)
	}
	return gs.initiatorService
}

// GetPerformancePolicyService - returns service of a type *PerformancePolicyService
func (gs *NsGroupService) GetPerformancePolicyService() (vs *PerformancePolicyService) {
	if gs.performancePolicyService == nil {
		gs.performancePolicyService = NewPerformancePolicyService(gs)
	}
	return gs.performancePolicyService
}

// GetSpaceDomainService - returns service of a type *SpaceDomainService
func (gs *NsGroupService) GetSpaceDomainService() (vs *SpaceDomainService) {
	if gs.spaceDomainService == nil {
		gs.spaceDomainService = NewSpaceDomainService(gs)
	}
	return gs.spaceDomainService
}

// GetSnapshotCollectionService - returns service of a type *SnapshotCollectionService
func (gs *NsGroupService) GetSnapshotCollectionService() (vs *SnapshotCollectionService) {
	if gs.snapshotCollectionService == nil {
		gs.snapshotCollectionService = NewSnapshotCollectionService(gs)
	}
	return gs.snapshotCollectionService
}

// GetReplicationPartnerService - returns service of a type *ReplicationPartnerService
func (gs *NsGroupService) GetReplicationPartnerService() (vs *ReplicationPartnerService) {
	if gs.replicationPartnerService == nil {
		gs.replicationPartnerService = NewReplicationPartnerService(gs)
	}
	return gs.replicationPartnerService
}

// GetEventService - returns service of a type *EventService
func (gs *NsGroupService) GetEventService() (vs *EventService) {
	if gs.eventService == nil {
		gs.eventService = NewEventService(gs)
	}
	return gs.eventService
}

// GetSnapshotService - returns service of a type *SnapshotService
func (gs *NsGroupService) GetSnapshotService() (vs *SnapshotService) {
	if gs.snapshotService == nil {
		gs.snapshotService = NewSnapshotService(gs)
	}
	return gs.snapshotService
}

// GetApplicationServerService - returns service of a type *ApplicationServerService
func (gs *NsGroupService) GetApplicationServerService() (vs *ApplicationServerService) {
	if gs.applicationServerService == nil {
		gs.applicationServerService = NewApplicationServerService(gs)
	}
	return gs.applicationServerService
}

// GetUserPolicyService - returns service of a type *UserPolicyService
func (gs *NsGroupService) GetUserPolicyService() (vs *UserPolicyService) {
	if gs.userPolicyService == nil {
		gs.userPolicyService = NewUserPolicyService(gs)
	}
	return gs.userPolicyService
}

// GetUserGroupService - returns service of a type *UserGroupService
func (gs *NsGroupService) GetUserGroupService() (vs *UserGroupService) {
	if gs.userGroupService == nil {
		gs.userGroupService = NewUserGroupService(gs)
	}
	return gs.userGroupService
}

// GetSubnetService - returns service of a type *SubnetService
func (gs *NsGroupService) GetSubnetService() (vs *SubnetService) {
	if gs.subnetService == nil {
		gs.subnetService = NewSubnetService(gs)
	}
	return gs.subnetService
}

// GetControllerService - returns service of a type *ControllerService
func (gs *NsGroupService) GetControllerService() (vs *ControllerService) {
	if gs.controllerService == nil {
		gs.controllerService = NewControllerService(gs)
	}
	return gs.controllerService
}

// GetFibreChannelSessionService - returns service of a type *FibreChannelSessionService
func (gs *NsGroupService) GetFibreChannelSessionService() (vs *FibreChannelSessionService) {
	if gs.fibreChannelSessionService == nil {
		gs.fibreChannelSessionService = NewFibreChannelSessionService(gs)
	}
	return gs.fibreChannelSessionService
}

// GetUserService - returns service of a type *UserService
func (gs *NsGroupService) GetUserService() (vs *UserService) {
	if gs.userService == nil {
		gs.userService = NewUserService(gs)
	}
	return gs.userService
}

// GetProtectionScheduleService - returns service of a type *ProtectionScheduleService
func (gs *NsGroupService) GetProtectionScheduleService() (vs *ProtectionScheduleService) {
	if gs.protectionScheduleService == nil {
		gs.protectionScheduleService = NewProtectionScheduleService(gs)
	}
	return gs.protectionScheduleService
}

// GetInitiatorGroupService - returns service of a type *InitiatorGroupService
func (gs *NsGroupService) GetInitiatorGroupService() (vs *InitiatorGroupService) {
	if gs.initiatorGroupService == nil {
		gs.initiatorGroupService = NewInitiatorGroupService(gs)
	}
	return gs.initiatorGroupService
}

// GetAccessControlRecordService - returns service of a type *AccessControlRecordService
func (gs *NsGroupService) GetAccessControlRecordService() (vs *AccessControlRecordService) {
	if gs.accessControlRecordService == nil {
		gs.accessControlRecordService = NewAccessControlRecordService(gs)
	}
	return gs.accessControlRecordService
}

// GetActiveDirectoryMembershipService - returns service of a type *ActiveDirectoryMembershipService
func (gs *NsGroupService) GetActiveDirectoryMembershipService() (vs *ActiveDirectoryMembershipService) {
	if gs.activeDirectoryMembershipService == nil {
		gs.activeDirectoryMembershipService = NewActiveDirectoryMembershipService(gs)
	}
	return gs.activeDirectoryMembershipService
}

// GetFibreChannelPortService - returns service of a type *FibreChannelPortService
func (gs *NsGroupService) GetFibreChannelPortService() (vs *FibreChannelPortService) {
	if gs.fibreChannelPortService == nil {
		gs.fibreChannelPortService = NewFibreChannelPortService(gs)
	}
	return gs.fibreChannelPortService
}

// GetProtocolEndpointService - returns service of a type *ProtocolEndpointService
func (gs *NsGroupService) GetProtocolEndpointService() (vs *ProtocolEndpointService) {
	if gs.protocolEndpointService == nil {
		gs.protocolEndpointService = NewProtocolEndpointService(gs)
	}
	return gs.protocolEndpointService
}

// GetWitnessService - returns service of a type *WitnessService
func (gs *NsGroupService) GetWitnessService() (vs *WitnessService) {
	if gs.witnessService == nil {
		gs.witnessService = NewWitnessService(gs)
	}
	return gs.witnessService
}

// GetJobService - returns service of a type *JobService
func (gs *NsGroupService) GetJobService() (vs *JobService) {
	if gs.jobService == nil {
		gs.jobService = NewJobService(gs)
	}
	return gs.jobService
}

// GetAuditLogService - returns service of a type *AuditLogService
func (gs *NsGroupService) GetAuditLogService() (vs *AuditLogService) {
	if gs.auditLogService == nil {
		gs.auditLogService = NewAuditLogService(gs)
	}
	return gs.auditLogService
}

// GetPoolService - returns service of a type *PoolService
func (gs *NsGroupService) GetPoolService() (vs *PoolService) {
	if gs.poolService == nil {
		gs.poolService = NewPoolService(gs)
	}
	return gs.poolService
}

// GetVolumeCollectionService - returns service of a type *VolumeCollectionService
func (gs *NsGroupService) GetVolumeCollectionService() (vs *VolumeCollectionService) {
	if gs.volumeCollectionService == nil {
		gs.volumeCollectionService = NewVolumeCollectionService(gs)
	}
	return gs.volumeCollectionService
}

// GetDiskService - returns service of a type *DiskService
func (gs *NsGroupService) GetDiskService() (vs *DiskService) {
	if gs.diskService == nil {
		gs.diskService = NewDiskService(gs)
	}
	return gs.diskService
}

// GetFibreChannelInitiatorAliasService - returns service of a type *FibreChannelInitiatorAliasService
func (gs *NsGroupService) GetFibreChannelInitiatorAliasService() (vs *FibreChannelInitiatorAliasService) {
	if gs.fibreChannelInitiatorAliasService == nil {
		gs.fibreChannelInitiatorAliasService = NewFibreChannelInitiatorAliasService(gs)
	}
	return gs.fibreChannelInitiatorAliasService
}

// GetGroupService - returns service of a type *GroupService
func (gs *NsGroupService) GetGroupService() (vs *GroupService) {
	if gs.groupService == nil {
		gs.groupService = NewGroupService(gs)
	}
	return gs.groupService
}

// GetSoftwareVersionService - returns service of a type *SoftwareVersionService
func (gs *NsGroupService) GetSoftwareVersionService() (vs *SoftwareVersionService) {
	if gs.softwareVersionService == nil {
		gs.softwareVersionService = NewSoftwareVersionService(gs)
	}
	return gs.softwareVersionService
}

// GetNetworkConfigService - returns service of a type *NetworkConfigService
func (gs *NsGroupService) GetNetworkConfigService() (vs *NetworkConfigService) {
	if gs.networkConfigService == nil {
		gs.networkConfigService = NewNetworkConfigService(gs)
	}
	return gs.networkConfigService
}
