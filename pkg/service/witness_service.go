// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Witness Service - Manage witness host configuration.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// WitnessService type
type WitnessService struct {
	objectSet *client.WitnessObjectSet
}

// NewWitnessService - method to initialize "WitnessService"
func NewWitnessService(gs *NsGroupService) *WitnessService {
	objectSet := gs.client.GetWitnessObjectSet()
	return &WitnessService{objectSet: objectSet}
}

// GetWitnesses - method returns a array of pointers of type "Witnesses"
func (svc *WitnessService) GetWitnesses(params *param.GetParams) ([]*nimbleos.Witness, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	witnessResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return witnessResp, nil
}

// CreateWitness - method creates a "Witness"
func (svc *WitnessService) CreateWitness(obj *nimbleos.Witness) (*nimbleos.Witness, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	witnessResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return witnessResp, nil
}

// UpdateWitness - method modifies  the "Witness"
func (svc *WitnessService) UpdateWitness(id string, obj *nimbleos.Witness) (*nimbleos.Witness, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	witnessResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return witnessResp, nil
}

// GetWitnessById - method returns a pointer to "Witness"
func (svc *WitnessService) GetWitnessById(id string) (*nimbleos.Witness, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	witnessResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return witnessResp, nil
}

// DeleteWitness - deletes the "Witness"
func (svc *WitnessService) DeleteWitness(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
