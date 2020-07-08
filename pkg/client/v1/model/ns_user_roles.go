// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsUserRoles Enum.

type NsUserRoles string

const (
	cNsUserRolesAdministrator NsUserRoles = "administrator"
	cNsUserRolesGuest         NsUserRoles = "guest"
	cNsUserRolesOperator      NsUserRoles = "operator"
	cNsUserRolesPoweruser     NsUserRoles = "poweruser"
)

var pNsUserRolesAdministrator NsUserRoles
var pNsUserRolesGuest NsUserRoles
var pNsUserRolesOperator NsUserRoles
var pNsUserRolesPoweruser NsUserRoles

// Export
var NsUserRolesAdministrator *NsUserRoles
var NsUserRolesGuest *NsUserRoles
var NsUserRolesOperator *NsUserRoles
var NsUserRolesPoweruser *NsUserRoles

func init() {
	pNsUserRolesAdministrator = cNsUserRolesAdministrator
	NsUserRolesAdministrator = &pNsUserRolesAdministrator

	pNsUserRolesGuest = cNsUserRolesGuest
	NsUserRolesGuest = &pNsUserRolesGuest

	pNsUserRolesOperator = cNsUserRolesOperator
	NsUserRolesOperator = &pNsUserRolesOperator

	pNsUserRolesPoweruser = cNsUserRolesPoweruser
	NsUserRolesPoweruser = &pNsUserRolesPoweruser

}
