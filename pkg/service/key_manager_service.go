// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// KeyManager Service - Key Manager stores encryption keys for the array volumes / dedupe domains.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// KeyManagerService type 
type KeyManagerService struct {
	objectSet *client.KeyManagerObjectSet
}

// NewKeyManagerService - method to initialize "KeyManagerService" 
func NewKeyManagerService(gs *NsGroupService) (*KeyManagerService) {
	objectSet := gs.client.GetKeyManagerObjectSet()
	return &KeyManagerService{objectSet: objectSet}
}

// GetKeyManagers - method returns a array of pointers of type "KeyManagers"
func (svc *KeyManagerService) GetKeyManagers(params *util.GetParams) ([]*model.KeyManager, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetKeyManagersWithFields - method returns a array of pointers of type "KeyManager" 
func (svc *KeyManagerService) GetKeyManagersWithFields(fields []string) ([]*model.KeyManager, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateKeyManager - method creates a "KeyManager"
func (svc *KeyManagerService) CreateKeyManager(obj *model.KeyManager) (*model.KeyManager, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditKeyManager - method modifies  the "KeyManager" 
func (svc *KeyManagerService) EditKeyManager(id string, obj *model.KeyManager) (*model.KeyManager, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyKeyManager - private method for more than one element check. 
func onlyKeyManager(objs []*model.KeyManager) (*model.KeyManager, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one KeyManager found with the given filter")
	}

	return objs[0], nil
}

 
// GetKeyManagersByID - method returns associative a array of pointers of type "KeyManager", filter by Id
func (svc *KeyManagerService) GetKeyManagersByID(pool *model.Pool, fields []string) (map[string]*model.KeyManager, error) {
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
	objs, err := svc.GetKeyManagers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.KeyManager)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetKeyManagerById - method returns a pointer to "KeyManager"
func (svc *KeyManagerService) GetKeyManagerById(id string) (*model.KeyManager, error) {
	return svc.objectSet.GetObject(id)
}

// GetKeyManagersByName - method returns a associative array of pointers of type "KeyManager", filter by name 
func (svc *KeyManagerService) GetKeyManagersByName(pool *model.Pool, fields []string) (map[string]*model.KeyManager, error) {
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
	objs, err := svc.GetKeyManagers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.KeyManager)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetKeyManagerByName - method returns a pointer "KeyManager" 
func (svc *KeyManagerService) GetKeyManagerByName(name string) (*model.KeyManager, error) {
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
	return onlyKeyManager(objs)
}	

