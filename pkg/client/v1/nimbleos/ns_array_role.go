// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsArrayRole Enum.

type NsArrayRole string

const (
 cNsArrayRoleLeader NsArrayRole = "leader"
 cNsArrayRoleNonMember NsArrayRole = "non_member"
 cNsArrayRoleInvalid NsArrayRole = "invalid"
 cNsArrayRoleBackupLeader NsArrayRole = "backup_leader"
 cNsArrayRoleMember NsArrayRole = "member"
 cNsArrayRoleFailed NsArrayRole = "failed"
)

var pNsArrayRoleLeader NsArrayRole
var pNsArrayRoleNonMember NsArrayRole
var pNsArrayRoleInvalid NsArrayRole
var pNsArrayRoleBackupLeader NsArrayRole
var pNsArrayRoleMember NsArrayRole
var pNsArrayRoleFailed NsArrayRole

// NsArrayRoleLeader enum exports
var NsArrayRoleLeader *NsArrayRole

// NsArrayRoleNonMember enum exports
var NsArrayRoleNonMember *NsArrayRole

// NsArrayRoleInvalid enum exports
var NsArrayRoleInvalid *NsArrayRole

// NsArrayRoleBackupLeader enum exports
var NsArrayRoleBackupLeader *NsArrayRole

// NsArrayRoleMember enum exports
var NsArrayRoleMember *NsArrayRole

// NsArrayRoleFailed enum exports
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

