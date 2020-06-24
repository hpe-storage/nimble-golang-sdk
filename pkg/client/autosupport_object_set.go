package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Get status of autosupport.
 *
 */
const (
    autosupportPath = "autosupport"
)

/**
 * AutosupportObjectSet
*/
type AutosupportObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Autosupport object
func (objectSet *AutosupportObjectSet) CreateObject(payload *model.Autosupport) (*model.Autosupport, error) {
	response, err := objectSet.Client.Post(autosupportPath, payload)
	return response.(*model.Autosupport), err
}

// UpdateObject Modify existing Autosupport object
func (objectSet *AutosupportObjectSet) UpdateObject(id string, payload *model.Autosupport) (*model.Autosupport, error) {
	response, err := objectSet.Client.Put(autosupportPath, id, payload)
	return response.(*model.Autosupport), err
}

// DeleteObject deletes the Autosupport object with the specified ID
func (objectSet *AutosupportObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(autosupportPath, id)
}

// GetObject returns a Autosupport object with the given ID
func (objectSet *AutosupportObjectSet) GetObject(id string) (*model.Autosupport, error) {
	response, err := objectSet.Client.Get(autosupportPath, id, model.Autosupport{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Autosupport), err
}

// GetObjectList returns the list of Autosupport objects
func (objectSet *AutosupportObjectSet) GetObjectList() ([]*model.Autosupport, error) {
	response, err := objectSet.Client.List(autosupportPath)
	return buildAutosupportObjectSet(response), err
}

// GetObjectListFromParams returns the list of Autosupport objects using the given params query info
func (objectSet *AutosupportObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Autosupport, error) {
	response, err := objectSet.Client.ListFromParams(autosupportPath, params)
	return buildAutosupportObjectSet(response), err
}

// generated function to build the appropriate response types
func buildAutosupportObjectSet(response interface{}) ([]*model.Autosupport) {
	values := reflect.ValueOf(response)
	results := make([]*model.Autosupport, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Autosupport{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}