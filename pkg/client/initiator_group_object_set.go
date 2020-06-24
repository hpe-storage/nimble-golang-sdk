package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Manage initiator groups for initiator authentication. An initiator group is a set of initiators that can be configured as part of your ACL to access a specific volume through group membership.
 *
 */
const (
    initiatorGroupPath = "initiator_groups"
)

/**
 * InitiatorGroupObjectSet
*/
type InitiatorGroupObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new InitiatorGroup object
func (objectSet *InitiatorGroupObjectSet) CreateObject(payload *model.InitiatorGroup) (*model.InitiatorGroup, error) {
	response, err := objectSet.Client.Post(initiatorGroupPath, payload)
	return response.(*model.InitiatorGroup), err
}

// UpdateObject Modify existing InitiatorGroup object
func (objectSet *InitiatorGroupObjectSet) UpdateObject(id string, payload *model.InitiatorGroup) (*model.InitiatorGroup, error) {
	response, err := objectSet.Client.Put(initiatorGroupPath, id, payload)
	return response.(*model.InitiatorGroup), err
}

// DeleteObject deletes the InitiatorGroup object with the specified ID
func (objectSet *InitiatorGroupObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(initiatorGroupPath, id)
}

// GetObject returns a InitiatorGroup object with the given ID
func (objectSet *InitiatorGroupObjectSet) GetObject(id string) (*model.InitiatorGroup, error) {
	response, err := objectSet.Client.Get(initiatorGroupPath, id, model.InitiatorGroup{})
	if response == nil {
		return nil, err
	}
	return response.(*model.InitiatorGroup), err
}

// GetObjectList returns the list of InitiatorGroup objects
func (objectSet *InitiatorGroupObjectSet) GetObjectList() ([]*model.InitiatorGroup, error) {
	response, err := objectSet.Client.List(initiatorGroupPath)
	return buildInitiatorGroupObjectSet(response), err
}

// GetObjectListFromParams returns the list of InitiatorGroup objects using the given params query info
func (objectSet *InitiatorGroupObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.InitiatorGroup, error) {
	response, err := objectSet.Client.ListFromParams(initiatorGroupPath, params)
	return buildInitiatorGroupObjectSet(response), err
}

// generated function to build the appropriate response types
func buildInitiatorGroupObjectSet(response interface{}) ([]*model.InitiatorGroup) {
	values := reflect.ValueOf(response)
	results := make([]*model.InitiatorGroup, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.InitiatorGroup{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}