// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// ChapUser - Manage Challenge-Response Handshake Authentication Protocol (CHAP) user accounts. CHAP users are one method of access control for iSCSI initiators. Each CHAP user has a CHAP password, sometimes called a CHAP secret. The CHAP passwords must match on the array and on the iSCSI initiator in order for the array to authenicate the initiator and allow it access. The CHAP user information must exist on both the array and the iSCSI initiator. Target authentication gives security only for the specific iSCSI target.
// Export ChapUserFields for advance operations like search filter etc.
var ChapUserFields *ChapUser

func init(){
	IDfield:= "id"
	Namefield:= "name"
	FullNamefield:= "full_name"
	SearchNamefield:= "search_name"
	Descriptionfield:= "description"
	Passwordfield:= "password"
		
	ChapUserFields= &ChapUser{
		ID:            &IDfield,
		Name:          &Namefield,
		FullName:      &FullNamefield,
		SearchName:    &SearchNamefield,
		Description:   &Descriptionfield,
		Password:      &Passwordfield,
	}
}

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
	// CreationTime - Time when this CHAP user was created.
   	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this CHAP user was last modified.
   	LastModified *int64 `json:"last_modified,omitempty"`
	// VolList - List of volumes associated with this CHAP user.
   	VolList []*NsVolumeSummary `json:"vol_list,omitempty"`
	// VolCount - Count of volumes associated with this CHAP user.
   	VolCount *int64 `json:"vol_count,omitempty"`
}
