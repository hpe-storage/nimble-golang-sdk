// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// ClientCredential - Credential that this device will trust.
// Export ClientCredentialFields for advance operations like search filter etc.
var ClientCredentialFields *ClientCredential

func init(){
 IDfield:= "id"
 Namefield:= "name"
 ClientIdfield:= "client_id"
 Secretfield:= "secret"
 Descriptionfield:= "description"

 ClientCredentialFields= &ClientCredential{
  ID:           &IDfield,
  Name:         &Namefield,
  ClientId:     &ClientIdfield,
  Secret:       &Secretfield,
  Description:  &Descriptionfield,
 }
}


type ClientCredential struct {
 // ID - Identifier for the client credentials record.
  ID *string `json:"id,omitempty"`
 // Name - Client name.
  Name *string `json:"name,omitempty"`
 // ClientId - Uniqely identify this client, preferably uuid.
  ClientId *string `json:"client_id,omitempty"`
 // Secret - Client secret corresponding to this client id and name.
  Secret *string `json:"secret,omitempty"`
 // CreationTime - Time when this client credentials was created.
    CreationTime *int64 `json:"creation_time,omitempty"`
 // LastModified - Time when this client credentials was last modified.
    LastModified *int64 `json:"last_modified,omitempty"`
 // Description - Description of client.
  Description *string `json:"description,omitempty"`
}
