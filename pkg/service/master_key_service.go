// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// MasterKey Service - Manage the master key. Data encryption keys for volumes are encrypted by using a master key that must be initialized before encrypted volumes can be created. The master key in
// turn is protected by a passphrase that is set when the master key is created. The passphrase may have to be entered to enable the master key when it is not available, for
// example, after an array reboot.

import (
	"fmt"
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
	if params == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",params)
	}
	
	masterKeyResp,err := svc.objectSet.GetObjectListFromParams(params)
	if err !=nil {
		return nil,err
	}
	return masterKeyResp,nil
}

// CreateMasterKey - method creates a "MasterKey"
func (svc *MasterKeyService) CreateMasterKey(obj *model.MasterKey) (*model.MasterKey, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	masterKeyResp,err := svc.objectSet.CreateObject(obj)
	if err !=nil {
		return nil,err
	}
	return masterKeyResp,nil
}

// UpdateMasterKey - method modifies  the "MasterKey" 
func (svc *MasterKeyService) UpdateMasterKey(id string, obj *model.MasterKey) (*model.MasterKey, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	masterKeyResp,err :=svc.objectSet.UpdateObject(id, obj)
	if err !=nil {
		return nil,err
	}
	return masterKeyResp,nil
}

// GetMasterKeyById - method returns a pointer to "MasterKey"
func (svc *MasterKeyService) GetMasterKeyById(id string) (*model.MasterKey, error) {
	if len(id) == 0 {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",id)
	}
	
	masterKeyResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil,err
	}
	return masterKeyResp,nil
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
	masterKeyResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(masterKeyResp) == 0 {
    	return nil, nil
    }
    
	return masterKeyResp[0],nil
}	

// DeleteMasterKey - deletes the "MasterKey"
func (svc *MasterKeyService) DeleteMasterKey(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s",id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
