// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsNetConfigRole Enum.

type NsNetConfigRole string

const (
	cNsNetConfigRoleBackup NsNetConfigRole = "backup"
	cNsNetConfigRoleDraft  NsNetConfigRole = "draft"
	cNsNetConfigRoleActive NsNetConfigRole = "active"
)

var pNsNetConfigRoleBackup NsNetConfigRole
var pNsNetConfigRoleDraft NsNetConfigRole
var pNsNetConfigRoleActive NsNetConfigRole

// NsNetConfigRoleBackup enum exports
var NsNetConfigRoleBackup *NsNetConfigRole

// NsNetConfigRoleDraft enum exports
var NsNetConfigRoleDraft *NsNetConfigRole

// NsNetConfigRoleActive enum exports
var NsNetConfigRoleActive *NsNetConfigRole

func init() {
	pNsNetConfigRoleBackup = cNsNetConfigRoleBackup
	NsNetConfigRoleBackup = &pNsNetConfigRoleBackup

	pNsNetConfigRoleDraft = cNsNetConfigRoleDraft
	NsNetConfigRoleDraft = &pNsNetConfigRoleDraft

	pNsNetConfigRoleActive = cNsNetConfigRoleActive
	NsNetConfigRoleActive = &pNsNetConfigRoleActive

}
