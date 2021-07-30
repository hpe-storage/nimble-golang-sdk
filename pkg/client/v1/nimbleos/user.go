// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export UserFields provides field names to use in filter parameters, for example.
var UserFields *UserFieldHandles

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldSearchName := "search_name"
	fieldDescription := "description"
	fieldRoleId := "role_id"
	fieldRole := "role"
	fieldPassword := "password"
	fieldAuthPassword := "auth_password"
	fieldOtpType := "otp_type"
	fieldOtpReset := "otp_reset"
	fieldInactivityTimeout := "inactivity_timeout"
	fieldCreationTime := "creation_time"
	fieldLastModified := "last_modified"
	fieldFullName := "full_name"
	fieldEmailAddr := "email_addr"
	fieldTenantId := "tenant_id"
	fieldTenantKey := "tenant_key"
	fieldDisabled := "disabled"
	fieldAuthLock := "auth_lock"
	fieldLastLogin := "last_login"
	fieldLastLogout := "last_logout"
	fieldLoggedIn := "logged_in"

	UserFields = &UserFieldHandles{
		ID:                &fieldID,
		Name:              &fieldName,
		SearchName:        &fieldSearchName,
		Description:       &fieldDescription,
		RoleId:            &fieldRoleId,
		Role:              &fieldRole,
		Password:          &fieldPassword,
		AuthPassword:      &fieldAuthPassword,
		OtpType:           &fieldOtpType,
		OtpReset:          &fieldOtpReset,
		InactivityTimeout: &fieldInactivityTimeout,
		CreationTime:      &fieldCreationTime,
		LastModified:      &fieldLastModified,
		FullName:          &fieldFullName,
		EmailAddr:         &fieldEmailAddr,
		TenantId:          &fieldTenantId,
		TenantKey:         &fieldTenantKey,
		Disabled:          &fieldDisabled,
		AuthLock:          &fieldAuthLock,
		LastLogin:         &fieldLastLogin,
		LastLogout:        &fieldLastLogout,
		LoggedIn:          &fieldLoggedIn,
	}
}

// User - Represents users configured to manage the system.
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
	// OtpType - Type of One Time Password authentication in use.
	OtpType *NsOTPType `json:"otp_type,omitempty"`
	// OtpReset - When sent as true, this causes a reset of the One Time Password secret for the user.
	OtpReset *bool `json:"otp_reset,omitempty"`
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
	// TenantId - Identifier for the tenant.
	TenantId *string `json:"tenant_id,omitempty"`
	// TenantKey - Tenant secret key for encrypting the password.
	TenantKey *string `json:"tenant_key,omitempty"`
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

// UserFieldHandles provides a string representation for each AccessControlRecord field.
type UserFieldHandles struct {
	ID                *string
	Name              *string
	SearchName        *string
	Description       *string
	RoleId            *string
	Role              *string
	Password          *string
	AuthPassword      *string
	OtpType           *string
	OtpReset          *string
	InactivityTimeout *string
	CreationTime      *string
	LastModified      *string
	FullName          *string
	EmailAddr         *string
	TenantId          *string
	TenantKey         *string
	Disabled          *string
	AuthLock          *string
	LastLogin         *string
	LastLogout        *string
	LoggedIn          *string
}
