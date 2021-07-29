// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Manage protection templates. Protection templates are sets of snapshot schedules, replication schedules, and retention limits that can be used to prefill the protection
// information when creating new volume collections. A volume collection, once created, is not affected by edits to the protection template that was used to create it. All the
// volumes assigned to a volume collection use the same settings. You cannot edit or delete the predefined protection templates provided by storage array, but you can create custom
// protection templates as needed.
const (
	protectionTemplatePath = "protection_templates"
)

// ProtectionTemplateObjectSet
type ProtectionTemplateObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new ProtectionTemplate object
func (objectSet *ProtectionTemplateObjectSet) CreateObject(payload *nimbleos.ProtectionTemplate) (*nimbleos.ProtectionTemplate, error) {
	resp, err := objectSet.Client.Post(protectionTemplatePath, payload, &nimbleos.ProtectionTemplate{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.ProtectionTemplate), err
}

// UpdateObject Modify existing ProtectionTemplate object
func (objectSet *ProtectionTemplateObjectSet) UpdateObject(id string, payload *nimbleos.ProtectionTemplate) (*nimbleos.ProtectionTemplate, error) {
	resp, err := objectSet.Client.Put(protectionTemplatePath, id, payload, &nimbleos.ProtectionTemplate{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.ProtectionTemplate), err
}

// DeleteObject deletes the ProtectionTemplate object with the specified ID
func (objectSet *ProtectionTemplateObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(protectionTemplatePath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a ProtectionTemplate object with the given ID
func (objectSet *ProtectionTemplateObjectSet) GetObject(id string) (*nimbleos.ProtectionTemplate, error) {
	resp, err := objectSet.Client.Get(protectionTemplatePath, id, &nimbleos.ProtectionTemplate{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.ProtectionTemplate), err
}

// GetObjectList returns the list of ProtectionTemplate objects
func (objectSet *ProtectionTemplateObjectSet) GetObjectList() ([]*nimbleos.ProtectionTemplate, error) {
	resp, err := objectSet.Client.List(protectionTemplatePath)
	if err != nil {
		return nil, err
	}
	return buildProtectionTemplateObjectSet(resp), err
}

// GetObjectListFromParams returns the list of ProtectionTemplate objects using the given params query info
func (objectSet *ProtectionTemplateObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.ProtectionTemplate, error) {
	protectionTemplateObjectSetResp, err := objectSet.Client.ListFromParams(protectionTemplatePath, params)
	if err != nil {
		return nil, err
	}
	return buildProtectionTemplateObjectSet(protectionTemplateObjectSetResp), err
}

// generated function to build the appropriate response types
func buildProtectionTemplateObjectSet(response interface{}) []*nimbleos.ProtectionTemplate {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.ProtectionTemplate, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.ProtectionTemplate{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
