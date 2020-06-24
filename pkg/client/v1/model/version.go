/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Version - Show the API version.
// Export VersionFields for advance operations like search filter etc.
var VersionFields *Version

func init(){
	Namefield:= "name"
	SoftwareVersionfield:= "software_version"
		
	VersionFields= &Version{
		Name: &Namefield,
		SoftwareVersion: &SoftwareVersionfield,
		
	}
}

type Version struct {
   
    // API version number.
    
 	Name *string `json:"name,omitempty"`
   
    // Software version of the group.
    
 	SoftwareVersion *string `json:"software_version,omitempty"`
}
