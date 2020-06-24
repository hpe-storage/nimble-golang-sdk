// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsAuditCategory Enum.</p>
 */

type NsAuditCategory string 

const (
	NSAUDITCATEGORY_DATA_PROTECTION NsAuditCategory = "data_protection"
	NSAUDITCATEGORY_SYSTEM_CONFIGURATION NsAuditCategory = "system_configuration"
	NSAUDITCATEGORY_SOFTWARE_UPDATE NsAuditCategory = "software_update"
	NSAUDITCATEGORY_USER_ACCESS NsAuditCategory = "user_access"
	NSAUDITCATEGORY_DATA_ACCESS NsAuditCategory = "data_access"
	NSAUDITCATEGORY_DATA_PROVISIONING NsAuditCategory = "data_provisioning"

) 
