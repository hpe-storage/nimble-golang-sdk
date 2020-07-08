// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsSpacePolicy Enum.

type NsSpacePolicy string

const (
	cNsSpacePolicyOffline     NsSpacePolicy = "offline"
	cNsSpacePolicyLoginOnly   NsSpacePolicy = "login_only"
	cNsSpacePolicyNonWritable NsSpacePolicy = "non_writable"
	cNsSpacePolicyReadOnly    NsSpacePolicy = "read_only"
	cNsSpacePolicyInvalid     NsSpacePolicy = "invalid"
)

var pNsSpacePolicyOffline NsSpacePolicy
var pNsSpacePolicyLoginOnly NsSpacePolicy
var pNsSpacePolicyNonWritable NsSpacePolicy
var pNsSpacePolicyReadOnly NsSpacePolicy
var pNsSpacePolicyInvalid NsSpacePolicy

// Export
var NsSpacePolicyOffline *NsSpacePolicy
var NsSpacePolicyLoginOnly *NsSpacePolicy
var NsSpacePolicyNonWritable *NsSpacePolicy
var NsSpacePolicyReadOnly *NsSpacePolicy
var NsSpacePolicyInvalid *NsSpacePolicy

func init() {
	pNsSpacePolicyOffline = cNsSpacePolicyOffline
	NsSpacePolicyOffline = &pNsSpacePolicyOffline

	pNsSpacePolicyLoginOnly = cNsSpacePolicyLoginOnly
	NsSpacePolicyLoginOnly = &pNsSpacePolicyLoginOnly

	pNsSpacePolicyNonWritable = cNsSpacePolicyNonWritable
	NsSpacePolicyNonWritable = &pNsSpacePolicyNonWritable

	pNsSpacePolicyReadOnly = cNsSpacePolicyReadOnly
	NsSpacePolicyReadOnly = &pNsSpacePolicyReadOnly

	pNsSpacePolicyInvalid = cNsSpacePolicyInvalid
	NsSpacePolicyInvalid = &pNsSpacePolicyInvalid

}
