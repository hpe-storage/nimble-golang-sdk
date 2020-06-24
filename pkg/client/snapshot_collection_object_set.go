package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Snapshot collections are collections of scheduled snapshots that are taken from volumes sharing a volume collection. Snapshot collections are replicated in the order that the collections were taken.
 *
 */
const (
    snapshotCollectionPath = "snapshot_collections"
)

/**
 * SnapshotCollectionObjectSet
*/
type SnapshotCollectionObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new SnapshotCollection object
func (objectSet *SnapshotCollectionObjectSet) CreateObject(payload *model.SnapshotCollection) (*model.SnapshotCollection, error) {
	response, err := objectSet.Client.Post(snapshotCollectionPath, payload)
	return response.(*model.SnapshotCollection), err
}

// UpdateObject Modify existing SnapshotCollection object
func (objectSet *SnapshotCollectionObjectSet) UpdateObject(id string, payload *model.SnapshotCollection) (*model.SnapshotCollection, error) {
	response, err := objectSet.Client.Put(snapshotCollectionPath, id, payload)
	return response.(*model.SnapshotCollection), err
}

// DeleteObject deletes the SnapshotCollection object with the specified ID
func (objectSet *SnapshotCollectionObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(snapshotCollectionPath, id)
}

// GetObject returns a SnapshotCollection object with the given ID
func (objectSet *SnapshotCollectionObjectSet) GetObject(id string) (*model.SnapshotCollection, error) {
	response, err := objectSet.Client.Get(snapshotCollectionPath, id, model.SnapshotCollection{})
	if response == nil {
		return nil, err
	}
	return response.(*model.SnapshotCollection), err
}

// GetObjectList returns the list of SnapshotCollection objects
func (objectSet *SnapshotCollectionObjectSet) GetObjectList() ([]*model.SnapshotCollection, error) {
	response, err := objectSet.Client.List(snapshotCollectionPath)
	return buildSnapshotCollectionObjectSet(response), err
}

// GetObjectListFromParams returns the list of SnapshotCollection objects using the given params query info
func (objectSet *SnapshotCollectionObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.SnapshotCollection, error) {
	response, err := objectSet.Client.ListFromParams(snapshotCollectionPath, params)
	return buildSnapshotCollectionObjectSet(response), err
}

// generated function to build the appropriate response types
func buildSnapshotCollectionObjectSet(response interface{}) ([]*model.SnapshotCollection) {
	values := reflect.ValueOf(response)
	results := make([]*model.SnapshotCollection, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.SnapshotCollection{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}