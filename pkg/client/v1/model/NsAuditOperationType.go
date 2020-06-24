/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsAuditOperationType.</p>
 */

type NsAuditOperationType string 

const (
	NSAUDITOPERATIONTYPE_OTHER NsAuditOperationType = "other"
	NSAUDITOPERATIONTYPE_READ NsAuditOperationType = "read"
	NSAUDITOPERATIONTYPE_CREATE NsAuditOperationType = "create"
	NSAUDITOPERATIONTYPE_UPDATE NsAuditOperationType = "update"
	NSAUDITOPERATIONTYPE_DELETE NsAuditOperationType = "delete"

) 
