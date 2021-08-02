// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcFabricInfoFields provides field names to use in filter parameters, for example.
var NsFcFabricInfoFields *NsFcFabricInfoFieldHandles

func init() {
	NsFcFabricInfoFields = &NsFcFabricInfoFieldHandles{
		FabricSwitchName: "fabric_switch_name",
		FabricSwitchPort: "fabric_switch_port",
		FabricSwitchWwnn: "fabric_switch_wwnn",
		FabricSwitchWwpn: "fabric_switch_wwpn",
		FabricWwn:        "fabric_wwn",
		FcId:             "fc_id",
		LoggedIn:         "logged_in",
	}
}

// NsFcFabricInfo - Fibre Channel fabric information.
type NsFcFabricInfo struct {
	// FabricSwitchName - Name of the Fibre Channel switch.
	FabricSwitchName *string `json:"fabric_switch_name,omitempty"`
	// FabricSwitchPort - Port on the Fibre Channel switch to which connection is established.
	FabricSwitchPort *string `json:"fabric_switch_port,omitempty"`
	// FabricSwitchWwnn - WWNN (World Wide Node Name) for the connected port on the fabric switch.
	FabricSwitchWwnn *string `json:"fabric_switch_wwnn,omitempty"`
	// FabricSwitchWwpn - WWPN (World Wide Port Name) for the connected port on the fabric switch.
	FabricSwitchWwpn *string `json:"fabric_switch_wwpn,omitempty"`
	// FabricWwn - WWNN(World Wide Node Name) for the Fabric Switch.
	FabricWwn *string `json:"fabric_wwn,omitempty"`
	// FcId - FCID assigned to the Fabric Channel fabric port.
	FcId *string `json:"fc_id,omitempty"`
	// LoggedIn - Login information for interface. True if interface has logged in to the Fibre Channel fabric, else false.
	LoggedIn *bool `json:"logged_in,omitempty"`
}

// NsFcFabricInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsFcFabricInfoFieldHandles struct {
	FabricSwitchName string
	FabricSwitchPort string
	FabricSwitchWwnn string
	FabricSwitchWwpn string
	FabricWwn        string
	FcId             string
	LoggedIn         string
}
