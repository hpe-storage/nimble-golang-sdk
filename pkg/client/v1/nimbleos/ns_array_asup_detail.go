// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsArrayAsupDetailFields provides field names to use in filter parameters, for example.
var NsArrayAsupDetailFields *NsArrayAsupDetailFieldHandles

func init() {
	fieldArrayName := "array_name"
	fieldAsupValidate := "asup_validate"
	fieldNameResolution := "name_resolution"
	fieldPingfromMgmtip := "pingfrom_mgmtip"
	fieldPingfromCtrlra := "pingfrom_ctrlra"
	fieldPingfromCtrlrb := "pingfrom_ctrlrb"
	fieldHeartbeat := "heartbeat"
	fieldMessages := "messages"

	NsArrayAsupDetailFields = &NsArrayAsupDetailFieldHandles{
		ArrayName:      &fieldArrayName,
		AsupValidate:   &fieldAsupValidate,
		NameResolution: &fieldNameResolution,
		PingfromMgmtip: &fieldPingfromMgmtip,
		PingfromCtrlra: &fieldPingfromCtrlra,
		PingfromCtrlrb: &fieldPingfromCtrlrb,
		Heartbeat:      &fieldHeartbeat,
		Messages:       &fieldMessages,
	}
}

// NsArrayAsupDetail - Detailed array asup information.
type NsArrayAsupDetail struct {
	// ArrayName - Unique name of array.
	ArrayName *string `json:"array_name,omitempty"`
	// AsupValidate - This indicates whether the ASUP setting is enabled or disabled.
	AsupValidate *bool `json:"asup_validate,omitempty"`
	// NameResolution - Indicates whether DNS resolution succeeded.
	NameResolution *bool `json:"name_resolution,omitempty"`
	// PingfromMgmtip - Indicates whether connection from Management IP address to support server succeed.
	PingfromMgmtip *bool `json:"pingfrom_mgmtip,omitempty"`
	// PingfromCtrlra - Indicates whether connection from Controller-A to support server succeed.
	PingfromCtrlra *bool `json:"pingfrom_ctrlra,omitempty"`
	// PingfromCtrlrb - Indicates whether connection from Controller-B to support server succeed.
	PingfromCtrlrb *bool `json:"pingfrom_ctrlrb,omitempty"`
	// Heartbeat - Indicates whether heartbeat to support server succeed.
	Heartbeat *bool `json:"heartbeat,omitempty"`
	// Messages - A list of error messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// NsArrayAsupDetailFieldHandles provides a string representation for each AccessControlRecord field.
type NsArrayAsupDetailFieldHandles struct {
	ArrayName      *string
	AsupValidate   *string
	NameResolution *string
	PingfromMgmtip *string
	PingfromCtrlra *string
	PingfromCtrlrb *string
	Heartbeat      *string
	Messages       *string
}
