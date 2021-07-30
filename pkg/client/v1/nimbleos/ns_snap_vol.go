// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSnapVolFields provides field names to use in filter parameters, for example.
var NsSnapVolFields *NsSnapVolFieldHandles

func init() {
	fieldVolId := "vol_id"
	fieldSnapName := "snap_name"
	fieldSnapDescription := "snap_description"
	fieldCookie := "cookie"
	fieldOnline := "online"
	fieldWritable := "writable"
	fieldAppUuid := "app_uuid"
	fieldAgentType := "agent_type"
	fieldMetadata := "metadata"

	NsSnapVolFields = &NsSnapVolFieldHandles{
		VolId:           &fieldVolId,
		SnapName:        &fieldSnapName,
		SnapDescription: &fieldSnapDescription,
		Cookie:          &fieldCookie,
		Online:          &fieldOnline,
		Writable:        &fieldWritable,
		AppUuid:         &fieldAppUuid,
		AgentType:       &fieldAgentType,
		Metadata:        &fieldMetadata,
	}
}

// NsSnapVol - Select fields containing volume info.
type NsSnapVol struct {
	// VolId - ID of volume.
	VolId *string `json:"vol_id,omitempty"`
	// SnapName - Snapshot name.
	SnapName *string `json:"snap_name,omitempty"`
	// SnapDescription - Snapshot description.
	SnapDescription *string `json:"snap_description,omitempty"`
	// Cookie - A cookie.
	Cookie *string `json:"cookie,omitempty"`
	// Online - Snapshot is online.
	Online *bool `json:"online,omitempty"`
	// Writable - Snapshot is writable.
	Writable *bool `json:"writable,omitempty"`
	// AppUuid - Application identifier of snapshots.
	AppUuid *string `json:"app_uuid,omitempty"`
	// AgentType - External management agent type.
	AgentType *NsAgentType `json:"agent_type,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
}

// NsSnapVolFieldHandles provides a string representation for each AccessControlRecord field.
type NsSnapVolFieldHandles struct {
	VolId           *string
	SnapName        *string
	SnapDescription *string
	Cookie          *string
	Online          *string
	Writable        *string
	AppUuid         *string
	AgentType       *string
	Metadata        *string
}
