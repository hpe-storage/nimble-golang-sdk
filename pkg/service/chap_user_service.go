// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ChapUser Service - Manage Challenge-Response Handshake Authentication Protocol (CHAP) user accounts. CHAP users are one method of access control for iSCSI initiators. Each CHAP user has a CHAP
// password, sometimes called a CHAP secret. The CHAP passwords must match on the array and on the iSCSI initiator in order for the array to authenicate the initiator and allow it
// access. The CHAP user information must exist on both the array and the iSCSI initiator. Target authentication gives security only for the specific iSCSI target.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// ChapUserService type 
type ChapUserService struct {
	objectSet *client.ChapUserObjectSet
}

// NewChapUserService - method to initialize "ChapUserService" 
func NewChapUserService(gs *NsGroupService) (*ChapUserService) {
	objectSet := gs.client.GetChapUserObjectSet()
	return &ChapUserService{objectSet: objectSet}
}

// GetChapUsers - method returns a array of pointers of type "ChapUsers"
func (svc *ChapUserService) GetChapUsers(params *util.GetParams) ([]*model.ChapUser, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateChapUser - method creates a "ChapUser"
func (svc *ChapUserService) CreateChapUser(obj *model.ChapUser) (*model.ChapUser, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateChapUser - method modifies  the "ChapUser" 
func (svc *ChapUserService) UpdateChapUser(id string, obj *model.ChapUser) (*model.ChapUser, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetChapUserById - method returns a pointer to "ChapUser"
func (svc *ChapUserService) GetChapUserById(id string) (*model.ChapUser, error) {
	return svc.objectSet.GetObject(id)
}

// GetChapUserByName - method returns a pointer "ChapUser" 
func (svc *ChapUserService) GetChapUserByName(name string) (*model.ChapUser, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteChapUser - deletes the "ChapUser"
func (svc *ChapUserService) DeleteChapUser(id string) error {
	return svc.objectSet.DeleteObject(id)
}
