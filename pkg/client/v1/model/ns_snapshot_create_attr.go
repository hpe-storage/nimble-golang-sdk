/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSnapshotCreateAttr - Select fields containing volume info.
// Export NsSnapshotCreateAttrFields for advance operations like search filter etc.
var NsSnapshotCreateAttrFields *NsSnapshotCreateAttr

func init(){
	VolIDfield:= "vol_id"
	Namefield:= "name"
	Descriptionfield:= "description"
		
	NsSnapshotCreateAttrFields= &NsSnapshotCreateAttr{
		VolID: &VolIDfield,
		Name: &Namefield,
		Description: &Descriptionfield,
		
	}
}

type NsSnapshotCreateAttr struct {
   
    // ID of volume.
    
 	VolID *string `json:"vol_id,omitempty"`
   
    // Snapshot name.
    
 	Name *string `json:"name,omitempty"`
   
    // Snapshot description.
    
 	Description *string `json:"description,omitempty"`
   
    // Snapshot is online.
    
 	Online *bool `json:"online,omitempty"`
   
    // Snapshot is writable.
    
 	Writable *bool `json:"writable,omitempty"`
   
    // External management agent type.
    
   	AgentType *NsAgentType `json:"agent_type,omitempty"`
   
    // Key-value pairs that augment a snapshot's attributes.
    
   	Metadata []*NsKeyValue `json:"metadata,omitempty"`
}
