// Copyright 2020 Hewlett Packard Enterprise Development LP

package param

// NewBool : Lame golang way to initialize a *bool.
// This is needed because the json marshaller considers false as
// empty and omits it from the request.
func NewBool(b bool) *bool {
	return &b
}

// NewInt - type int pointer
func NewInt(i int) *int {
	return &i
}

// NewInt64 - type int64 pointer
func NewInt64(i int64) *int64 {
	return &i
}
