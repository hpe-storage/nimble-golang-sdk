// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

import (
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
)

type GroupService struct {
	ip     string
	client *client.GroupMgmtClient

	accessControlRecordService *AccessControlRecordService
	performancePolicyService   *PerformancePolicyService
	tokenService               *TokenService
	volumeService              *VolumeService
}

func NewGroupService(ip string, username string, password string) (gs *GroupService, err error) {
	client, err := client.NewClient(ip, username, password, "v1")
	if err != nil {
		return nil, err
	}

	return &GroupService{ip: ip, client: client}, nil
}
func(gs *GroupService) EnableDebug(){
	gs.client.EnableDebug()
}
func (gs *GroupService) GetAccessControlRecordService() (vs *AccessControlRecordService) {
	if gs.accessControlRecordService == nil {
		gs.accessControlRecordService = NewAccessControlRecordService(gs)
	}
	return gs.accessControlRecordService
}

func (gs *GroupService) GetPerformancePolicyService() (vs *PerformancePolicyService) {
	if gs.performancePolicyService == nil {
		gs.performancePolicyService = NewPerformancePolicyService(gs)
	}
	return gs.performancePolicyService
}

func (gs *GroupService) GetTokenService() (vs *TokenService) {
	if gs.tokenService == nil {
		gs.tokenService = NewTokenService(gs)
	}
	return gs.tokenService
}

func (gs *GroupService) GetVolumeService() (vs *VolumeService) {
	if gs.volumeService == nil {
		gs.volumeService = NewVolumeService(gs)
	}
	return gs.volumeService
}
