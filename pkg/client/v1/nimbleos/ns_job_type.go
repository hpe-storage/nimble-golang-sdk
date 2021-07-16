// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

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

// NsJobTypeBulkChild enum exports
var NsJobTypeBulkChild *NsJobType

// NsJobTypeBulkParent enum exports
var NsJobTypeBulkParent *NsJobType

// NsJobTypeSimple enum exports
var NsJobTypeSimple *NsJobType

func init() {
	pNsJobTypeBulkChild = cNsJobTypeBulkChild
	NsJobTypeBulkChild = &pNsJobTypeBulkChild

	pNsJobTypeBulkParent = cNsJobTypeBulkParent
	NsJobTypeBulkParent = &pNsJobTypeBulkParent

	pNsJobTypeSimple = cNsJobTypeSimple
	NsJobTypeSimple = &pNsJobTypeSimple

}
