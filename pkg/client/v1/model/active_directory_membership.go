/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// ActiveDirectoryMembership - Manages the storage array's membership with the Active Directory.
// Export ActiveDirectoryMembershipFields for advance operations like search filter etc.
var ActiveDirectoryMembershipFields *ActiveDirectoryMembership

func init(){
	IDfield:= "id"
	Descriptionfield:= "description"
	Namefield:= "name"
	Netbiosfield:= "netbios"
	ServerListfield:= "server_list"
	ComputerNamefield:= "computer_name"
	OrganizationalUnitfield:= "organizational_unit"
	Userfield:= "user"
	Passwordfield:= "password"
		
	ActiveDirectoryMembershipFields= &ActiveDirectoryMembership{
		ID: &IDfield,
		Description: &Descriptionfield,
		Name: &Namefield,
		Netbios: &Netbiosfield,
		ServerList: &ServerListfield,
		ComputerName: &ComputerNamefield,
		OrganizationalUnit: &OrganizationalUnitfield,
		User: &Userfield,
		Password: &Passwordfield,
		
	}
}

type ActiveDirectoryMembership struct {
   
    // Identifier for the Active Directory Domain.
    
 	ID *string `json:"id,omitempty"`
   
    // Description for the Active Directory Domain.
    
 	Description *string `json:"description,omitempty"`
   
    // Identifier for the Active Directory domain.
    
 	Name *string `json:"name,omitempty"`
   
    // Netbios name for the Active Directory domain.
    
 	Netbios *string `json:"netbios,omitempty"`
   
    // List of IP addresses or names for the backup domain controller.
    
 	ServerList *string `json:"server_list,omitempty"`
   
    // The name of the computer account in the domain controller.
    
 	ComputerName *string `json:"computer_name,omitempty"`
   
    // The location for the computer account.
    
 	OrganizationalUnit *string `json:"organizational_unit,omitempty"`
   
    // Name of the Activer Directory user with Administrator's privilege.
    
 	User *string `json:"user,omitempty"`
   
    // Password for the Active Directory user.
    
 	Password *string `json:"password,omitempty"`
   
    // Active Directory authentication is currently enabled.
    
 	Enabled *bool `json:"enabled,omitempty"`
}
