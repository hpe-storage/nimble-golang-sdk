// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Token - Manage user's session information.

// Export TokenFields provides field names to use in filter parameters, for example.
var TokenFields *TokenStringFields

func init() {
	fieldID := "id"
	fieldSessionToken := "session_token"
	fieldUsername := "username"
	fieldPassword := "password"
	fieldOtpCode := "otp_code"
	fieldAppName := "app_name"
	fieldSdkName := "sdk_name"
	fieldSourceIp := "source_ip"
	fieldCreationTime := "creation_time"
	fieldLastModified := "last_modified"
	fieldExpiryTime := "expiry_time"
	fieldServerUuid := "server_uuid"
	fieldGrantType := "grant_type"
	fieldAssertion := "assertion"

	TokenFields = &TokenStringFields{
		ID:           &fieldID,
		SessionToken: &fieldSessionToken,
		Username:     &fieldUsername,
		Password:     &fieldPassword,
		OtpCode:      &fieldOtpCode,
		AppName:      &fieldAppName,
		SdkName:      &fieldSdkName,
		SourceIp:     &fieldSourceIp,
		CreationTime: &fieldCreationTime,
		LastModified: &fieldLastModified,
		ExpiryTime:   &fieldExpiryTime,
		ServerUuid:   &fieldServerUuid,
		GrantType:    &fieldGrantType,
		Assertion:    &fieldAssertion,
	}
}

type Token struct {
	// ID - Object identifier for the session token.
	ID *string `json:"id,omitempty"`
	// SessionToken - Token used for authentication.
	SessionToken *string `json:"session_token,omitempty"`
	// Username - User name for the session.
	Username *string `json:"username,omitempty"`
	// Password - Password for the user. A password is required for creating a token.
	Password *string `json:"password,omitempty"`
	// OtpCode - One time password code.
	OtpCode *string `json:"otp_code,omitempty"`
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
	// GrantType - OAuth grant type, currently only support 'urn:ietf:params:oauth:grant-type:jwt-bearer'.
	GrantType *string `json:"grant_type,omitempty"`
	// Assertion - OAuth assertion, currently expecting a JWT token.
	Assertion *string `json:"assertion,omitempty"`
}

// Struct for TokenFields
type TokenStringFields struct {
	ID           *string
	SessionToken *string
	Username     *string
	Password     *string
	OtpCode      *string
	AppName      *string
	SdkName      *string
	SourceIp     *string
	CreationTime *string
	LastModified *string
	ExpiryTime   *string
	ServerUuid   *string
	GrantType    *string
	Assertion    *string
}
