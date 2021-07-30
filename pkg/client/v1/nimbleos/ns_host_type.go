// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsHostTypeFields provides field names to use in filter parameters, for example.
var NsHostTypeFields *NsHostTypeFieldHandles

func init() {
	NsHostTypeFields = &NsHostTypeFieldHandles{
		InitiatorName:             "initiator_name",
		SourceInitiatorGroup:      "source_initiator_group",
		SourceHostType:            "source_host_type",
		DestinationInitiatorGroup: "destination_initiator_group",
		DestinationHostType:       "destination_host_type",
	}
}

// NsHostType - Host type attribute.
type NsHostType struct {
	// InitiatorName - The initiator for which the conflict exists.
	InitiatorName *string `json:"initiator_name,omitempty"`
	// SourceInitiatorGroup - Source initiator group.
	SourceInitiatorGroup []*string `json:"source_initiator_group,omitempty"`
	// SourceHostType - Source host type.
	SourceHostType *string `json:"source_host_type,omitempty"`
	// DestinationInitiatorGroup - Destination initiator group.
	DestinationInitiatorGroup []*string `json:"destination_initiator_group,omitempty"`
	// DestinationHostType - Destination Host type.
	DestinationHostType *string `json:"destination_host_type,omitempty"`
}

// NsHostTypeFieldHandles provides a string representation for each AccessControlRecord field.
type NsHostTypeFieldHandles struct {
	InitiatorName             string
	SourceInitiatorGroup      string
	SourceHostType            string
	DestinationInitiatorGroup string
	DestinationHostType       string
}
