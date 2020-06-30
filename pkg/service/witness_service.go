// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Witness Service - Manage witness host configuration.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// WitnessService type 
type WitnessService struct {
	objectSet *client.WitnessObjectSet
}

// NewWitnessService - method to initialize "WitnessService" 
func NewWitnessService(gs *NsGroupService) (*WitnessService) {
	objectSet := gs.client.GetWitnessObjectSet()
	return &WitnessService{objectSet: objectSet}
}

// GetWitnesss - method returns a array of pointers of type "Witnesss"
func (svc *WitnessService) GetWitnesss(params *util.GetParams) ([]*model.Witness, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetWitnesssWithFields - method returns a array of pointers of type "Witness" 
func (svc *WitnessService) GetWitnesssWithFields(fields []string) ([]*model.Witness, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateWitness - method creates a "Witness"
func (svc *WitnessService) CreateWitness(obj *model.Witness) (*model.Witness, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditWitness - method modifies  the "Witness" 
func (svc *WitnessService) EditWitness(id string, obj *model.Witness) (*model.Witness, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyWitness - private method for more than one element check. 
func onlyWitness(objs []*model.Witness) (*model.Witness, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Witness found with the given filter")
	}

	return objs[0], nil
}

 
// GetWitnesssByID - method returns associative a array of pointers of type "Witness", filter by Id
func (svc *WitnessService) GetWitnesssByID(pool *model.Pool, fields []string) (map[string]*model.Witness, error) {
	params := &util.GetParams{}

	// make sure ID field is selected
	if _, found := params.FindField("id"); !found {
		fields = append(fields, "id")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetWitnesss(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Witness)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetWitnessById - method returns a pointer to "Witness"
func (svc *WitnessService) GetWitnessById(id string) (*model.Witness, error) {
	return svc.objectSet.GetObject(id)
}


