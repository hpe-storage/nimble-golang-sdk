// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// LdapDomainFields provides field names to use in filter parameters, for example.
var LdapDomainFields *LdapDomainFieldHandles

func init() {
	LdapDomainFields = &LdapDomainFieldHandles{
		ID:                  "id",
		DomainName:          "domain_name",
		DomainDescription:   "domain_description",
		DomainEnabled:       "domain_enabled",
		ServerUriList:       "server_uri_list",
		BindUser:            "bind_user",
		BindPassword:        "bind_password",
		BaseDn:              "base_dn",
		UserSearchFilter:    "user_search_filter",
		UserSearchBaseList:  "user_search_base_list",
		GroupSearchFilter:   "group_search_filter",
		GroupSearchBaseList: "group_search_base_list",
		SchemaType:          "schema_type",
	}
}

// LdapDomain - Manages the storage array's membership with LDAP servers.
type LdapDomain struct {
	// ID - Identifier for the LDAP Domain.
	ID *string `json:"id,omitempty"`
	// DomainName - Domain name.
	DomainName *string `json:"domain_name,omitempty"`
	// DomainDescription - Description of the domain.
	DomainDescription *string `json:"domain_description,omitempty"`
	// DomainEnabled - Indicates whether the LDAP domain is currently active or not.
	DomainEnabled *bool `json:"domain_enabled,omitempty"`
	// ServerUriList - A set of up to 3 LDAP URIs.
	ServerUriList []*string `json:"server_uri_list,omitempty"`
	// BindUser - Full Distinguished Name of LDAP admin user. If empty, attempt to bind anonymously.
	BindUser *string `json:"bind_user,omitempty"`
	// BindPassword - Password for the Full Distinguished Name of LDAP admin user.  This parameter is mandatory if the bind_user is given.
	BindPassword *string `json:"bind_password,omitempty"`
	// BaseDn - The Distinguished Name of the base object from which to start all searches.
	BaseDn *string `json:"base_dn,omitempty"`
	// UserSearchFilter - Limit the results returned based on specific search criteria.
	UserSearchFilter *string `json:"user_search_filter,omitempty"`
	// UserSearchBaseList - A set of upto 10 Relative Distinguished Names, relative to the Base DN, from which to search for User objects.
	UserSearchBaseList []*string `json:"user_search_base_list,omitempty"`
	// GroupSearchFilter - Limit the results returned based on specific search criteria.
	GroupSearchFilter *string `json:"group_search_filter,omitempty"`
	// GroupSearchBaseList - A set of upto 10 Relative Distinguished Names, relative to the Base DN, from which to search for Group objects.
	GroupSearchBaseList []*string `json:"group_search_base_list,omitempty"`
	// SchemaType - Enum values are OpenLDAP or AD.
	SchemaType *NsLdapSchemaType `json:"schema_type,omitempty"`
}

// LdapDomainFieldHandles provides a string representation for each LdapDomain field.
type LdapDomainFieldHandles struct {
	ID                  string
	DomainName          string
	DomainDescription   string
	DomainEnabled       string
	ServerUriList       string
	BindUser            string
	BindPassword        string
	BaseDn              string
	UserSearchFilter    string
	UserSearchBaseList  string
	GroupSearchFilter   string
	GroupSearchBaseList string
	SchemaType          string
}
