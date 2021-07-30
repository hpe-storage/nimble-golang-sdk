// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// ClientCredential - Credential that this device will trust.

// Export ClientCredentialFields provides field names to use in filter parameters, for example.
var ClientCredentialFields *ClientCredentialStringFields

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldClientId := "client_id"
	fieldSecret := "secret"
	fieldCreationTime := "creation_time"
	fieldLastModified := "last_modified"
	fieldDescription := "description"

	ClientCredentialFields = &ClientCredentialStringFields{
		ID:           &fieldID,
		Name:         &fieldName,
		ClientId:     &fieldClientId,
		Secret:       &fieldSecret,
		CreationTime: &fieldCreationTime,
		LastModified: &fieldLastModified,
		Description:  &fieldDescription,
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

// Struct for ClientCredentialFields
type ClientCredentialStringFields struct {
	ID           *string
	Name         *string
	ClientId     *string
	Secret       *string
	CreationTime *string
	LastModified *string
	Description  *string
}
