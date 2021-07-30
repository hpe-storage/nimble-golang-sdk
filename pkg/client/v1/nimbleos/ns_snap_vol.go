// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapVol - Select fields containing volume info.
// Export NsSnapVolFields for advance operations like search filter etc.
var NsSnapVolFields *NsSnapVolStringFields

func init() {
	VolIdfield := "vol_id"
	SnapNamefield := "snap_name"
	SnapDescriptionfield := "snap_description"
	Cookiefield := "cookie"
	Onlinefield := "online"
	Writablefield := "writable"
	AppUuidfield := "app_uuid"
	AgentTypefield := "agent_type"
	Metadatafield := "metadata"

	NsSnapVolFields = &NsSnapVolStringFields{
		VolId:           &VolIdfield,
		SnapName:        &SnapNamefield,
		SnapDescription: &SnapDescriptionfield,
		Cookie:          &Cookiefield,
		Online:          &Onlinefield,
		Writable:        &Writablefield,
		AppUuid:         &AppUuidfield,
		AgentType:       &AgentTypefield,
		Metadata:        &Metadatafield,
	}
}

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

// Struct for NsSnapVolFields
type NsSnapVolStringFields struct {
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
