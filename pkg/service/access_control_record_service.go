// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

import (
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
)

type AccessControlRecordService struct {
	objectSet *client.AccessControlRecordObjectSet
}

func NewAccessControlRecordService(gs *GroupService) (vs *AccessControlRecordService) {
	objectSet := gs.client.GetAccessControlRecordObjectSet()
	return &AccessControlRecordService{objectSet: objectSet}
}
