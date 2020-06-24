package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Represents Active Directory groups configured to manage the system.
 *
 */
const (
    userGroupPath = "user_groups"
)

/**
 * UserGroupObjectSet
*/
type UserGroupObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new UserGroup object
func (objectSet *UserGroupObjectSet) CreateObject(payload *model.UserGroup) (*model.UserGroup, error) {
	response, err := objectSet.Client.Post(userGroupPath, payload)
	return response.(*model.UserGroup), err
}

// UpdateObject Modify existing UserGroup object
func (objectSet *UserGroupObjectSet) UpdateObject(id string, payload *model.UserGroup) (*model.UserGroup, error) {
	response, err := objectSet.Client.Put(userGroupPath, id, payload)
	return response.(*model.UserGroup), err
}

// DeleteObject deletes the UserGroup object with the specified ID
func (objectSet *UserGroupObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(userGroupPath, id)
}

// GetObject returns a UserGroup object with the given ID
func (objectSet *UserGroupObjectSet) GetObject(id string) (*model.UserGroup, error) {
	response, err := objectSet.Client.Get(userGroupPath, id, model.UserGroup{})
	if response == nil {
		return nil, err
	}
	return response.(*model.UserGroup), err
}

// GetObjectList returns the list of UserGroup objects
func (objectSet *UserGroupObjectSet) GetObjectList() ([]*model.UserGroup, error) {
	response, err := objectSet.Client.List(userGroupPath)
	return buildUserGroupObjectSet(response), err
}

// GetObjectListFromParams returns the list of UserGroup objects using the given params query info
func (objectSet *UserGroupObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.UserGroup, error) {
	response, err := objectSet.Client.ListFromParams(userGroupPath, params)
	return buildUserGroupObjectSet(response), err
}

// generated function to build the appropriate response types
func buildUserGroupObjectSet(response interface{}) ([]*model.UserGroup) {
	values := reflect.ValueOf(response)
	results := make([]*model.UserGroup, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.UserGroup{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}