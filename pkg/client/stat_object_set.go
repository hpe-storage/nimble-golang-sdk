package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Access generic stats interface via REST for internal testing.
 *
 */
const (
    statPath = "stats"
)

/**
 * StatObjectSet
*/
type StatObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Stat object
func (objectSet *StatObjectSet) CreateObject(payload *model.Stat) (*model.Stat, error) {
	response, err := objectSet.Client.Post(statPath, payload)
	return response.(*model.Stat), err
}

// UpdateObject Modify existing Stat object
func (objectSet *StatObjectSet) UpdateObject(id string, payload *model.Stat) (*model.Stat, error) {
	response, err := objectSet.Client.Put(statPath, id, payload)
	return response.(*model.Stat), err
}

// DeleteObject deletes the Stat object with the specified ID
func (objectSet *StatObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(statPath, id)
}

// GetObject returns a Stat object with the given ID
func (objectSet *StatObjectSet) GetObject(id string) (*model.Stat, error) {
	response, err := objectSet.Client.Get(statPath, id, model.Stat{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Stat), err
}

// GetObjectList returns the list of Stat objects
func (objectSet *StatObjectSet) GetObjectList() ([]*model.Stat, error) {
	response, err := objectSet.Client.List(statPath)
	return buildStatObjectSet(response), err
}

// GetObjectListFromParams returns the list of Stat objects using the given params query info
func (objectSet *StatObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Stat, error) {
	response, err := objectSet.Client.ListFromParams(statPath, params)
	return buildStatObjectSet(response), err
}

// generated function to build the appropriate response types
func buildStatObjectSet(response interface{}) ([]*model.Stat) {
	values := reflect.ValueOf(response)
	results := make([]*model.Stat, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Stat{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}