package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Lists discovered arrays that are not members of any group and are in the same subnet.
 *
 */
const (
    uninitializedArrayPath = "uninitialized_arrays"
)

/**
 * UninitializedArrayObjectSet
*/
type UninitializedArrayObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new UninitializedArray object
func (objectSet *UninitializedArrayObjectSet) CreateObject(payload *model.UninitializedArray) (*model.UninitializedArray, error) {
	response, err := objectSet.Client.Post(uninitializedArrayPath, payload)
	return response.(*model.UninitializedArray), err
}

// UpdateObject Modify existing UninitializedArray object
func (objectSet *UninitializedArrayObjectSet) UpdateObject(id string, payload *model.UninitializedArray) (*model.UninitializedArray, error) {
	response, err := objectSet.Client.Put(uninitializedArrayPath, id, payload)
	return response.(*model.UninitializedArray), err
}

// DeleteObject deletes the UninitializedArray object with the specified ID
func (objectSet *UninitializedArrayObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(uninitializedArrayPath, id)
}

// GetObject returns a UninitializedArray object with the given ID
func (objectSet *UninitializedArrayObjectSet) GetObject(id string) (*model.UninitializedArray, error) {
	response, err := objectSet.Client.Get(uninitializedArrayPath, id, model.UninitializedArray{})
	if response == nil {
		return nil, err
	}
	return response.(*model.UninitializedArray), err
}

// GetObjectList returns the list of UninitializedArray objects
func (objectSet *UninitializedArrayObjectSet) GetObjectList() ([]*model.UninitializedArray, error) {
	response, err := objectSet.Client.List(uninitializedArrayPath)
	return buildUninitializedArrayObjectSet(response), err
}

// GetObjectListFromParams returns the list of UninitializedArray objects using the given params query info
func (objectSet *UninitializedArrayObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.UninitializedArray, error) {
	response, err := objectSet.Client.ListFromParams(uninitializedArrayPath, params)
	return buildUninitializedArrayObjectSet(response), err
}

// generated function to build the appropriate response types
func buildUninitializedArrayObjectSet(response interface{}) ([]*model.UninitializedArray) {
	values := reflect.ValueOf(response)
	results := make([]*model.UninitializedArray, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.UninitializedArray{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}