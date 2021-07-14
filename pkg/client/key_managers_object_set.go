// Copyright 2020 Hewlett Packard Enterprise Development LP


package client

import (

    "fmt"
        "github.com/hpe-storage/common-host-libs/jsonutil"
        "reflect"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"

)

// Key Manager stores encryption keys for the array volumes / dedupe domains.
const (
    keyManagerPath = "key_managers"
)

// KeyManagerObjectSet
type KeyManagerObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new KeyManager object
func (objectSet *KeyManagerObjectSet) CreateObject(payload *nimbleos.KeyManager) (*nimbleos.KeyManager, error) {
    resp, err := objectSet.Client.Post(keyManagerPath, payload, &nimbleos.KeyManager{})
    if err != nil {
        return nil, err
    }

    return resp.(*nimbleos.KeyManager), err
}

// UpdateObject Modify existing KeyManager object
func (objectSet *KeyManagerObjectSet) UpdateObject(id string, payload *nimbleos.KeyManager) (*nimbleos.KeyManager, error) {
    resp, err:= objectSet.Client.Put(keyManagerPath, id, payload, &nimbleos.KeyManager{})
    if err != nil {
        return nil, err
    }

    return resp.(*nimbleos.KeyManager), err
}

// DeleteObject deletes the KeyManager object with the specified ID
func (objectSet *KeyManagerObjectSet) DeleteObject(id string) error {
    return fmt.Errorf("Unsupported operation 'delete' on KeyManager")
}

// GetObject returns a KeyManager object with the given ID
func (objectSet *KeyManagerObjectSet) GetObject(id string) (*nimbleos.KeyManager, error) {
    resp, err:= objectSet.Client.Get(keyManagerPath, id, &nimbleos.KeyManager{})
    if err != nil {
        return nil, err
    }

    // null check
    if resp == nil {
        return nil, nil
    }
    return resp.(*nimbleos.KeyManager), err
}

// GetObjectList returns the list of KeyManager objects
func (objectSet *KeyManagerObjectSet) GetObjectList() ([]*nimbleos.KeyManager, error) {
    resp, err:= objectSet.Client.List(keyManagerPath)
    if err != nil {
        return nil, err
    }
    return buildKeyManagerObjectSet(resp), err
}

// GetObjectListFromParams returns the list of KeyManager objects using the given params query info
func (objectSet *KeyManagerObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.KeyManager, error) {
    keyManagerObjectSetResp,err:= objectSet.Client.ListFromParams(keyManagerPath, params)
    if err != nil {
        return nil, err
    }
    return buildKeyManagerObjectSet(keyManagerObjectSetResp), err
}
// generated function to build the appropriate response types
func buildKeyManagerObjectSet(response interface{}) ([]*nimbleos.KeyManager) {
    values := reflect.ValueOf(response)
    results := make([]*nimbleos.KeyManager, values.Len())

    for i := 0; i < values.Len(); i++ {
        value := &nimbleos.KeyManager{}
        jsonutil.Decode(values.Index(i).Interface(), value)
        results[i] = value
    }

    return results
}

// List of supported actions on object sets

// Remove - Remove external key manager. You must migrate the keys to an inactive external key manager before removing the active key manager. If you remove the active external key manager the passphrase is used to enable the internal key manager.
func (objectSet *KeyManagerObjectSet) Remove ( id *string , passphrase *string )(error) {
     removeUri := keyManagerPath
     removeUri = removeUri + "/" + *id
     removeUri = removeUri + "/actions/" + "remove"

     payload := &struct {
                Id *string `json:"id,omitempty"`
                Passphrase *string `json:"passphrase,omitempty"`
     }{
            id,
            passphrase,
     }

        _, err := objectSet.Client.Post(removeUri, payload, &nimbleos.KeyManager{})
        return err
}

// MigrateKeys - Migrate volume encryption keys from the active key manager to the destination id given in the input. After successfully migrating the encryption keys, the destination key manager is made the active key manager.
func (objectSet *KeyManagerObjectSet) MigrateKeys ( id *string )(error) {
     migrateKeysUri := keyManagerPath
     migrateKeysUri = migrateKeysUri + "/" + *id
     migrateKeysUri = migrateKeysUri + "/actions/" + "migrate_keys"

     payload := &struct {
                Id *string `json:"id,omitempty"`
     }{
            id,
     }

        _, err := objectSet.Client.Post(migrateKeysUri, payload, &nimbleos.KeyManager{})
        return err
}
