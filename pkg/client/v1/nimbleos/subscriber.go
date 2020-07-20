// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Subscriber - Subscribers are websocket based notification clients that can subscribe to interesting operations and events and recieve notifications whenever the subscribed to operations and events happen on the array.
// Export SubscriberFields for advance operations like search filter etc.
var SubscriberFields *Subscriber

func init() {

	SubscriberFields = &Subscriber{
		ID:   "id",
		Type: "type",
	}
}

type Subscriber struct {
	// ID - Identifier for subscriber.
	ID string `json:"id,omitempty"`
	// Type - This is generally used to indicate the type of subscriber e.g. SMIS/GUI/ThirdParty etc. This is free form and doesn't need to be unique.
	Type string `json:"type,omitempty"`
	// RenewInterval - The interval in seconds within which the subscriber is expected to send a renew message over the websocket channel in case there is no traffic on the websocket channel.
	RenewInterval int64 `json:"renew_interval,omitempty"`
	// RenewResponseTimeout - The interval in seconds after the subscriber sends a renew message within which the subscriber expects to get a response.
	RenewResponseTimeout int64 `json:"renew_response_timeout,omitempty"`
	// IsConnected - True if the subscriber has an active websocket connection.
	IsConnected *bool `json:"is_connected,omitempty"`
	// NotificationCount - Number of notifications sent to subscriber.
	NotificationCount int64 `json:"notification_count,omitempty"`
	// Force - Forcibly modify a connected subscriber.
	Force *bool `json:"force,omitempty"`
}

// sdk internal struct
type subscriber struct {
	ID                   *string `json:"id,omitempty"`
	Type                 *string `json:"type,omitempty"`
	RenewInterval        *int64  `json:"renew_interval,omitempty"`
	RenewResponseTimeout *int64  `json:"renew_response_timeout,omitempty"`
	IsConnected          *bool   `json:"is_connected,omitempty"`
	NotificationCount    *int64  `json:"notification_count,omitempty"`
	Force                *bool   `json:"force,omitempty"`
}

// EncodeSubscriber - Transform Subscriber to subscriber type
func EncodeSubscriber(request interface{}) (*subscriber, error) {
	reqSubscriber := request.(*Subscriber)
	byte, err := json.Marshal(reqSubscriber)
	if err != nil {
		return nil, err
	}
	respSubscriberPtr := &subscriber{}
	err = json.Unmarshal(byte, respSubscriberPtr)
	return respSubscriberPtr, err
}

// DecodeSubscriber Transform subscriber to Subscriber type
func DecodeSubscriber(request interface{}) (*Subscriber, error) {
	reqSubscriber := request.(*subscriber)
	byte, err := json.Marshal(reqSubscriber)
	if err != nil {
		return nil, err
	}
	respSubscriber := &Subscriber{}
	err = json.Unmarshal(byte, respSubscriber)
	return respSubscriber, err
}
