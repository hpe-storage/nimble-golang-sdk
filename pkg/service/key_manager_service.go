// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// KeyManager Service - Key Manager stores encryption keys for the array volumes / dedupe domains.

import (
    "fmt"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (svc *KeyManagerService) GetKeyManagers(params *param.GetParams) ([]*nimbleos.KeyManager, error) {
    keyManagerResp,err := svc.objectSet.GetObjectListFromParams(params)
    if err !=nil {
        return nil,err
    }
    return keyManagerResp,nil
}

// CreateKeyManager - method creates a "KeyManager"
func (svc *KeyManagerService) CreateKeyManager(obj *nimbleos.KeyManager) (*nimbleos.KeyManager, error) {
    if obj == nil {
        return nil,fmt.Errorf("CreateKeyManager: invalid parameter specified, %v",obj)
    }

    keyManagerResp,err := svc.objectSet.CreateObject(obj)
    if err !=nil {
        return nil,err
    }
    return keyManagerResp,nil
}

// UpdateKeyManager - method modifies  the "KeyManager"
func (svc *KeyManagerService) UpdateKeyManager(id string, obj *nimbleos.KeyManager) (*nimbleos.KeyManager, error) {
    if obj == nil {
        return nil,fmt.Errorf("UpdateKeyManager: invalid parameter specified, %v",obj)
    }

    keyManagerResp,err :=svc.objectSet.UpdateObject(id, obj)
    if err !=nil {
        return nil,err
    }
    return keyManagerResp,nil
}

// GetKeyManagerById - method returns a pointer to "KeyManager"
func (svc *KeyManagerService) GetKeyManagerById(id string) (*nimbleos.KeyManager, error) {
    if len(id) == 0 {
        return nil,fmt.Errorf("GetKeyManagerById: invalid parameter specified, %v",id)
    }

    keyManagerResp, err := svc.objectSet.GetObject(id)
    if err != nil {
        return nil,err
    }
    return keyManagerResp,nil
}

// GetKeyManagerByName - method returns a pointer "KeyManager"
func (svc *KeyManagerService) GetKeyManagerByName(name string) (*nimbleos.KeyManager, error) {
    params := &param.GetParams{
        Filter: &param.SearchFilter{
            FieldName: nimbleos.VolumeFields.Name,
            Operator:  param.EQUALS.String(),
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

    return keyManagerResp[0],nil
}
// DeleteKeyManager - deletes the "KeyManager"
func (svc *KeyManagerService) DeleteKeyManager(id string) error {
    if len(id) == 0 {
        return fmt.Errorf("DeleteKeyManager: invalid parameter specified, %s",id)
    }
    err := svc.objectSet.DeleteObject(id)
    if err != nil {
        return err
    }
    return nil
}

// RemoveKeyManager - remove external key manager. You must migrate the keys to an inactive external key manager before removing the active key manager. If you remove the active external key manager the passphrase is used to enable the internal key manager.
//   Required parameters:
//       id - ID of the external key manager.

//   Optional parameters:
//       passphrase - Passphrase used to protect the master key, required during deletion of external key manager.

func (svc *KeyManagerService) RemoveKeyManager( id string , passphrase *string )(error) {

     if  len(id) == 0  {
return fmt.Errorf("RemoveKeyManager: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.Remove( &id , passphrase )
     return err
}
// MigrateKeysKeyManager - migrate volume encryption keys from the active key manager to the destination id given in the input. After successfully migrating the encryption keys, the destination key manager is made the active key manager.
//   Required parameters:
//       id - ID of the destination external key manager.


func (svc *KeyManagerService) MigrateKeysKeyManager( id string )(error) {

     if  len(id) == 0  {
return fmt.Errorf("MigrateKeysKeyManager: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.MigrateKeys( &id )
     return err
}
