// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export ChapUserFields provides field names to use in filter parameters, for example.
var ChapUserFields *ChapUserFieldHandles

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldFullName := "full_name"
	fieldSearchName := "search_name"
	fieldDescription := "description"
	fieldPassword := "password"
	fieldInitiatorIqns := "initiator_iqns"
	fieldTenantId := "tenant_id"
	fieldCreationTime := "creation_time"
	fieldLastModified := "last_modified"
	fieldVolList := "vol_list"
	fieldVolCount := "vol_count"

	ChapUserFields = &ChapUserFieldHandles{
		ID:            &fieldID,
		Name:          &fieldName,
		FullName:      &fieldFullName,
		SearchName:    &fieldSearchName,
		Description:   &fieldDescription,
		Password:      &fieldPassword,
		InitiatorIqns: &fieldInitiatorIqns,
		TenantId:      &fieldTenantId,
		CreationTime:  &fieldCreationTime,
		LastModified:  &fieldLastModified,
		VolList:       &fieldVolList,
		VolCount:      &fieldVolCount,
	}
}

// ChapUser - Manage Challenge-Response Handshake Authentication Protocol (CHAP) user accounts. CHAP users are one method of access control for iSCSI initiators. Each CHAP user has a CHAP password, sometimes called a CHAP secret. The CHAP passwords must match on the array and on the iSCSI initiator in order for the array to authenicate the initiator and allow it access. The CHAP user information must exist on both the array and the iSCSI initiator. Target authentication gives security only for the specific iSCSI target.
type ChapUser struct {
	// ID - Identifier for the CHAP user.
	ID *string `json:"id,omitempty"`
	// Name - Name of CHAP user.
	Name *string `json:"name,omitempty"`
	// FullName - CHAP user's fully qualified name.
	FullName *string `json:"full_name,omitempty"`
	// SearchName - CHAP user name used for object search.
	SearchName *string `json:"search_name,omitempty"`
	// Description - Text description of CHAP user.
	Description *string `json:"description,omitempty"`
	// Password - CHAP secret.The CHAP secret should be between 12-16 characters and cannot contain spaces or most punctuation.
	Password *string `json:"password,omitempty"`
	// InitiatorIqns - List of iSCSI initiators. To be configured with this CHAP user for iSCSI Group Target CHAP authentication. This attribute cannot be modified at the same time with other attributes. If any specified initiator is already associated with another CHAP user, it will be replaced by this CHAP user for future CHAP authentication.
	InitiatorIqns []*NsISCSIIQN `json:"initiator_iqns,omitempty"`
	// TenantId - Identifier for the tenant.
	TenantId *string `json:"tenant_id,omitempty"`
	// CreationTime - Time when this CHAP user was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this CHAP user was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// VolList - List of volumes associated with this CHAP user.
	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
	// VolCount - Count of volumes associated with this CHAP user.
	VolCount *int64 `json:"vol_count,omitempty"`
}

// ChapUserFieldHandles provides a string representation for each AccessControlRecord field.
type ChapUserFieldHandles struct {
	ID            *string
	Name          *string
	FullName      *string
	SearchName    *string
	Description   *string
	Password      *string
	InitiatorIqns *string
	TenantId      *string
	CreationTime  *string
	LastModified  *string
	VolList       *string
	VolCount      *string
}
