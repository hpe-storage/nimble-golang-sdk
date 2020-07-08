// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsEventCategory Enum.

type NsEventCategory string

const (
	cNsEventCategoryReplication   NsEventCategory = "replication"
	cNsEventCategoryVolume        NsEventCategory = "volume"
	cNsEventCategorySecurity      NsEventCategory = "security"
	cNsEventCategoryTest          NsEventCategory = "test"
	cNsEventCategoryConfiguration NsEventCategory = "configuration"
	cNsEventCategoryService       NsEventCategory = "service"
	cNsEventCategoryUpdate        NsEventCategory = "update"
	cNsEventCategoryArrayUpgrade  NsEventCategory = "array_upgrade"
	cNsEventCategoryUnknown       NsEventCategory = "unknown"
	cNsEventCategoryHardware      NsEventCategory = "hardware"
)

var pNsEventCategoryReplication NsEventCategory
var pNsEventCategoryVolume NsEventCategory
var pNsEventCategorySecurity NsEventCategory
var pNsEventCategoryTest NsEventCategory
var pNsEventCategoryConfiguration NsEventCategory
var pNsEventCategoryService NsEventCategory
var pNsEventCategoryUpdate NsEventCategory
var pNsEventCategoryArrayUpgrade NsEventCategory
var pNsEventCategoryUnknown NsEventCategory
var pNsEventCategoryHardware NsEventCategory

// Export
var NsEventCategoryReplication *NsEventCategory
var NsEventCategoryVolume *NsEventCategory
var NsEventCategorySecurity *NsEventCategory
var NsEventCategoryTest *NsEventCategory
var NsEventCategoryConfiguration *NsEventCategory
var NsEventCategoryService *NsEventCategory
var NsEventCategoryUpdate *NsEventCategory
var NsEventCategoryArrayUpgrade *NsEventCategory
var NsEventCategoryUnknown *NsEventCategory
var NsEventCategoryHardware *NsEventCategory

func init() {
	pNsEventCategoryReplication = cNsEventCategoryReplication
	NsEventCategoryReplication = &pNsEventCategoryReplication

	pNsEventCategoryVolume = cNsEventCategoryVolume
	NsEventCategoryVolume = &pNsEventCategoryVolume

	pNsEventCategorySecurity = cNsEventCategorySecurity
	NsEventCategorySecurity = &pNsEventCategorySecurity

	pNsEventCategoryTest = cNsEventCategoryTest
	NsEventCategoryTest = &pNsEventCategoryTest

	pNsEventCategoryConfiguration = cNsEventCategoryConfiguration
	NsEventCategoryConfiguration = &pNsEventCategoryConfiguration

	pNsEventCategoryService = cNsEventCategoryService
	NsEventCategoryService = &pNsEventCategoryService

	pNsEventCategoryUpdate = cNsEventCategoryUpdate
	NsEventCategoryUpdate = &pNsEventCategoryUpdate

	pNsEventCategoryArrayUpgrade = cNsEventCategoryArrayUpgrade
	NsEventCategoryArrayUpgrade = &pNsEventCategoryArrayUpgrade

	pNsEventCategoryUnknown = cNsEventCategoryUnknown
	NsEventCategoryUnknown = &pNsEventCategoryUnknown

	pNsEventCategoryHardware = cNsEventCategoryHardware
	NsEventCategoryHardware = &pNsEventCategoryHardware

}
