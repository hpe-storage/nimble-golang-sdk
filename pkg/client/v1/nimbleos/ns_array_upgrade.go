// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArrayUpgrade - Array upgrade attributes.

// Export NsArrayUpgradeFields provides field names to use in filter parameters, for example.
var NsArrayUpgradeFields *NsArrayUpgradeStringFields

func init() {
	fieldType := "type"
	fieldState := "state"
	fieldStage := "stage"
	fieldCtrlrAPortList := "ctrlr_a_port_list"
	fieldCtrlrBPortList := "ctrlr_b_port_list"
	fieldMessages := "messages"
	fieldMetadata := "metadata"

	NsArrayUpgradeFields = &NsArrayUpgradeStringFields{
		Type:           &fieldType,
		State:          &fieldState,
		Stage:          &fieldStage,
		CtrlrAPortList: &fieldCtrlrAPortList,
		CtrlrBPortList: &fieldCtrlrBPortList,
		Messages:       &fieldMessages,
		Metadata:       &fieldMetadata,
	}
}

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

// Struct for NsArrayUpgradeFields
type NsArrayUpgradeStringFields struct {
	Type           *string
	State          *string
	Stage          *string
	CtrlrAPortList *string
	CtrlrBPortList *string
	Messages       *string
	Metadata       *string
}
