// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ChapUser Service - Manage Challenge-Response Handshake Authentication Protocol (CHAP) user accounts. CHAP users are one method of access control for iSCSI initiators. Each CHAP user has a CHAP
// password, sometimes called a CHAP secret. The CHAP passwords must match on the array and on the iSCSI initiator in order for the array to authenicate the initiator and allow it
// access. The CHAP user information must exist on both the array and the iSCSI initiator. Target authentication gives security only for the specific iSCSI target.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetChapUsersWithFields - method returns a array of pointers of type "ChapUser" 
func (svc *ChapUserService) GetChapUsersWithFields(fields []string) ([]*model.ChapUser, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateChapUser - method creates a "ChapUser"
func (svc *ChapUserService) CreateChapUser(obj *model.ChapUser) (*model.ChapUser, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditChapUser - method modifies  the "ChapUser" 
func (svc *ChapUserService) EditChapUser(id string, obj *model.ChapUser) (*model.ChapUser, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyChapUser - private method for more than one element check. 
func onlyChapUser(objs []*model.ChapUser) (*model.ChapUser, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one ChapUser found with the given filter")
	}

	return objs[0], nil
}

 
// GetChapUsersByID - method returns associative a array of pointers of type "ChapUser", filter by Id
func (svc *ChapUserService) GetChapUsersByID(pool *model.Pool, fields []string) (map[string]*model.ChapUser, error) {
	params := &util.GetParams{}

	// make sure ID field is selected
	if _, found := params.FindField("id"); !found {
		fields = append(fields, "id")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetChapUsers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ChapUser)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetChapUserById - method returns a pointer to "ChapUser"
func (svc *ChapUserService) GetChapUserById(id string) (*model.ChapUser, error) {
	return svc.objectSet.GetObject(id)
}

// GetChapUsersByName - method returns a associative array of pointers of type "ChapUser", filter by name 
func (svc *ChapUserService) GetChapUsersByName(pool *model.Pool, fields []string) (map[string]*model.ChapUser, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetChapUsers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ChapUser)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
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
	return onlyChapUser(objs)
}	

