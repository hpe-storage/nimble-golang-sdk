// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsHostType - Host type attribute.

// Export NsHostTypeFields provides field names to use in filter parameters, for example.
var NsHostTypeFields *NsHostTypeStringFields

func init() {
	fieldInitiatorName := "initiator_name"
	fieldSourceInitiatorGroup := "source_initiator_group"
	fieldSourceHostType := "source_host_type"
	fieldDestinationInitiatorGroup := "destination_initiator_group"
	fieldDestinationHostType := "destination_host_type"

	NsHostTypeFields = &NsHostTypeStringFields{
		InitiatorName:             &fieldInitiatorName,
		SourceInitiatorGroup:      &fieldSourceInitiatorGroup,
		SourceHostType:            &fieldSourceHostType,
		DestinationInitiatorGroup: &fieldDestinationInitiatorGroup,
		DestinationHostType:       &fieldDestinationHostType,
	}
}

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

// Struct for NsHostTypeFields
type NsHostTypeStringFields struct {
	InitiatorName             *string
	SourceInitiatorGroup      *string
	SourceHostType            *string
	DestinationInitiatorGroup *string
	DestinationHostType       *string
}
