// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsAuditCategoryDataProtection *NsAuditCategory
var NsAuditCategorySystemConfiguration *NsAuditCategory
var NsAuditCategorySoftwareUpdate *NsAuditCategory
var NsAuditCategoryUserAccess *NsAuditCategory
var NsAuditCategoryDataAccess *NsAuditCategory
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
