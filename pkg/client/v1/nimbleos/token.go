// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Token - Manage user's session information.
// Export TokenFields for advance operations like search filter etc.
var TokenFields *Token

func init() {

	TokenFields = &Token{
		ID:           "id",
		SessionToken: "session_token",
		Username:     "username",
		Password:     "password",
		AppName:      "app_name",
		SdkName:      "sdk_name",
		SourceIp:     "source_ip",
		ServerUuid:   "server_uuid",
	}
}

type Token struct {
	// ID - Object identifier for the session token.
	ID string `json:"id,omitempty"`
	// SessionToken - Token used for authentication.
	SessionToken string `json:"session_token,omitempty"`
	// Username - User name for the session.
	Username string `json:"username,omitempty"`
	// Password - Password for the user. A password is required for creating a token.
	Password string `json:"password,omitempty"`
	// AppName - Application name.
	AppName string `json:"app_name,omitempty"`
	// SdkName - SDK name.
	SdkName string `json:"sdk_name,omitempty"`
	// SourceIp - IP address from which the session originates.
	SourceIp string `json:"source_ip,omitempty"`
	// CreationTime - Time when this token was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this token was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// ExpiryTime - Time when this token will expire.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// ServerUuid - Non mandatory 36 character uuid returned by the server. Currently only the witness REST server returns one.
	ServerUuid string `json:"server_uuid,omitempty"`
}

// sdk internal struct
type token struct {
	// ID - Object identifier for the session token.
	ID *string `json:"id,omitempty"`
	// SessionToken - Token used for authentication.
	SessionToken *string `json:"session_token,omitempty"`
	// Username - User name for the session.
	Username *string `json:"username,omitempty"`
	// Password - Password for the user. A password is required for creating a token.
	Password *string `json:"password,omitempty"`
	// AppName - Application name.
	AppName *string `json:"app_name,omitempty"`
	// SdkName - SDK name.
	SdkName *string `json:"sdk_name,omitempty"`
	// SourceIp - IP address from which the session originates.
	SourceIp *string `json:"source_ip,omitempty"`
	// CreationTime - Time when this token was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this token was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// ExpiryTime - Time when this token will expire.
	ExpiryTime *int64 `json:"expiry_time,omitempty"`
	// ServerUuid - Non mandatory 36 character uuid returned by the server. Currently only the witness REST server returns one.
	ServerUuid *string `json:"server_uuid,omitempty"`
}

// EncodeToken - Transform Token to token type
func EncodeToken(request interface{}) (*token, error) {
	reqToken := request.(*Token)
	byte, err := json.Marshal(reqToken)
	objPtr := &token{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeToken Transform token to Token type
func DecodeToken(request interface{}) (*Token, error) {
	reqToken := request.(*token)
	byte, err := json.Marshal(reqToken)
	obj := &Token{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
