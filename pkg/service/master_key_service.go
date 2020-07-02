// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// MasterKey Service - Manage the master key. Data encryption keys for volumes are encrypted by using a master key that must be initialized before encrypted volumes can be created. The master key in
// turn is protected by a passphrase that is set when the master key is created. The passphrase may have to be entered to enable the master key when it is not available, for
// example, after an array reboot.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// CreateMasterKey - method creates a "MasterKey"
func (svc *MasterKeyService) CreateMasterKey(obj *model.MasterKey) (*model.MasterKey, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateMasterKey - method modifies  the "MasterKey" 
func (svc *MasterKeyService) UpdateMasterKey(id string, obj *model.MasterKey) (*model.MasterKey, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetMasterKeyById - method returns a pointer to "MasterKey"
func (svc *MasterKeyService) GetMasterKeyById(id string) (*model.MasterKey, error) {
	return svc.objectSet.GetObject(id)
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
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteMasterKey - deletes the "MasterKey"
func (svc *MasterKeyService) DeleteMasterKey(id string) error {
	return svc.objectSet.DeleteObject(id)
}
