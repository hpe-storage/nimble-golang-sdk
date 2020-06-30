// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Folder set represents a set of folder each on separate pools that represent a group-scope datastore spanning the entire group.
const (
    folderSetPath = "folder_sets"
)

// FolderSetObjectSet
type FolderSetObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new FolderSet object
func (objectSet *FolderSetObjectSet) CreateObject(payload *model.FolderSet) (*model.FolderSet, error) {
	response, err := objectSet.Client.Post(folderSetPath, payload)
	return response.(*model.FolderSet), err
}

// UpdateObject Modify existing FolderSet object
func (objectSet *FolderSetObjectSet) UpdateObject(id string, payload *model.FolderSet) (*model.FolderSet, error) {
	response, err := objectSet.Client.Put(folderSetPath, id, payload)
	return response.(*model.FolderSet), err
}

// DeleteObject deletes the FolderSet object with the specified ID
func (objectSet *FolderSetObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(folderSetPath, id)
}

// GetObject returns a FolderSet object with the given ID
func (objectSet *FolderSetObjectSet) GetObject(id string) (*model.FolderSet, error) {
	response, err := objectSet.Client.Get(folderSetPath, id, model.FolderSet{})
	if response == nil {
		return nil, err
	}
	return response.(*model.FolderSet), err
}

// GetObjectList returns the list of FolderSet objects
func (objectSet *FolderSetObjectSet) GetObjectList() ([]*model.FolderSet, error) {
	response, err := objectSet.Client.List(folderSetPath)
	return buildFolderSetObjectSet(response), err
}

// GetObjectListFromParams returns the list of FolderSet objects using the given params query info
func (objectSet *FolderSetObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.FolderSet, error) {
	response, err := objectSet.Client.ListFromParams(folderSetPath, params)
	return buildFolderSetObjectSet(response), err
}

// generated function to build the appropriate response types
func buildFolderSetObjectSet(response interface{}) ([]*model.FolderSet) {
	values := reflect.ValueOf(response)
	results := make([]*model.FolderSet, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.FolderSet{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}