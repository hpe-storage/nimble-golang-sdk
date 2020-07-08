// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
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
	newPayload, err := model.EncodeFolder(payload)
	resp, err := objectSet.Client.Post(folderPath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return model.DecodeFolder(resp)
}

// UpdateObject Modify existing Folder object
func (objectSet *FolderObjectSet) UpdateObject(id string, payload *model.Folder) (*model.Folder, error) {
	newPayload, err := model.EncodeFolder(payload)
	resp, err := objectSet.Client.Put(folderPath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return model.DecodeFolder(resp)
}

// DeleteObject deletes the Folder object with the specified ID
func (objectSet *FolderObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(folderPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Folder object with the given ID
func (objectSet *FolderObjectSet) GetObject(id string) (*model.Folder, error) {
	folderObjectSetResp, err := objectSet.Client.Get(folderPath, id, model.Folder{})
	if err != nil {
		return nil, err
	}

	// null check
	if folderObjectSetResp == nil {
		return nil, nil
	}
	return folderObjectSetResp.(*model.Folder), err
}

// GetObjectList returns the list of Folder objects
func (objectSet *FolderObjectSet) GetObjectList() ([]*model.Folder, error) {
	folderObjectSetResp, err := objectSet.Client.List(folderPath)
	if err != nil {
		return nil, err
	}
	return buildFolderObjectSet(folderObjectSetResp), err
}

// GetObjectListFromParams returns the list of Folder objects using the given params query info
func (objectSet *FolderObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Folder, error) {
	folderObjectSetResp, err := objectSet.Client.ListFromParams(folderPath, params)
	if err != nil {
		return nil, err
	}
	return buildFolderObjectSet(folderObjectSetResp), err
}

// generated function to build the appropriate response types
func buildFolderObjectSet(response interface{}) []*model.Folder {
	values := reflect.ValueOf(response)
	results := make([]*model.Folder, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Folder{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
