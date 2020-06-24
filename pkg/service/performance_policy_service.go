// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

import (
	"fmt"

	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

type PerformancePolicyService struct {
	objectSet *client.PerformancePolicyObjectSet
}

func NewPerformancePolicyService(gs *GroupService) (vs *PerformancePolicyService) {
	objectSet := gs.client.GetPerformancePolicyObjectSet()
	return &PerformancePolicyService{objectSet: objectSet}
}

func (vs *PerformancePolicyService) GetPerformancePolicyByName(name string) (*model.PerformancePolicy, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.PerformancePolicyFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	performancePolicies, err := vs.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return onlyPerformancePolicy(performancePolicies)
}

func onlyPerformancePolicy(performancePolicies []*model.PerformancePolicy) (*model.PerformancePolicy, error) {
	if len(performancePolicies) == 0 {
		return nil, nil
	}

	if len(performancePolicies) > 1 {
		return nil, fmt.Errorf("More than one performance policy found with the given filter")
	}

	return performancePolicies[0], nil
}
