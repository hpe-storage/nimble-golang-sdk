// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Manage the master key. Data encryption keys for volumes are encrypted by using a master key that must be initialized before encrypted volumes can be created. The master key in turn is protected by a passphrase that is set when the master key is created. The passphrase may have to be entered to enable the master key when it is not available, for example, after an array reboot.
 *
 */
const (
    masterKeyPath = "master_key"
)

/**
 * MasterKeyObjectSet
*/
type MasterKeyObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new MasterKey object
func (objectSet *MasterKeyObjectSet) CreateObject(payload *model.MasterKey) (*model.MasterKey, error) {
	response, err := objectSet.Client.Post(masterKeyPath, payload)
	return response.(*model.MasterKey), err
}

// UpdateObject Modify existing MasterKey object
func (objectSet *MasterKeyObjectSet) UpdateObject(id string, payload *model.MasterKey) (*model.MasterKey, error) {
	response, err := objectSet.Client.Put(masterKeyPath, id, payload)
	return response.(*model.MasterKey), err
}

// DeleteObject deletes the MasterKey object with the specified ID
func (objectSet *MasterKeyObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(masterKeyPath, id)
}

// GetObject returns a MasterKey object with the given ID
func (objectSet *MasterKeyObjectSet) GetObject(id string) (*model.MasterKey, error) {
	response, err := objectSet.Client.Get(masterKeyPath, id, model.MasterKey{})
	if response == nil {
		return nil, err
	}
	return response.(*model.MasterKey), err
}

// GetObjectList returns the list of MasterKey objects
func (objectSet *MasterKeyObjectSet) GetObjectList() ([]*model.MasterKey, error) {
	response, err := objectSet.Client.List(masterKeyPath)
	return buildMasterKeyObjectSet(response), err
}

// GetObjectListFromParams returns the list of MasterKey objects using the given params query info
func (objectSet *MasterKeyObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.MasterKey, error) {
	response, err := objectSet.Client.ListFromParams(masterKeyPath, params)
	return buildMasterKeyObjectSet(response), err
}

// generated function to build the appropriate response types
func buildMasterKeyObjectSet(response interface{}) ([]*model.MasterKey) {
	values := reflect.ValueOf(response)
	results := make([]*model.MasterKey, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.MasterKey{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}