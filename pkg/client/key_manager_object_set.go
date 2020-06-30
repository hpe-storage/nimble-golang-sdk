// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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
func (objectSet *KeyManagerObjectSet) CreateObject(payload *model.KeyManager) (*model.KeyManager, error) {
	response, err := objectSet.Client.Post(keyManagerPath, payload)
	return response.(*model.KeyManager), err
}

// UpdateObject Modify existing KeyManager object
func (objectSet *KeyManagerObjectSet) UpdateObject(id string, payload *model.KeyManager) (*model.KeyManager, error) {
	response, err := objectSet.Client.Put(keyManagerPath, id, payload)
	return response.(*model.KeyManager), err
}

// DeleteObject deletes the KeyManager object with the specified ID
func (objectSet *KeyManagerObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(keyManagerPath, id)
}

// GetObject returns a KeyManager object with the given ID
func (objectSet *KeyManagerObjectSet) GetObject(id string) (*model.KeyManager, error) {
	response, err := objectSet.Client.Get(keyManagerPath, id, model.KeyManager{})
	if response == nil {
		return nil, err
	}
	return response.(*model.KeyManager), err
}

// GetObjectList returns the list of KeyManager objects
func (objectSet *KeyManagerObjectSet) GetObjectList() ([]*model.KeyManager, error) {
	response, err := objectSet.Client.List(keyManagerPath)
	return buildKeyManagerObjectSet(response), err
}

// GetObjectListFromParams returns the list of KeyManager objects using the given params query info
func (objectSet *KeyManagerObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.KeyManager, error) {
	response, err := objectSet.Client.ListFromParams(keyManagerPath, params)
	return buildKeyManagerObjectSet(response), err
}

// generated function to build the appropriate response types
func buildKeyManagerObjectSet(response interface{}) ([]*model.KeyManager) {
	values := reflect.ValueOf(response)
	results := make([]*model.KeyManager, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.KeyManager{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}