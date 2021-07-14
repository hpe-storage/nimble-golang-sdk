// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// Dna - Get list of DNAs.
// Export DnaFields for advance operations like search filter etc.
var DnaFields *Dna

func init(){
 IDfield:= "id"
 ArrayNamefield:= "array_name"
 DnaNamefield:= "dna_name"

 DnaFields= &Dna{
  ID:             &IDfield,
  ArrayName:      &ArrayNamefield,
  DnaName:        &DnaNamefield,
 }
}


type Dna struct {
 // ID - Identifier for the DNA.
  ID *string `json:"id,omitempty"`
 // ArrayName - Name of array holding the DNA.
  ArrayName *string `json:"array_name,omitempty"`
 // CreationTime - Creation time for the DNA.
    CreationTime *int64 `json:"creation_time,omitempty"`
 // FileSize - Size of the DNA in bytes.
    FileSize *int64 `json:"file_size,omitempty"`
 // DnaName - Name of the DNA, unique in the context of its controller.
  DnaName *string `json:"dna_name,omitempty"`
 // ControllerName - Name of the controller storing the DNA.
    ControllerName *NsControllerSide `json:"controller_name,omitempty"`
}
