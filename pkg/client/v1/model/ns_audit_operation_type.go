// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsAuditOperationType Enum.

type NsAuditOperationType string

const (
	cNsAuditOperationTypeOther  NsAuditOperationType = "other"
	cNsAuditOperationTypeRead   NsAuditOperationType = "read"
	cNsAuditOperationTypeCreate NsAuditOperationType = "create"
	cNsAuditOperationTypeUpdate NsAuditOperationType = "update"
	cNsAuditOperationTypeDelete NsAuditOperationType = "delete"
)

var pNsAuditOperationTypeOther NsAuditOperationType
var pNsAuditOperationTypeRead NsAuditOperationType
var pNsAuditOperationTypeCreate NsAuditOperationType
var pNsAuditOperationTypeUpdate NsAuditOperationType
var pNsAuditOperationTypeDelete NsAuditOperationType

// Export
var NsAuditOperationTypeOther *NsAuditOperationType
var NsAuditOperationTypeRead *NsAuditOperationType
var NsAuditOperationTypeCreate *NsAuditOperationType
var NsAuditOperationTypeUpdate *NsAuditOperationType
var NsAuditOperationTypeDelete *NsAuditOperationType

func init() {
	pNsAuditOperationTypeOther = cNsAuditOperationTypeOther
	NsAuditOperationTypeOther = &pNsAuditOperationTypeOther

	pNsAuditOperationTypeRead = cNsAuditOperationTypeRead
	NsAuditOperationTypeRead = &pNsAuditOperationTypeRead

	pNsAuditOperationTypeCreate = cNsAuditOperationTypeCreate
	NsAuditOperationTypeCreate = &pNsAuditOperationTypeCreate

	pNsAuditOperationTypeUpdate = cNsAuditOperationTypeUpdate
	NsAuditOperationTypeUpdate = &pNsAuditOperationTypeUpdate

	pNsAuditOperationTypeDelete = cNsAuditOperationTypeDelete
	NsAuditOperationTypeDelete = &pNsAuditOperationTypeDelete

}
