// Copyright 2020 Hewlett Packard Enterprise Development LP

package util

// NewBool : Lame golang way to initialize a *bool.
// This is needed because the json marshaller considers false as
// empty and omits it from the request.
func NewBool(b bool) *bool {
	return &b
}
