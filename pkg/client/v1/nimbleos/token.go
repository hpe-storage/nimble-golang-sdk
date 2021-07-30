// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Token - Manage user's session information.
// Export TokenFields for advance operations like search filter etc.
var TokenFields *TokenStringFields

func init() {
	IDfield := "id"
	SessionTokenfield := "session_token"
	Usernamefield := "username"
	Passwordfield := "password"
	OtpCodefield := "otp_code"
	AppNamefield := "app_name"
	SdkNamefield := "sdk_name"
	SourceIpfield := "source_ip"
	CreationTimefield := "creation_time"
	LastModifiedfield := "last_modified"
	ExpiryTimefield := "expiry_time"
	ServerUuidfield := "server_uuid"
	GrantTypefield := "grant_type"
	Assertionfield := "assertion"

	TokenFields = &TokenStringFields{
		ID:           &IDfield,
		SessionToken: &SessionTokenfield,
		Username:     &Usernamefield,
		Password:     &Passwordfield,
		OtpCode:      &OtpCodefield,
		AppName:      &AppNamefield,
		SdkName:      &SdkNamefield,
		SourceIp:     &SourceIpfield,
		CreationTime: &CreationTimefield,
		LastModified: &LastModifiedfield,
		ExpiryTime:   &ExpiryTimefield,
		ServerUuid:   &ServerUuidfield,
		GrantType:    &GrantTypefield,
		Assertion:    &Assertionfield,
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
