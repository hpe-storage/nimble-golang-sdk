// Copyright 2020 Hewlett Packard Enterprise Development LP

package param

import "fmt"

type SortOrder struct {
	Field     string
	Ascending bool
}

func sortOrderToString(sortOrder SortOrder) string {
	if sortOrder.Ascending {
		return sortOrder.Field
	}
	return fmt.Sprintf("-%s", sortOrder.Field)
}
