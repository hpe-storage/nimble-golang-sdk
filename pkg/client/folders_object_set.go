// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (objectSet *FolderObjectSet) CreateObject(payload *nimbleos.Folder) (*nimbleos.Folder, error) {
	resp, err := objectSet.Client.Post(folderPath, payload, &nimbleos.Folder{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Folder), err
}

// UpdateObject Modify existing Folder object
func (objectSet *FolderObjectSet) UpdateObject(id string, payload *nimbleos.Folder) (*nimbleos.Folder, error) {
	resp, err := objectSet.Client.Put(folderPath, id, payload, &nimbleos.Folder{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Folder), err
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
func (objectSet *FolderObjectSet) GetObject(id string) (*nimbleos.Folder, error) {
	resp, err := objectSet.Client.Get(folderPath, id, &nimbleos.Folder{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Folder), err
}

// GetObjectList returns the list of Folder objects
func (objectSet *FolderObjectSet) GetObjectList() ([]*nimbleos.Folder, error) {
	resp, err := objectSet.Client.List(folderPath)
	if err != nil {
		return nil, err
	}
	return buildFolderObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Folder objects using the given params query info
func (objectSet *FolderObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Folder, error) {
	folderObjectSetResp, err := objectSet.Client.ListFromParams(folderPath, params)
	if err != nil {
		return nil, err
	}
	return buildFolderObjectSet(folderObjectSetResp), err
}

// generated function to build the appropriate response types
func buildFolderObjectSet(response interface{}) []*nimbleos.Folder {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Folder, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Folder{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// SetDedupe - Set dedupe enabled/disabled for all applicable volumes inside a folder.
func (objectSet *FolderObjectSet) SetDedupe(dedupeEnabled *bool, id *string) error {

	setDedupeUri := folderPath
	setDedupeUri = setDedupeUri + "/" + *id
	setDedupeUri = setDedupeUri + "/actions/" + "set_dedupe"

	payload := &struct {
		DedupeEnabled *bool   `json:"dedupe_enabled,omitempty"`
		Id            *string `json:"id,omitempty"`
	}{
		dedupeEnabled,
		id,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(setDedupeUri, payload, &emptyStruct)
	return err
}
