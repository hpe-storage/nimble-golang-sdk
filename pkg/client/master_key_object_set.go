// Copyright 2020 Hewlett Packard Enterprise Development LP


package client

import (

        "github.com/hpe-storage/common-host-libs/jsonutil"
        "reflect"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"

)

// Manage the master key. Data encryption keys for volumes are encrypted by using a master key that must be initialized before encrypted volumes can be created. The master key in
// turn is protected by a passphrase that is set when the master key is created. The passphrase may have to be entered to enable the master key when it is not available, for
// example, after an array reboot.
const (
    masterKeyPath = "master_key"
)

// MasterKeyObjectSet
type MasterKeyObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new MasterKey object
func (objectSet *MasterKeyObjectSet) CreateObject(payload *nimbleos.MasterKey) (*nimbleos.MasterKey, error) {
    resp, err := objectSet.Client.Post(masterKeyPath, payload, &nimbleos.MasterKey{})
    if err != nil {
        return nil, err
    }

    return resp.(*nimbleos.MasterKey), err
}

// UpdateObject Modify existing MasterKey object
func (objectSet *MasterKeyObjectSet) UpdateObject(id string, payload *nimbleos.MasterKey) (*nimbleos.MasterKey, error) {
    resp, err:= objectSet.Client.Put(masterKeyPath, id, payload, &nimbleos.MasterKey{})
    if err != nil {
        return nil, err
    }

    return resp.(*nimbleos.MasterKey), err
}

// DeleteObject deletes the MasterKey object with the specified ID
func (objectSet *MasterKeyObjectSet) DeleteObject(id string) error {
    err := objectSet.Client.Delete(masterKeyPath, id)
    if err != nil {
        return err
    }
    return nil
}

// GetObject returns a MasterKey object with the given ID
func (objectSet *MasterKeyObjectSet) GetObject(id string) (*nimbleos.MasterKey, error) {
    resp, err:= objectSet.Client.Get(masterKeyPath, id, &nimbleos.MasterKey{})
    if err != nil {
        return nil, err
    }

    // null check
    if resp == nil {
        return nil, nil
    }
    return resp.(*nimbleos.MasterKey), err
}

// GetObjectList returns the list of MasterKey objects
func (objectSet *MasterKeyObjectSet) GetObjectList() ([]*nimbleos.MasterKey, error) {
    resp, err:= objectSet.Client.List(masterKeyPath)
    if err != nil {
        return nil, err
    }
    return buildMasterKeyObjectSet(resp), err
}

// GetObjectListFromParams returns the list of MasterKey objects using the given params query info
func (objectSet *MasterKeyObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.MasterKey, error) {
    masterKeyObjectSetResp,err:= objectSet.Client.ListFromParams(masterKeyPath, params)
    if err != nil {
        return nil, err
    }
    return buildMasterKeyObjectSet(masterKeyObjectSetResp), err
}
// generated function to build the appropriate response types
func buildMasterKeyObjectSet(response interface{}) ([]*nimbleos.MasterKey) {
    values := reflect.ValueOf(response)
    results := make([]*nimbleos.MasterKey, values.Len())

    for i := 0; i < values.Len(); i++ {
        value := &nimbleos.MasterKey{}
        jsonutil.Decode(values.Index(i).Interface(), value)
        results[i] = value
    }

    return results
}

// List of supported actions on object sets

// PurgeInactive - Purges encryption keys that have been inactive for the age or longer. If you do not specify an age, the keys will be purged immediately.
func (objectSet *MasterKeyObjectSet) PurgeInactive ( id *string , age *int )(error) {
     purgeInactiveUri := masterKeyPath
     purgeInactiveUri = purgeInactiveUri + "/" + *id
     purgeInactiveUri = purgeInactiveUri + "/actions/" + "purge_inactive"

     payload := &struct {
                Id *string `json:"id,omitempty"`
                Age *int `json:"age,omitempty"`
     }{
            id,
            age,
     }

        _, err := objectSet.Client.Post(purgeInactiveUri, payload, &nimbleos.MasterKey{})
        return err
}
