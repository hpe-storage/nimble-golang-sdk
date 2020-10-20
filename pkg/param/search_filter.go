// Copyright 2020 Hewlett Packard Enterprise Development LP

package param

import (
	"fmt"
	"strings"
)

type SearchFilter struct {
	Constructor           *string         `json:"_constructor,omitempty"`
	FieldName             *string         `json:"fieldName,omitempty"`
	Operator              string          `json:"operator"`
	Criteria              []*SearchFilter `json:"criteria,omitempty"`
	Value                 interface{}     `json:"value,omitempty"`
	forceAdvancedCriteria bool
	simpleMap             map[string]string
	isMetadata            bool
}

type LogicalOperator int

const (
	AND LogicalOperator = iota
	OR
)

func (e LogicalOperator) String() string {
	switch e {
	case AND:
		return "and"
	case OR:
		return "or"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

const (
	RESERVED_URL_CHARACTERS = "<> : = &"
)

type SingleOperator int

const (
	// keep in parity with gmd/rest/gm_rest_criteria.h operatorMap (line 121)
	// except AND and OR belong in LogicalOperator
	ICONTAINS SingleOperator = iota
	CONTAINS
	IN_SET
	NOT_IN_SET
	IEQUALS
	EQUALS
	NOT_EQUAL
	INOT_EQUAL
	BETWEEN_INCLUSIVE
	BETWEEN
	GREATER_OR_EQUAL
	GREATER_THAN
	LESS_OR_EQUAL
	LESS_THAN
	REGEXP
	ASSIGN
)

func (e SingleOperator) String() string {
	switch e {
	case ICONTAINS:
		return "iContains"
	case CONTAINS:
		return "contains"
	case IN_SET:
		return "inSet"
	case NOT_IN_SET:
		return "notInSet"
	case IEQUALS:
		return "iEquals"
	case EQUALS:
		return "equals"
	case NOT_EQUAL:
		return "notEqual"
	case INOT_EQUAL:
		return "iNotEqual"
	case BETWEEN_INCLUSIVE:
		return "betweenInclusive"
	case BETWEEN:
		return "between"
	case GREATER_OR_EQUAL:
		return "greaterOrEqual"
	case GREATER_THAN:
		return "greaterThan"
	case LESS_OR_EQUAL:
		return "lessOrEqual"
	case LESS_THAN:
		return "lessThan"
	case REGEXP:
		return "regexp"
	case ASSIGN:
		return "assign"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

func (filter *SearchFilter) prepareForPost() {
	temp := "AdvancedCriteria"
	filter.Constructor = &temp
	if filter.Criteria == nil {
		criteria := make([]*SearchFilter, 1)
		criteria[0] = &SearchFilter{
			FieldName: filter.FieldName,
			Operator:  filter.Operator,
			Value:     filter.Value,
		}
		filter.Criteria = criteria
		filter.Operator = LogicalOperator(AND).String()
		filter.FieldName = nil
		filter.Value = nil
	}
	filter.simpleMap = filter.makeSimpleMap()

}
func (filter *SearchFilter) makeSimpleMap() map[string]string {
	if filter.forceAdvancedCriteria {
		return nil
	}
	if filter.Criteria == nil {
		// only the equals operator is supported for a simple map
		if !strings.EqualFold(SingleOperator(EQUALS).String(), filter.Operator) {
			return nil
		}
		valueStr := filter.Value.(string)
		// if the value contains a reserved character that's not valid for a URL, return null to force AdvancedCriteria
		if strings.ContainsAny(valueStr, RESERVED_URL_CHARACTERS) {
			return nil
		}

		simpleMap := make(map[string]string)
		simpleMap[*filter.FieldName] = valueStr
		return simpleMap
	} else {

		simpleMap := make(map[string]string)
		for _, filter := range filter.Criteria {
			submap := filter.makeSimpleMap()
			if submap == nil {
				return nil
			}
			for k, v := range submap {
				simpleMap[k] = v
			}
		}
		return simpleMap
	}
}

func (filter *SearchFilter) AsSimpleMap() (map[string]string, error) {
	if filter.Constructor == nil {
		return nil, fmt.Errorf("Should be called only on the top level filter, and only after prepareForPost")
	}
	return filter.simpleMap, nil
}

func (filter *SearchFilter) Init(fieldName string, operator SingleOperator, value string, isMetadata bool) {
	filter.FieldName = &fieldName
	filter.Operator = SingleOperator(operator).String()
	filter.Value = value
	filter.isMetadata = isMetadata
}
