// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Folders are a way to group volumes, as well as a way to apply space constraints to them.
const (
    folderPath = "folders"
)

// FolderObjectSet
type FolderObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Folder object
func (objectSet *FolderObjectSet) CreateObject(payload *model.Folder) (*model.Folder, error) {
	response, err := objectSet.Client.Post(folderPath, payload)
	return response.(*model.Folder), err
}

// UpdateObject Modify existing Folder object
func (objectSet *FolderObjectSet) UpdateObject(id string, payload *model.Folder) (*model.Folder, error) {
	response, err := objectSet.Client.Put(folderPath, id, payload)
	return response.(*model.Folder), err
}

// DeleteObject deletes the Folder object with the specified ID
func (objectSet *FolderObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(folderPath, id)
}

// GetObject returns a Folder object with the given ID
func (objectSet *FolderObjectSet) GetObject(id string) (*model.Folder, error) {
	response, err := objectSet.Client.Get(folderPath, id, model.Folder{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Folder), err
}

// GetObjectList returns the list of Folder objects
func (objectSet *FolderObjectSet) GetObjectList() ([]*model.Folder, error) {
	response, err := objectSet.Client.List(folderPath)
	return buildFolderObjectSet(response), err
}

// GetObjectListFromParams returns the list of Folder objects using the given params query info
func (objectSet *FolderObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Folder, error) {
	response, err := objectSet.Client.ListFromParams(folderPath, params)
	return buildFolderObjectSet(response), err
}

// generated function to build the appropriate response types
func buildFolderObjectSet(response interface{}) ([]*model.Folder) {
	values := reflect.ValueOf(response)
	results := make([]*model.Folder, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Folder{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}