// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSnapVolFields provides field names to use in filter parameters, for example.
var NsSnapVolFields *NsSnapVolFieldHandles

func init() {
	NsSnapVolFields = &NsSnapVolFieldHandles{
		VolId:           "vol_id",
		SnapName:        "snap_name",
		SnapDescription: "snap_description",
		Cookie:          "cookie",
		Online:          "online",
		Writable:        "writable",
		AppUuid:         "app_uuid",
		AgentType:       "agent_type",
		Metadata:        "metadata",
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
	VolId           string
	SnapName        string
	SnapDescription string
	Cookie          string
	Online          string
	Writable        string
	AppUuid         string
	AgentType       string
	Metadata        string
}
