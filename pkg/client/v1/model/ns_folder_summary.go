/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsFolderSummary - Select fields containing folder info.
// Export NsFolderSummaryFields for advance operations like search filter etc.
var NsFolderSummaryFields *NsFolderSummary

func init(){
	IDfield:= "id"
	Fqnfield:= "fqn"
		
	NsFolderSummaryFields= &NsFolderSummary{
		ID: &IDfield,
		Fqn: &Fqnfield,
		
	}
}

type NsFolderSummary struct {
   
    // ID of folder.
    
 	ID *string `json:"id,omitempty"`
   
    // Fully qualified name of folder.
    
 	Fqn *string `json:"fqn,omitempty"`
}
