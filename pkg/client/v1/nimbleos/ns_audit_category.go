// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsAuditCategory Enum.

type NsAuditCategory string

const (
	cNsAuditCategoryDataProtection      NsAuditCategory = "data_protection"
	cNsAuditCategorySystemConfiguration NsAuditCategory = "system_configuration"
	cNsAuditCategorySoftwareUpdate      NsAuditCategory = "software_update"
	cNsAuditCategoryUserAccess          NsAuditCategory = "user_access"
	cNsAuditCategoryDataAccess          NsAuditCategory = "data_access"
	cNsAuditCategoryDataProvisioning    NsAuditCategory = "data_provisioning"
)

var pNsAuditCategoryDataProtection NsAuditCategory
var pNsAuditCategorySystemConfiguration NsAuditCategory
var pNsAuditCategorySoftwareUpdate NsAuditCategory
var pNsAuditCategoryUserAccess NsAuditCategory
var pNsAuditCategoryDataAccess NsAuditCategory
var pNsAuditCategoryDataProvisioning NsAuditCategory

// NsAuditCategoryDataProtection enum exports
var NsAuditCategoryDataProtection *NsAuditCategory

// NsAuditCategorySystemConfiguration enum exports
var NsAuditCategorySystemConfiguration *NsAuditCategory

// NsAuditCategorySoftwareUpdate enum exports
var NsAuditCategorySoftwareUpdate *NsAuditCategory

// NsAuditCategoryUserAccess enum exports
var NsAuditCategoryUserAccess *NsAuditCategory

// NsAuditCategoryDataAccess enum exports
var NsAuditCategoryDataAccess *NsAuditCategory

// NsAuditCategoryDataProvisioning enum exports
var NsAuditCategoryDataProvisioning *NsAuditCategory

func init() {
	pNsAuditCategoryDataProtection = cNsAuditCategoryDataProtection
	NsAuditCategoryDataProtection = &pNsAuditCategoryDataProtection

	pNsAuditCategorySystemConfiguration = cNsAuditCategorySystemConfiguration
	NsAuditCategorySystemConfiguration = &pNsAuditCategorySystemConfiguration

	pNsAuditCategorySoftwareUpdate = cNsAuditCategorySoftwareUpdate
	NsAuditCategorySoftwareUpdate = &pNsAuditCategorySoftwareUpdate

	pNsAuditCategoryUserAccess = cNsAuditCategoryUserAccess
	NsAuditCategoryUserAccess = &pNsAuditCategoryUserAccess

	pNsAuditCategoryDataAccess = cNsAuditCategoryDataAccess
	NsAuditCategoryDataAccess = &pNsAuditCategoryDataAccess

	pNsAuditCategoryDataProvisioning = cNsAuditCategoryDataProvisioning
	NsAuditCategoryDataProvisioning = &pNsAuditCategoryDataProvisioning

}
