// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArrayUpgradeFields provides field names to use in filter parameters, for example.
var NsArrayUpgradeFields *NsArrayUpgradeFieldHandles

func init() {
	NsArrayUpgradeFields = &NsArrayUpgradeFieldHandles{
		Type:           "type",
		State:          "state",
		Stage:          "stage",
		CtrlrAPortList: "ctrlr_a_port_list",
		CtrlrBPortList: "ctrlr_b_port_list",
		Messages:       "messages",
		Metadata:       "metadata",
	}
}

// NsArrayUpgrade - Array upgrade attributes.
type NsArrayUpgrade struct {
	// Type - Array upgrade type.
	Type *NsArrayUpgradeType `json:"type,omitempty"`
	// State - Array upgrade state.
	State *NsArrayUpgradeState `json:"state,omitempty"`
	// Stage - Array upgrade stage.
	Stage *NsArrayUpgradeStage `json:"stage,omitempty"`
	// CtrlrAPortList - Comma-separated list of array upgrade port names for controller A.
	CtrlrAPortList *string `json:"ctrlr_a_port_list,omitempty"`
	// CtrlrBPortList - Comma-separated list of array upgrade port names for controller B.
	CtrlrBPortList *string `json:"ctrlr_b_port_list,omitempty"`
	// Messages - List of error or info messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
	// Metadata - Key-value pairs that augment an array upgrade attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
}

// NsArrayUpgradeFieldHandles provides a string representation for each NsArrayUpgrade field.
type NsArrayUpgradeFieldHandles struct {
	Type           string
	State          string
	Stage          string
	CtrlrAPortList string
	CtrlrBPortList string
	Messages       string
	Metadata       string
}
