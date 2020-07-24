// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// User - Represents users configured to manage the system.
// Export UserFields for advance operations like search filter etc.
var UserFields *User

func init() {
	IDfield := "id"
	Namefield := "name"
	SearchNamefield := "search_name"
	Descriptionfield := "description"
	RoleIdfield := "role_id"
	Passwordfield := "password"
	AuthPasswordfield := "auth_password"
	FullNamefield := "full_name"
	EmailAddrfield := "email_addr"

	UserFields = &User{
		ID:           &IDfield,
		Name:         &Namefield,
		SearchName:   &SearchNamefield,
		Description:  &Descriptionfield,
		RoleId:       &RoleIdfield,
		Password:     &Passwordfield,
		AuthPassword: &AuthPasswordfield,
		FullName:     &FullNamefield,
		EmailAddr:    &EmailAddrfield,
	}
}

type User struct {
	// ID - Identifier for the user.
	ID *string `json:"id,omitempty"`
	// Name - Name of the user.
	Name *string `json:"name,omitempty"`
	// SearchName - Name of the user used for object search.
	SearchName *string `json:"search_name,omitempty"`
	// Description - Description of the user.
	Description *string `json:"description,omitempty"`
	// RoleId - Identifier for the user's role.
	RoleId *string `json:"role_id,omitempty"`
	// Role - Role of the user.
	Role *NsUserRoles `json:"role,omitempty"`
	// Password - User's login password.
	Password *string `json:"password,omitempty"`
	// AuthPassword - Authorization password for changing password.
	AuthPassword *string `json:"auth_password,omitempty"`
	// InactivityTimeout - The amount of time that the user session is inactive before timing out. A value of 0 indicates that the timeout is taken from the group setting.
	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`
	// CreationTime - Time when this user was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this user was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// FullName - Fully qualified name of the user.
	FullName *string `json:"full_name,omitempty"`
	// EmailAddr - Email address of the user.
	EmailAddr *string `json:"email_addr,omitempty"`
	// Disabled - User is currently disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// AuthLock - User was locked due to failed logins.
	AuthLock *bool `json:"auth_lock,omitempty"`
	// LastLogin - Last login time.
	LastLogin *int64 `json:"last_login,omitempty"`
	// LastLogout - Last logout time.
	LastLogout *int64 `json:"last_logout,omitempty"`
	// LoggedIn - User is currently logged in.
	LoggedIn *bool `json:"logged_in,omitempty"`
}
