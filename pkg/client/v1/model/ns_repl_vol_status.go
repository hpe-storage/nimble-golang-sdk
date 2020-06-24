/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsReplVolStatus - The replication status of a volume undergoing replication.
// Export NsReplVolStatusFields for advance operations like search filter etc.
var NsReplVolStatusFields *NsReplVolStatus

func init(){
	Namefield:= "name"
	SnapNamefield:= "snap_name"
		
	NsReplVolStatusFields= &NsReplVolStatus{
		Name: &Namefield,
		SnapName: &SnapNamefield,
		
	}
}

type NsReplVolStatus struct {
   
    // Name of the volume being replicated.
    
 	Name *string `json:"name,omitempty"`
   
    // Name of the snapshot being replicated.
    
 	SnapName *string `json:"snap_name,omitempty"`
   
    // Replication status of the volume.
    
   	Status *NsReplVolPartnerStatus `json:"status,omitempty"`
   
    // Internal replication status of the volume.
    
   	InternalStatus *NsReplVolPartnerStatus `json:"internal_status,omitempty"`
   
    // Transferred bytes.
    
   	ReplBytesDone *int64 `json:"repl_bytes_done,omitempty"`
   
    // Total number of bytes to be transferred.
    
   	ReplBytesTotal *int64 `json:"repl_bytes_total,omitempty"`
}
