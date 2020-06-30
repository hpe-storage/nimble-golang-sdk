// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ApplicationServer Service - An application server is an external agent that collaborates with an array to manage storage resources; for example, Volume Shadow Copy Service (VSS) or VMware.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// ApplicationServerService type 
type ApplicationServerService struct {
	objectSet *client.ApplicationServerObjectSet
}

// NewApplicationServerService - method to initialize "ApplicationServerService" 
func NewApplicationServerService(gs *NsGroupService) (*ApplicationServerService) {
	objectSet := gs.client.GetApplicationServerObjectSet()
	return &ApplicationServerService{objectSet: objectSet}
}

// GetApplicationServers - method returns a array of pointers of type "ApplicationServers"
func (svc *ApplicationServerService) GetApplicationServers(params *util.GetParams) ([]*model.ApplicationServer, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetApplicationServersWithFields - method returns a array of pointers of type "ApplicationServer" 
func (svc *ApplicationServerService) GetApplicationServersWithFields(fields []string) ([]*model.ApplicationServer, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateApplicationServer - method creates a "ApplicationServer"
func (svc *ApplicationServerService) CreateApplicationServer(obj *model.ApplicationServer) (*model.ApplicationServer, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditApplicationServer - method modifies  the "ApplicationServer" 
func (svc *ApplicationServerService) EditApplicationServer(id string, obj *model.ApplicationServer) (*model.ApplicationServer, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyApplicationServer - private method for more than one element check. 
func onlyApplicationServer(objs []*model.ApplicationServer) (*model.ApplicationServer, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one ApplicationServer found with the given filter")
	}

	return objs[0], nil
}

 
// GetApplicationServersByID - method returns associative a array of pointers of type "ApplicationServer", filter by Id
func (svc *ApplicationServerService) GetApplicationServersByID(pool *model.Pool, fields []string) (map[string]*model.ApplicationServer, error) {
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
	objs, err := svc.GetApplicationServers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ApplicationServer)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetApplicationServerById - method returns a pointer to "ApplicationServer"
func (svc *ApplicationServerService) GetApplicationServerById(id string) (*model.ApplicationServer, error) {
	return svc.objectSet.GetObject(id)
}

// GetApplicationServersByName - method returns a associative array of pointers of type "ApplicationServer", filter by name 
func (svc *ApplicationServerService) GetApplicationServersByName(pool *model.Pool, fields []string) (map[string]*model.ApplicationServer, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetApplicationServers(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.ApplicationServer)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetApplicationServerByName - method returns a pointer "ApplicationServer" 
func (svc *ApplicationServerService) GetApplicationServerByName(name string) (*model.ApplicationServer, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return onlyApplicationServer(objs)
}	

