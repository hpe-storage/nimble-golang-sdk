package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Retrieve roles and privileges for role-based access control.
 *
 */
const (
    rolePath = "roles"
)

/**
 * RoleObjectSet
*/
type RoleObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Role object
func (objectSet *RoleObjectSet) CreateObject(payload *model.Role) (*model.Role, error) {
	response, err := objectSet.Client.Post(rolePath, payload)
	return response.(*model.Role), err
}

// UpdateObject Modify existing Role object
func (objectSet *RoleObjectSet) UpdateObject(id string, payload *model.Role) (*model.Role, error) {
	response, err := objectSet.Client.Put(rolePath, id, payload)
	return response.(*model.Role), err
}

// DeleteObject deletes the Role object with the specified ID
func (objectSet *RoleObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(rolePath, id)
}

// GetObject returns a Role object with the given ID
func (objectSet *RoleObjectSet) GetObject(id string) (*model.Role, error) {
	response, err := objectSet.Client.Get(rolePath, id, model.Role{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Role), err
}

// GetObjectList returns the list of Role objects
func (objectSet *RoleObjectSet) GetObjectList() ([]*model.Role, error) {
	response, err := objectSet.Client.List(rolePath)
	return buildRoleObjectSet(response), err
}

// GetObjectListFromParams returns the list of Role objects using the given params query info
func (objectSet *RoleObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Role, error) {
	response, err := objectSet.Client.ListFromParams(rolePath, params)
	return buildRoleObjectSet(response), err
}

// generated function to build the appropriate response types
func buildRoleObjectSet(response interface{}) ([]*model.Role) {
	values := reflect.ValueOf(response)
	results := make([]*model.Role, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Role{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}