// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// KeyManager Service - Key Manager stores encryption keys for the array volumes / dedupe domains.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// KeyManagerService type
type KeyManagerService struct {
	objectSet *client.KeyManagerObjectSet
}

// NewKeyManagerService - method to initialize "KeyManagerService"
func NewKeyManagerService(gs *NsGroupService) *KeyManagerService {
	objectSet := gs.client.GetKeyManagerObjectSet()
	return &KeyManagerService{objectSet: objectSet}
}

// GetKeyManagers - method returns a array of pointers of type "KeyManagers"
func (svc *KeyManagerService) GetKeyManagers(params *util.GetParams) ([]*model.KeyManager, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	keyManagerResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return keyManagerResp, nil
}

// CreateKeyManager - method creates a "KeyManager"
func (svc *KeyManagerService) CreateKeyManager(obj *model.KeyManager) (*model.KeyManager, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	keyManagerResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return keyManagerResp, nil
}

// UpdateKeyManager - method modifies  the "KeyManager"
func (svc *KeyManagerService) UpdateKeyManager(id string, obj *model.KeyManager) (*model.KeyManager, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	keyManagerResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return keyManagerResp, nil
}

// GetKeyManagerById - method returns a pointer to "KeyManager"
func (svc *KeyManagerService) GetKeyManagerById(id string) (*model.KeyManager, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	keyManagerResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return keyManagerResp, nil
}

// GetKeyManagerByName - method returns a pointer "KeyManager"
func (svc *KeyManagerService) GetKeyManagerByName(name string) (*model.KeyManager, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: &model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	keyManagerResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(keyManagerResp) == 0 {
		return nil, nil
	}

	return keyManagerResp[0], nil
}

// DeleteKeyManager - deletes the "KeyManager"
func (svc *KeyManagerService) DeleteKeyManager(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
