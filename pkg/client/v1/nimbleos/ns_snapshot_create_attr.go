// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapshotCreateAttr - Select fields containing volume info.
// Export NsSnapshotCreateAttrFields for advance operations like search filter etc.
var NsSnapshotCreateAttrFields *NsSnapshotCreateAttr

func init() {
	VolIdfield := "vol_id"
	Namefield := "name"
	Descriptionfield := "description"

	NsSnapshotCreateAttrFields = &NsSnapshotCreateAttr{
		VolId:       &VolIdfield,
		Name:        &Namefield,
		Description: &Descriptionfield,
	}
}

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
