// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// KeyManager Service - Key Manager stores encryption keys for the array volumes / dedupe domains.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateKeyManager - method creates a "KeyManager"
func (svc *KeyManagerService) CreateKeyManager(obj *model.KeyManager) (*model.KeyManager, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateKeyManager - method modifies  the "KeyManager" 
func (svc *KeyManagerService) UpdateKeyManager(id string, obj *model.KeyManager) (*model.KeyManager, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetKeyManagerById - method returns a pointer to "KeyManager"
func (svc *KeyManagerService) GetKeyManagerById(id string) (*model.KeyManager, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteKeyManager - deletes the "KeyManager"
func (svc *KeyManagerService) DeleteKeyManager(id string) error {
	return svc.objectSet.DeleteObject(id)
}
