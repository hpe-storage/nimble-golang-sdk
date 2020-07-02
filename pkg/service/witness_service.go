// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Witness Service - Manage witness host configuration.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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

// GetWitnesses - method returns a array of pointers of type "Witnesses"
func (svc *WitnessService) GetWitnesses(params *util.GetParams) ([]*model.Witness, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateWitness - method creates a "Witness"
func (svc *WitnessService) CreateWitness(obj *model.Witness) (*model.Witness, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateWitness - method modifies  the "Witness" 
func (svc *WitnessService) UpdateWitness(id string, obj *model.Witness) (*model.Witness, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetWitnessById - method returns a pointer to "Witness"
func (svc *WitnessService) GetWitnessById(id string) (*model.Witness, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteWitness - deletes the "Witness"
func (svc *WitnessService) DeleteWitness(id string) error {
	return svc.objectSet.DeleteObject(id)
}
