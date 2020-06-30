// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// MasterKey Service - Manage the master key. Data encryption keys for volumes are encrypted by using a master key that must be initialized before encrypted volumes can be created. The master key in
// turn is protected by a passphrase that is set when the master key is created. The passphrase may have to be entered to enable the master key when it is not available, for
// example, after an array reboot.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// MasterKeyService type 
type MasterKeyService struct {
	objectSet *client.MasterKeyObjectSet
}

// NewMasterKeyService - method to initialize "MasterKeyService" 
func NewMasterKeyService(gs *NsGroupService) (*MasterKeyService) {
	objectSet := gs.client.GetMasterKeyObjectSet()
	return &MasterKeyService{objectSet: objectSet}
}

// GetMasterKeys - method returns a array of pointers of type "MasterKeys"
func (svc *MasterKeyService) GetMasterKeys(params *util.GetParams) ([]*model.MasterKey, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetMasterKeysWithFields - method returns a array of pointers of type "MasterKey" 
func (svc *MasterKeyService) GetMasterKeysWithFields(fields []string) ([]*model.MasterKey, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateMasterKey - method creates a "MasterKey"
func (svc *MasterKeyService) CreateMasterKey(obj *model.MasterKey) (*model.MasterKey, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditMasterKey - method modifies  the "MasterKey" 
func (svc *MasterKeyService) EditMasterKey(id string, obj *model.MasterKey) (*model.MasterKey, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyMasterKey - private method for more than one element check. 
func onlyMasterKey(objs []*model.MasterKey) (*model.MasterKey, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one MasterKey found with the given filter")
	}

	return objs[0], nil
}

 
// GetMasterKeysByID - method returns associative a array of pointers of type "MasterKey", filter by Id
func (svc *MasterKeyService) GetMasterKeysByID(pool *model.Pool, fields []string) (map[string]*model.MasterKey, error) {
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
	objs, err := svc.GetMasterKeys(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.MasterKey)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetMasterKeyById - method returns a pointer to "MasterKey"
func (svc *MasterKeyService) GetMasterKeyById(id string) (*model.MasterKey, error) {
	return svc.objectSet.GetObject(id)
}

// GetMasterKeysByName - method returns a associative array of pointers of type "MasterKey", filter by name 
func (svc *MasterKeyService) GetMasterKeysByName(pool *model.Pool, fields []string) (map[string]*model.MasterKey, error) {
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
	objs, err := svc.GetMasterKeys(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.MasterKey)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetMasterKeyByName - method returns a pointer "MasterKey" 
func (svc *MasterKeyService) GetMasterKeyByName(name string) (*model.MasterKey, error) {
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
	return onlyMasterKey(objs)
}	

