// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapshotCreateAttrFields provides field names to use in filter parameters, for example.
var NsSnapshotCreateAttrFields *NsSnapshotCreateAttrFieldHandles

func init() {
	NsSnapshotCreateAttrFields = &NsSnapshotCreateAttrFieldHandles{
		VolId:       "vol_id",
		Name:        "name",
		Description: "description",
		Online:      "online",
		Writable:    "writable",
		AgentType:   "agent_type",
		Metadata:    "metadata",
	}
}

// NsSnapshotCreateAttr - Select fields containing volume info.
type NsSnapshotCreateAttr struct {
	// VolId - ID of volume.
	VolId *string `json:"vol_id,omitempty"`
	// Name - Snapshot name.
	Name *string `json:"name,omitempty"`
	// Description - Snapshot description.
	Description *string `json:"description,omitempty"`
	// Online - Snapshot is online.
	Online *bool `json:"online,omitempty"`
	// Writable - Snapshot is writable.
	Writable *bool `json:"writable,omitempty"`
	// AgentType - External management agent type.
	AgentType *NsAgentType `json:"agent_type,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
}

// NsSnapshotCreateAttrFieldHandles provides a string representation for each AccessControlRecord field.
type NsSnapshotCreateAttrFieldHandles struct {
	VolId       string
	Name        string
	Description string
	Online      string
	Writable    string
	AgentType   string
	Metadata    string
}
