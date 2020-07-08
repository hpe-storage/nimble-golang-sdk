// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsVolStatus Enum.

type NsVolStatus string

const (
	cNsVolStatusOffline     NsVolStatus = "offline"
	cNsVolStatusLoginOnly   NsVolStatus = "login_only"
	cNsVolStatusNonWritable NsVolStatus = "non_writable"
	cNsVolStatusReadOnly    NsVolStatus = "read_only"
	cNsVolStatusOnline      NsVolStatus = "online"
)

var pNsVolStatusOffline NsVolStatus
var pNsVolStatusLoginOnly NsVolStatus
var pNsVolStatusNonWritable NsVolStatus
var pNsVolStatusReadOnly NsVolStatus
var pNsVolStatusOnline NsVolStatus

// Export
var NsVolStatusOffline *NsVolStatus
var NsVolStatusLoginOnly *NsVolStatus
var NsVolStatusNonWritable *NsVolStatus
var NsVolStatusReadOnly *NsVolStatus
var NsVolStatusOnline *NsVolStatus

func init() {
	pNsVolStatusOffline = cNsVolStatusOffline
	NsVolStatusOffline = &pNsVolStatusOffline

	pNsVolStatusLoginOnly = cNsVolStatusLoginOnly
	NsVolStatusLoginOnly = &pNsVolStatusLoginOnly

	pNsVolStatusNonWritable = cNsVolStatusNonWritable
	NsVolStatusNonWritable = &pNsVolStatusNonWritable

	pNsVolStatusReadOnly = cNsVolStatusReadOnly
	NsVolStatusReadOnly = &pNsVolStatusReadOnly

	pNsVolStatusOnline = cNsVolStatusOnline
	NsVolStatusOnline = &pNsVolStatusOnline

}
