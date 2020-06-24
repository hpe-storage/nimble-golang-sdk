package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Manage protection templates. Protection templates are sets of snapshot schedules, replication schedules, and retention limits that can be used to prefill the protection information when creating new volume collections. A volume collection, once created, is not affected by edits to the protection template that was used to create it. All the volumes assigned to a volume collection use the same settings. You cannot edit or delete the predefined protection templates provided by storage array, but you can create custom protection templates as needed.
 *
 */
const (
    protectionTemplatePath = "protection_templates"
)

/**
 * ProtectionTemplateObjectSet
*/
type ProtectionTemplateObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new ProtectionTemplate object
func (objectSet *ProtectionTemplateObjectSet) CreateObject(payload *model.ProtectionTemplate) (*model.ProtectionTemplate, error) {
	response, err := objectSet.Client.Post(protectionTemplatePath, payload)
	return response.(*model.ProtectionTemplate), err
}

// UpdateObject Modify existing ProtectionTemplate object
func (objectSet *ProtectionTemplateObjectSet) UpdateObject(id string, payload *model.ProtectionTemplate) (*model.ProtectionTemplate, error) {
	response, err := objectSet.Client.Put(protectionTemplatePath, id, payload)
	return response.(*model.ProtectionTemplate), err
}

// DeleteObject deletes the ProtectionTemplate object with the specified ID
func (objectSet *ProtectionTemplateObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(protectionTemplatePath, id)
}

// GetObject returns a ProtectionTemplate object with the given ID
func (objectSet *ProtectionTemplateObjectSet) GetObject(id string) (*model.ProtectionTemplate, error) {
	response, err := objectSet.Client.Get(protectionTemplatePath, id, model.ProtectionTemplate{})
	if response == nil {
		return nil, err
	}
	return response.(*model.ProtectionTemplate), err
}

// GetObjectList returns the list of ProtectionTemplate objects
func (objectSet *ProtectionTemplateObjectSet) GetObjectList() ([]*model.ProtectionTemplate, error) {
	response, err := objectSet.Client.List(protectionTemplatePath)
	return buildProtectionTemplateObjectSet(response), err
}

// GetObjectListFromParams returns the list of ProtectionTemplate objects using the given params query info
func (objectSet *ProtectionTemplateObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.ProtectionTemplate, error) {
	response, err := objectSet.Client.ListFromParams(protectionTemplatePath, params)
	return buildProtectionTemplateObjectSet(response), err
}

// generated function to build the appropriate response types
func buildProtectionTemplateObjectSet(response interface{}) ([]*model.ProtectionTemplate) {
	values := reflect.ValueOf(response)
	results := make([]*model.ProtectionTemplate, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.ProtectionTemplate{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}