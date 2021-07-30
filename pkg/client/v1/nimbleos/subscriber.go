// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Subscriber - Subscribers are websocket based notification clients that can subscribe to interesting operations and events and recieve notifications whenever the subscribed to operations and events happen on the array.

// Export SubscriberFields provides field names to use in filter parameters, for example.
var SubscriberFields *SubscriberStringFields

func init() {
	fieldID := "id"
	fieldType := "type"
	fieldRenewInterval := "renew_interval"
	fieldRenewResponseTimeout := "renew_response_timeout"
	fieldIsConnected := "is_connected"
	fieldNotificationCount := "notification_count"
	fieldForce := "force"

	SubscriberFields = &SubscriberStringFields{
		ID:                   &fieldID,
		Type:                 &fieldType,
		RenewInterval:        &fieldRenewInterval,
		RenewResponseTimeout: &fieldRenewResponseTimeout,
		IsConnected:          &fieldIsConnected,
		NotificationCount:    &fieldNotificationCount,
		Force:                &fieldForce,
	}
}

type Subscriber struct {
	// ID - Identifier for subscriber.
	ID *string `json:"id,omitempty"`
	// Type - This is generally used to indicate the type of subscriber e.g. SMIS/GUI/ThirdParty etc. This is free form and doesn't need to be unique.
	Type *string `json:"type,omitempty"`
	// RenewInterval - The interval in seconds within which the subscriber is expected to send a renew message over the websocket channel in case there is no traffic on the websocket channel.
	RenewInterval *int64 `json:"renew_interval,omitempty"`
	// RenewResponseTimeout - The interval in seconds after the subscriber sends a renew message within which the subscriber expects to get a response.
	RenewResponseTimeout *int64 `json:"renew_response_timeout,omitempty"`
	// IsConnected - True if the subscriber has an active websocket connection.
	IsConnected *bool `json:"is_connected,omitempty"`
	// NotificationCount - Number of notifications sent to subscriber.
	NotificationCount *int64 `json:"notification_count,omitempty"`
	// Force - Forcibly modify a connected subscriber.
	Force *bool `json:"force,omitempty"`
}

// Struct for SubscriberFields
type SubscriberStringFields struct {
	ID                   *string
	Type                 *string
	RenewInterval        *string
	RenewResponseTimeout *string
	IsConnected          *string
	NotificationCount    *string
	Force                *string
}
