// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsJobType Enum.

type NsJobType string

const (
	cNsJobTypeBulkChild  NsJobType = "bulk_child"
	cNsJobTypeBulkParent NsJobType = "bulk_parent"
	cNsJobTypeSimple     NsJobType = "simple"
)

var pNsJobTypeBulkChild NsJobType
var pNsJobTypeBulkParent NsJobType
var pNsJobTypeSimple NsJobType

// Export
var NsJobTypeBulkChild *NsJobType
var NsJobTypeBulkParent *NsJobType
var NsJobTypeSimple *NsJobType

func init() {
	pNsJobTypeBulkChild = cNsJobTypeBulkChild
	NsJobTypeBulkChild = &pNsJobTypeBulkChild

	pNsJobTypeBulkParent = cNsJobTypeBulkParent
	NsJobTypeBulkParent = &pNsJobTypeBulkParent

	pNsJobTypeSimple = cNsJobTypeSimple
	NsJobTypeSimple = &pNsJobTypeSimple

}
