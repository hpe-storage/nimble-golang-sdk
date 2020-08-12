// Copyright 2020 Hewlett Packard Enterprise Development LP

package param

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param/pagination"
	"strings"
)

type GetParams struct {
	Filter *SearchFilter
	Fields []string
	Page   *pagination.Page
	SortBy []SortOrder
}

func (params *GetParams) WithFields(fields []string) {
	params.Fields = fields
}

func (params *GetParams) WithSearchFilter(filter *SearchFilter) {
	params.Filter = filter
}

func (params *GetParams) WithSortBy(sortBy []SortOrder) {
	params.SortBy = sortBy
}

func (params *GetParams) BuildQueryParts() (map[string]string, error) {
	queryParts := make(map[string]string)
	if params.Page != nil && params.Page.StartRow != nil {
		queryParts["startRow"] = fmt.Sprintf("%d", *params.Page.StartRow)
	}

	if params.Page != nil && params.Page.EndRow != nil {
		queryParts["endRow"] = fmt.Sprintf("%d", *params.Page.EndRow)
	}

	if params.Fields != nil {
		if len(params.Fields) == 0 {
			return nil, fmt.Errorf("No fields specified")
		}
		queryParts["fields"] = strings.Join(params.Fields, ",")
	}

	if params.SortBy != nil {
		var sortByStrings []string
		for _, sortBy := range params.SortBy {
			sortByStrings = append(sortByStrings, sortOrderToString(sortBy))
		}
		queryParts["sortBy"] = strings.Join(sortByStrings, ",")
	}

	if params.Filter != nil {
		// prepare for post
		params.Filter.prepareForPost()
		filterMap, err := params.Filter.AsSimpleMap()
		if err != nil {
			return nil, err
		}

		if filterMap != nil {
			params.prepareUrlFilter(filterMap, queryParts)
		}
	}

	return queryParts, nil
}

func (params *GetParams) prepareUrlFilter(filterMap, queryParts map[string]string) {
	for k, v := range filterMap {
		queryParts[k] = v
	}
}

// FindField looks for given field in params. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func (params *GetParams) FindField(val string) (int, bool) {
	for i, item := range params.Fields {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
