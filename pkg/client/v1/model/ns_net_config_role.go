// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsNetConfigRoleBackup *NsNetConfigRole
var NsNetConfigRoleDraft *NsNetConfigRole
var NsNetConfigRoleActive *NsNetConfigRole

func init() {
	pNsNetConfigRoleBackup = cNsNetConfigRoleBackup
	NsNetConfigRoleBackup = &pNsNetConfigRoleBackup

	pNsNetConfigRoleDraft = cNsNetConfigRoleDraft
	NsNetConfigRoleDraft = &pNsNetConfigRoleDraft

	pNsNetConfigRoleActive = cNsNetConfigRoleActive
	NsNetConfigRoleActive = &pNsNetConfigRoleActive

}
