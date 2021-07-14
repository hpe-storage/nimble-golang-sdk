// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsAuditOperationTypeOrAll Enum.

type NsAuditOperationTypeOrAll string

const (
 cNsAuditOperationTypeOrAllAll NsAuditOperationTypeOrAll = "all"
 cNsAuditOperationTypeOrAllOther NsAuditOperationTypeOrAll = "other"
 cNsAuditOperationTypeOrAllRead NsAuditOperationTypeOrAll = "read"
 cNsAuditOperationTypeOrAllCreate NsAuditOperationTypeOrAll = "create"
 cNsAuditOperationTypeOrAllUpdate NsAuditOperationTypeOrAll = "update"
 cNsAuditOperationTypeOrAllDelete NsAuditOperationTypeOrAll = "delete"
)

var pNsAuditOperationTypeOrAllAll NsAuditOperationTypeOrAll
var pNsAuditOperationTypeOrAllOther NsAuditOperationTypeOrAll
var pNsAuditOperationTypeOrAllRead NsAuditOperationTypeOrAll
var pNsAuditOperationTypeOrAllCreate NsAuditOperationTypeOrAll
var pNsAuditOperationTypeOrAllUpdate NsAuditOperationTypeOrAll
var pNsAuditOperationTypeOrAllDelete NsAuditOperationTypeOrAll

// NsAuditOperationTypeOrAllAll enum exports
var NsAuditOperationTypeOrAllAll *NsAuditOperationTypeOrAll

// NsAuditOperationTypeOrAllOther enum exports
var NsAuditOperationTypeOrAllOther *NsAuditOperationTypeOrAll

// NsAuditOperationTypeOrAllRead enum exports
var NsAuditOperationTypeOrAllRead *NsAuditOperationTypeOrAll

// NsAuditOperationTypeOrAllCreate enum exports
var NsAuditOperationTypeOrAllCreate *NsAuditOperationTypeOrAll

// NsAuditOperationTypeOrAllUpdate enum exports
var NsAuditOperationTypeOrAllUpdate *NsAuditOperationTypeOrAll

// NsAuditOperationTypeOrAllDelete enum exports
var NsAuditOperationTypeOrAllDelete *NsAuditOperationTypeOrAll

func init() {
 pNsAuditOperationTypeOrAllAll = cNsAuditOperationTypeOrAllAll
 NsAuditOperationTypeOrAllAll = &pNsAuditOperationTypeOrAllAll

 pNsAuditOperationTypeOrAllOther = cNsAuditOperationTypeOrAllOther
 NsAuditOperationTypeOrAllOther = &pNsAuditOperationTypeOrAllOther

 pNsAuditOperationTypeOrAllRead = cNsAuditOperationTypeOrAllRead
 NsAuditOperationTypeOrAllRead = &pNsAuditOperationTypeOrAllRead

 pNsAuditOperationTypeOrAllCreate = cNsAuditOperationTypeOrAllCreate
 NsAuditOperationTypeOrAllCreate = &pNsAuditOperationTypeOrAllCreate

 pNsAuditOperationTypeOrAllUpdate = cNsAuditOperationTypeOrAllUpdate
 NsAuditOperationTypeOrAllUpdate = &pNsAuditOperationTypeOrAllUpdate

 pNsAuditOperationTypeOrAllDelete = cNsAuditOperationTypeOrAllDelete
 NsAuditOperationTypeOrAllDelete = &pNsAuditOperationTypeOrAllDelete

}

