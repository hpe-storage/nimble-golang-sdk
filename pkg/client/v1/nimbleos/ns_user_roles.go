// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

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

// NsUserRolesAdministrator enum exports
var NsUserRolesAdministrator *NsUserRoles

// NsUserRolesGuest enum exports
var NsUserRolesGuest *NsUserRoles

// NsUserRolesOperator enum exports
var NsUserRolesOperator *NsUserRoles

// NsUserRolesPoweruser enum exports
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
