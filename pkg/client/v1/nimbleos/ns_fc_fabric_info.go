// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcFabricInfo - Fibre Channel fabric information.
// Export NsFcFabricInfoFields for advance operations like search filter etc.
var NsFcFabricInfoFields *NsFcFabricInfo

func init() {
	FabricSwitchNamefield := "fabric_switch_name"
	FabricSwitchPortfield := "fabric_switch_port"
	FabricSwitchWwnnfield := "fabric_switch_wwnn"
	FabricSwitchWwpnfield := "fabric_switch_wwpn"
	FabricWwnfield := "fabric_wwn"
	FcIdfield := "fc_id"

	NsFcFabricInfoFields = &NsFcFabricInfo{
		FabricSwitchName: &FabricSwitchNamefield,
		FabricSwitchPort: &FabricSwitchPortfield,
		FabricSwitchWwnn: &FabricSwitchWwnnfield,
		FabricSwitchWwpn: &FabricSwitchWwpnfield,
		FabricWwn:        &FabricWwnfield,
		FcId:             &FcIdfield,
	}
}

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
