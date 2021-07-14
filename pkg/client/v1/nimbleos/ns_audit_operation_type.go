// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsAuditOperationType Enum.

type NsAuditOperationType string

const (
 cNsAuditOperationTypeOther NsAuditOperationType = "other"
 cNsAuditOperationTypeRead NsAuditOperationType = "read"
 cNsAuditOperationTypeCreate NsAuditOperationType = "create"
 cNsAuditOperationTypeUpdate NsAuditOperationType = "update"
 cNsAuditOperationTypeDelete NsAuditOperationType = "delete"
)

var pNsAuditOperationTypeOther NsAuditOperationType
var pNsAuditOperationTypeRead NsAuditOperationType
var pNsAuditOperationTypeCreate NsAuditOperationType
var pNsAuditOperationTypeUpdate NsAuditOperationType
var pNsAuditOperationTypeDelete NsAuditOperationType

// NsAuditOperationTypeOther enum exports
var NsAuditOperationTypeOther *NsAuditOperationType

// NsAuditOperationTypeRead enum exports
var NsAuditOperationTypeRead *NsAuditOperationType

// NsAuditOperationTypeCreate enum exports
var NsAuditOperationTypeCreate *NsAuditOperationType

// NsAuditOperationTypeUpdate enum exports
var NsAuditOperationTypeUpdate *NsAuditOperationType

// NsAuditOperationTypeDelete enum exports
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

