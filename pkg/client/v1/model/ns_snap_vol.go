/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSnapVol - Select fields containing volume info.
// Export NsSnapVolFields for advance operations like search filter etc.
var NsSnapVolFields *NsSnapVol

func init(){
	VolIDfield:= "vol_id"
	SnapNamefield:= "snap_name"
	SnapDescriptionfield:= "snap_description"
	Cookiefield:= "cookie"
	AppUuIDfield:= "app_uuid"
		
	NsSnapVolFields= &NsSnapVol{
		VolID: &VolIDfield,
		SnapName: &SnapNamefield,
		SnapDescription: &SnapDescriptionfield,
		Cookie: &Cookiefield,
		AppUuID: &AppUuIDfield,
		
	}
}

type NsSnapVol struct {
   
    // ID of volume.
    
 	VolID *string `json:"vol_id,omitempty"`
   
    // Snapshot name.
    
 	SnapName *string `json:"snap_name,omitempty"`
   
    // Snapshot description.
    
 	SnapDescription *string `json:"snap_description,omitempty"`
   
    // A cookie.
    
 	Cookie *string `json:"cookie,omitempty"`
   
    // Snapshot is online.
    
 	Online *bool `json:"online,omitempty"`
   
    // Snapshot is writable.
    
 	Writable *bool `json:"writable,omitempty"`
   
    // Application identifier of snapshots.
    
 	AppUuID *string `json:"app_uuid,omitempty"`
   
    // External management agent type.
    
   	AgentType *NsAgentType `json:"agent_type,omitempty"`
   
    // Key-value pairs that augment a snapshot's attributes.
    
   	Metadata []*NsKeyValue `json:"metadata,omitempty"`
}
