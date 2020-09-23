// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
	"strings"
)

// Manage initiators in initiator groups. An initiator group has a set of initiators that can be configured as part of your ACL to access a specific volume through group membership.
const (
	initiatorPath = "initiators"
)

// InitiatorObjectSet
type InitiatorObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Initiator object
func (objectSet *InitiatorObjectSet) CreateObject(payload *nimbleos.Initiator) (*nimbleos.Initiator, error) {
	resp, err := objectSet.Client.Post(initiatorPath, payload, &nimbleos.Initiator{})
	if err != nil {
		//process http code 202
		if strings.Contains(err.Error(), "status (202)") {
			if resp != nil {
				ID := resp.(string)
				// Get object
				return objectSet.GetObject(ID)
			}
		}
		return nil, err
	}

	return resp.(*nimbleos.Initiator), err
}

// UpdateObject Modify existing Initiator object
func (objectSet *InitiatorObjectSet) UpdateObject(id string, payload *nimbleos.Initiator) (*nimbleos.Initiator, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Initiator")
}

// DeleteObject deletes the Initiator object with the specified ID
func (objectSet *InitiatorObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(initiatorPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Initiator object with the given ID
func (objectSet *InitiatorObjectSet) GetObject(id string) (*nimbleos.Initiator, error) {
	resp, err := objectSet.Client.Get(initiatorPath, id, &nimbleos.Initiator{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Initiator), err
}

// GetObjectList returns the list of Initiator objects
func (objectSet *InitiatorObjectSet) GetObjectList() ([]*nimbleos.Initiator, error) {
	resp, err := objectSet.Client.List(initiatorPath)
	if err != nil {
		return nil, err
	}
	return buildInitiatorObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Initiator objects using the given params query info
func (objectSet *InitiatorObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Initiator, error) {
	initiatorObjectSetResp, err := objectSet.Client.ListFromParams(initiatorPath, params)
	if err != nil {
		return nil, err
	}
	return buildInitiatorObjectSet(initiatorObjectSetResp), err
}

// generated function to build the appropriate response types
func buildInitiatorObjectSet(response interface{}) []*nimbleos.Initiator {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Initiator, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Initiator{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
