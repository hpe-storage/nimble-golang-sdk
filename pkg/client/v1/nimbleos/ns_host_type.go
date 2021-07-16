// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsHostType - Host type attribute.
// Export NsHostTypeFields for advance operations like search filter etc.
var NsHostTypeFields *NsHostType

func init() {
	InitiatorNamefield := "initiator_name"
	SourceHostTypefield := "source_host_type"
	DestinationHostTypefield := "destination_host_type"

	NsHostTypeFields = &NsHostType{
		InitiatorName:       &InitiatorNamefield,
		SourceHostType:      &SourceHostTypefield,
		DestinationHostType: &DestinationHostTypefield,
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
