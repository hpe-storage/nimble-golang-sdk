// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// ActiveDirectoryMembership - Manages the storage array's membership with the Active Directory.
// Export ActiveDirectoryMembershipFields for advance operations like search filter etc.
var ActiveDirectoryMembershipFields *ActiveDirectoryMembership

func init() {

	ActiveDirectoryMembershipFields = &ActiveDirectoryMembership{
		ID:                 "id",
		Description:        "description",
		Name:               "name",
		Netbios:            "netbios",
		ServerList:         "server_list",
		ComputerName:       "computer_name",
		OrganizationalUnit: "organizational_unit",
		User:               "user",
		Password:           "password",
	}
}

type ActiveDirectoryMembership struct {
	// ID - Identifier for the Active Directory Domain.
	ID string `json:"id,omitempty"`
	// Description - Description for the Active Directory Domain.
	Description string `json:"description,omitempty"`
	// Name - Identifier for the Active Directory domain.
	Name string `json:"name,omitempty"`
	// Netbios - Netbios name for the Active Directory domain.
	Netbios string `json:"netbios,omitempty"`
	// ServerList - List of IP addresses or names for the backup domain controller.
	ServerList string `json:"server_list,omitempty"`
	// ComputerName - The name of the computer account in the domain controller.
	ComputerName string `json:"computer_name,omitempty"`
	// OrganizationalUnit - The location for the computer account.
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	// User - Name of the Activer Directory user with Administrator's privilege.
	User string `json:"user,omitempty"`
	// Password - Password for the Active Directory user.
	Password string `json:"password,omitempty"`
	// Enabled - Active Directory authentication is currently enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// sdk internal struct
type activeDirectoryMembership struct {
	ID                 *string `json:"id,omitempty"`
	Description        *string `json:"description,omitempty"`
	Name               *string `json:"name,omitempty"`
	Netbios            *string `json:"netbios,omitempty"`
	ServerList         *string `json:"server_list,omitempty"`
	ComputerName       *string `json:"computer_name,omitempty"`
	OrganizationalUnit *string `json:"organizational_unit,omitempty"`
	User               *string `json:"user,omitempty"`
	Password           *string `json:"password,omitempty"`
	Enabled            *bool   `json:"enabled,omitempty"`
}

// EncodeActiveDirectoryMembership - Transform ActiveDirectoryMembership to activeDirectoryMembership type
func EncodeActiveDirectoryMembership(request interface{}) (*activeDirectoryMembership, error) {
	reqActiveDirectoryMembership := request.(*ActiveDirectoryMembership)
	byte, err := json.Marshal(reqActiveDirectoryMembership)
	if err != nil {
		return nil, err
	}
	respActiveDirectoryMembershipPtr := &activeDirectoryMembership{}
	err = json.Unmarshal(byte, respActiveDirectoryMembershipPtr)
	return respActiveDirectoryMembershipPtr, err
}

// DecodeActiveDirectoryMembership Transform activeDirectoryMembership to ActiveDirectoryMembership type
func DecodeActiveDirectoryMembership(request interface{}) (*ActiveDirectoryMembership, error) {
	reqActiveDirectoryMembership := request.(*activeDirectoryMembership)
	byte, err := json.Marshal(reqActiveDirectoryMembership)
	if err != nil {
		return nil, err
	}
	respActiveDirectoryMembership := &ActiveDirectoryMembership{}
	err = json.Unmarshal(byte, respActiveDirectoryMembership)
	return respActiveDirectoryMembership, err
}
