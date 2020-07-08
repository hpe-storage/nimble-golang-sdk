// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsArrayRole Enum.

type NsArrayRole string

const (
	cNsArrayRoleLeader       NsArrayRole = "leader"
	cNsArrayRoleNonMember    NsArrayRole = "non_member"
	cNsArrayRoleInvalid      NsArrayRole = "invalid"
	cNsArrayRoleBackupLeader NsArrayRole = "backup_leader"
	cNsArrayRoleMember       NsArrayRole = "member"
	cNsArrayRoleFailed       NsArrayRole = "failed"
)

var pNsArrayRoleLeader NsArrayRole
var pNsArrayRoleNonMember NsArrayRole
var pNsArrayRoleInvalid NsArrayRole
var pNsArrayRoleBackupLeader NsArrayRole
var pNsArrayRoleMember NsArrayRole
var pNsArrayRoleFailed NsArrayRole

// Export
var NsArrayRoleLeader *NsArrayRole
var NsArrayRoleNonMember *NsArrayRole
var NsArrayRoleInvalid *NsArrayRole
var NsArrayRoleBackupLeader *NsArrayRole
var NsArrayRoleMember *NsArrayRole
var NsArrayRoleFailed *NsArrayRole

func init() {
	pNsArrayRoleLeader = cNsArrayRoleLeader
	NsArrayRoleLeader = &pNsArrayRoleLeader

	pNsArrayRoleNonMember = cNsArrayRoleNonMember
	NsArrayRoleNonMember = &pNsArrayRoleNonMember

	pNsArrayRoleInvalid = cNsArrayRoleInvalid
	NsArrayRoleInvalid = &pNsArrayRoleInvalid

	pNsArrayRoleBackupLeader = cNsArrayRoleBackupLeader
	NsArrayRoleBackupLeader = &pNsArrayRoleBackupLeader

	pNsArrayRoleMember = cNsArrayRoleMember
	NsArrayRoleMember = &pNsArrayRoleMember

	pNsArrayRoleFailed = cNsArrayRoleFailed
	NsArrayRoleFailed = &pNsArrayRoleFailed

}
