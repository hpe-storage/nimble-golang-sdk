// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Stat Service - Access generic stats interface via REST for internal testing.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// StatService type 
type StatService struct {
	objectSet *client.StatObjectSet
}

// NewStatService - method to initialize "StatService" 
func NewStatService(gs *NsGroupService) (*StatService) {
	objectSet := gs.client.GetStatObjectSet()
	return &StatService{objectSet: objectSet}
}

// GetStats - method returns a array of pointers of type "Stats"
func (svc *StatService) GetStats(params *util.GetParams) ([]*model.Stat, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetStatsWithFields - method returns a array of pointers of type "Stat" 
func (svc *StatService) GetStatsWithFields(fields []string) ([]*model.Stat, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateStat - method creates a "Stat"
func (svc *StatService) CreateStat(obj *model.Stat) (*model.Stat, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditStat - method modifies  the "Stat" 
func (svc *StatService) EditStat(id string, obj *model.Stat) (*model.Stat, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyStat - private method for more than one element check. 
func onlyStat(objs []*model.Stat) (*model.Stat, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Stat found with the given filter")
	}

	return objs[0], nil
}


