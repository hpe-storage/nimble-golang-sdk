// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export ActiveDirectoryMembershipFields provides field names to use in filter parameters, for example.
var ActiveDirectoryMembershipFields *ActiveDirectoryMembershipFieldHandles

func init() {
	fieldID := "id"
	fieldDescription := "description"
	fieldName := "name"
	fieldNetbios := "netbios"
	fieldServerList := "server_list"
	fieldComputerName := "computer_name"
	fieldOrganizationalUnit := "organizational_unit"
	fieldUser := "user"
	fieldPassword := "password"
	fieldEnabled := "enabled"

	ActiveDirectoryMembershipFields = &ActiveDirectoryMembershipFieldHandles{
		ID:                 &fieldID,
		Description:        &fieldDescription,
		Name:               &fieldName,
		Netbios:            &fieldNetbios,
		ServerList:         &fieldServerList,
		ComputerName:       &fieldComputerName,
		OrganizationalUnit: &fieldOrganizationalUnit,
		User:               &fieldUser,
		Password:           &fieldPassword,
		Enabled:            &fieldEnabled,
	}
}

// ActiveDirectoryMembership - Manages the storage array's membership with the Active Directory.
type ActiveDirectoryMembership struct {
	// ID - Identifier for the Active Directory Domain.
	ID *string `json:"id,omitempty"`
	// Description - Description for the Active Directory Domain.
	Description *string `json:"description,omitempty"`
	// Name - Identifier for the Active Directory domain.
	Name *string `json:"name,omitempty"`
	// Netbios - Netbios name for the Active Directory domain.
	Netbios *string `json:"netbios,omitempty"`
	// ServerList - List of IP addresses or names for the backup domain controller.
	ServerList *string `json:"server_list,omitempty"`
	// ComputerName - The name of the computer account in the domain controller.
	ComputerName *string `json:"computer_name,omitempty"`
	// OrganizationalUnit - The location for the computer account.
	OrganizationalUnit *string `json:"organizational_unit,omitempty"`
	// User - Name of the Activer Directory user with Administrator's privilege.
	User *string `json:"user,omitempty"`
	// Password - Password for the Active Directory user.
	Password *string `json:"password,omitempty"`
	// Enabled - Active Directory authentication is currently enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// ActiveDirectoryMembershipFieldHandles provides a string representation for each AccessControlRecord field.
type ActiveDirectoryMembershipFieldHandles struct {
	ID                 *string
	Description        *string
	Name               *string
	Netbios            *string
	ServerList         *string
	ComputerName       *string
	OrganizationalUnit *string
	User               *string
	Password           *string
	Enabled            *string
}
