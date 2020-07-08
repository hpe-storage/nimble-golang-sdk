// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsEventTargetType Enum.

type NsEventTargetType string

const (
	cNsEventTargetTypeAnon           NsEventTargetType = "anon"
	cNsEventTargetTypeController     NsEventTargetType = "controller"
	cNsEventTargetTypeTest           NsEventTargetType = "test"
	cNsEventTargetTypeProtectionSet  NsEventTargetType = "protection_set"
	cNsEventTargetTypePool           NsEventTargetType = "pool"
	cNsEventTargetTypeNic            NsEventTargetType = "nic"
	cNsEventTargetTypeShelf          NsEventTargetType = "shelf"
	cNsEventTargetTypeVolume         NsEventTargetType = "volume"
	cNsEventTargetTypeDisk           NsEventTargetType = "disk"
	cNsEventTargetTypeFan            NsEventTargetType = "fan"
	cNsEventTargetTypeIscsi          NsEventTargetType = "iscsi"
	cNsEventTargetTypeNvram          NsEventTargetType = "nvram"
	cNsEventTargetTypePowerSupply    NsEventTargetType = "power_supply"
	cNsEventTargetTypePartner        NsEventTargetType = "partner"
	cNsEventTargetTypeArray          NsEventTargetType = "array"
	cNsEventTargetTypeService        NsEventTargetType = "service"
	cNsEventTargetTypeTemperature    NsEventTargetType = "temperature"
	cNsEventTargetTypeNtb            NsEventTargetType = "ntb"
	cNsEventTargetTypeFc             NsEventTargetType = "fc"
	cNsEventTargetTypeInitiatorGroup NsEventTargetType = "initiator_group"
	cNsEventTargetTypeRaid           NsEventTargetType = "raid"
	cNsEventTargetTypeGroup          NsEventTargetType = "group"
)

var pNsEventTargetTypeAnon NsEventTargetType
var pNsEventTargetTypeController NsEventTargetType
var pNsEventTargetTypeTest NsEventTargetType
var pNsEventTargetTypeProtectionSet NsEventTargetType
var pNsEventTargetTypePool NsEventTargetType
var pNsEventTargetTypeNic NsEventTargetType
var pNsEventTargetTypeShelf NsEventTargetType
var pNsEventTargetTypeVolume NsEventTargetType
var pNsEventTargetTypeDisk NsEventTargetType
var pNsEventTargetTypeFan NsEventTargetType
var pNsEventTargetTypeIscsi NsEventTargetType
var pNsEventTargetTypeNvram NsEventTargetType
var pNsEventTargetTypePowerSupply NsEventTargetType
var pNsEventTargetTypePartner NsEventTargetType
var pNsEventTargetTypeArray NsEventTargetType
var pNsEventTargetTypeService NsEventTargetType
var pNsEventTargetTypeTemperature NsEventTargetType
var pNsEventTargetTypeNtb NsEventTargetType
var pNsEventTargetTypeFc NsEventTargetType
var pNsEventTargetTypeInitiatorGroup NsEventTargetType
var pNsEventTargetTypeRaid NsEventTargetType
var pNsEventTargetTypeGroup NsEventTargetType

// Export
var NsEventTargetTypeAnon *NsEventTargetType
var NsEventTargetTypeController *NsEventTargetType
var NsEventTargetTypeTest *NsEventTargetType
var NsEventTargetTypeProtectionSet *NsEventTargetType
var NsEventTargetTypePool *NsEventTargetType
var NsEventTargetTypeNic *NsEventTargetType
var NsEventTargetTypeShelf *NsEventTargetType
var NsEventTargetTypeVolume *NsEventTargetType
var NsEventTargetTypeDisk *NsEventTargetType
var NsEventTargetTypeFan *NsEventTargetType
var NsEventTargetTypeIscsi *NsEventTargetType
var NsEventTargetTypeNvram *NsEventTargetType
var NsEventTargetTypePowerSupply *NsEventTargetType
var NsEventTargetTypePartner *NsEventTargetType
var NsEventTargetTypeArray *NsEventTargetType
var NsEventTargetTypeService *NsEventTargetType
var NsEventTargetTypeTemperature *NsEventTargetType
var NsEventTargetTypeNtb *NsEventTargetType
var NsEventTargetTypeFc *NsEventTargetType
var NsEventTargetTypeInitiatorGroup *NsEventTargetType
var NsEventTargetTypeRaid *NsEventTargetType
var NsEventTargetTypeGroup *NsEventTargetType

func init() {
	pNsEventTargetTypeAnon = cNsEventTargetTypeAnon
	NsEventTargetTypeAnon = &pNsEventTargetTypeAnon

	pNsEventTargetTypeController = cNsEventTargetTypeController
	NsEventTargetTypeController = &pNsEventTargetTypeController

	pNsEventTargetTypeTest = cNsEventTargetTypeTest
	NsEventTargetTypeTest = &pNsEventTargetTypeTest

	pNsEventTargetTypeProtectionSet = cNsEventTargetTypeProtectionSet
	NsEventTargetTypeProtectionSet = &pNsEventTargetTypeProtectionSet

	pNsEventTargetTypePool = cNsEventTargetTypePool
	NsEventTargetTypePool = &pNsEventTargetTypePool

	pNsEventTargetTypeNic = cNsEventTargetTypeNic
	NsEventTargetTypeNic = &pNsEventTargetTypeNic

	pNsEventTargetTypeShelf = cNsEventTargetTypeShelf
	NsEventTargetTypeShelf = &pNsEventTargetTypeShelf

	pNsEventTargetTypeVolume = cNsEventTargetTypeVolume
	NsEventTargetTypeVolume = &pNsEventTargetTypeVolume

	pNsEventTargetTypeDisk = cNsEventTargetTypeDisk
	NsEventTargetTypeDisk = &pNsEventTargetTypeDisk

	pNsEventTargetTypeFan = cNsEventTargetTypeFan
	NsEventTargetTypeFan = &pNsEventTargetTypeFan

	pNsEventTargetTypeIscsi = cNsEventTargetTypeIscsi
	NsEventTargetTypeIscsi = &pNsEventTargetTypeIscsi

	pNsEventTargetTypeNvram = cNsEventTargetTypeNvram
	NsEventTargetTypeNvram = &pNsEventTargetTypeNvram

	pNsEventTargetTypePowerSupply = cNsEventTargetTypePowerSupply
	NsEventTargetTypePowerSupply = &pNsEventTargetTypePowerSupply

	pNsEventTargetTypePartner = cNsEventTargetTypePartner
	NsEventTargetTypePartner = &pNsEventTargetTypePartner

	pNsEventTargetTypeArray = cNsEventTargetTypeArray
	NsEventTargetTypeArray = &pNsEventTargetTypeArray

	pNsEventTargetTypeService = cNsEventTargetTypeService
	NsEventTargetTypeService = &pNsEventTargetTypeService

	pNsEventTargetTypeTemperature = cNsEventTargetTypeTemperature
	NsEventTargetTypeTemperature = &pNsEventTargetTypeTemperature

	pNsEventTargetTypeNtb = cNsEventTargetTypeNtb
	NsEventTargetTypeNtb = &pNsEventTargetTypeNtb

	pNsEventTargetTypeFc = cNsEventTargetTypeFc
	NsEventTargetTypeFc = &pNsEventTargetTypeFc

	pNsEventTargetTypeInitiatorGroup = cNsEventTargetTypeInitiatorGroup
	NsEventTargetTypeInitiatorGroup = &pNsEventTargetTypeInitiatorGroup

	pNsEventTargetTypeRaid = cNsEventTargetTypeRaid
	NsEventTargetTypeRaid = &pNsEventTargetTypeRaid

	pNsEventTargetTypeGroup = cNsEventTargetTypeGroup
	NsEventTargetTypeGroup = &pNsEventTargetTypeGroup

}
