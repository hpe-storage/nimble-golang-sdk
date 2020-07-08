// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// Event - View events.
// Export EventFields for advance operations like search filter etc.
var EventFields *Event

func init() {

	EventFields = &Event{
		ID:       "id",
		Name:     "name",
		Scope:    "scope",
		Target:   "target",
		Summary:  "summary",
		Activity: "activity",
		AlarmId:  "alarm_id",
		TenantId: "tenant_id",
	}
}

type Event struct {
	// ID - Identifier for the event record.
	ID string `json:"id,omitempty"`
	// Type - Type of the event record.
	Type int64 `json:"type,omitempty"`
	// Name - Name of alert macro to generate.
	Name string `json:"name,omitempty"`
	// Scope - The array name for array level event.
	Scope string `json:"scope,omitempty"`
	// Target - Name of object upon which the event occurred.
	Target string `json:"target,omitempty"`
	// TargetType - Target type of the event record.
	TargetType *NsEventTargetType `json:"target_type,omitempty"`
	// Timestamp - Time when this event happened.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Category - Category of the event record.
	Category *NsEventCategory `json:"category,omitempty"`
	// Severity - Severity level of the event.
	Severity *NsSeverityLevel `json:"severity,omitempty"`
	// Summary - Summary of the event.
	Summary string `json:"summary,omitempty"`
	// Activity - Description of the event.
	Activity string `json:"activity,omitempty"`
	// AlarmId - The alarm ID if the event is related to an alarm.
	AlarmId string `json:"alarm_id,omitempty"`
	// Params - Arguments provided for event creation in key-value structure. Until KV implementation for events, will ignore keys (though keys must be non-empty) and take values positionally.
	Params []*NsKeyValue `json:"params,omitempty"`
	// TenantId - Tenant ID of the event. This is used to determine what tenant context the event belongs to.
	TenantId string `json:"tenant_id,omitempty"`
}

// sdk internal struct
type event struct {
	// ID - Identifier for the event record.
	ID *string `json:"id,omitempty"`
	// Type - Type of the event record.
	Type *int64 `json:"type,omitempty"`
	// Name - Name of alert macro to generate.
	Name *string `json:"name,omitempty"`
	// Scope - The array name for array level event.
	Scope *string `json:"scope,omitempty"`
	// Target - Name of object upon which the event occurred.
	Target *string `json:"target,omitempty"`
	// TargetType - Target type of the event record.
	TargetType *NsEventTargetType `json:"target_type,omitempty"`
	// Timestamp - Time when this event happened.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// Category - Category of the event record.
	Category *NsEventCategory `json:"category,omitempty"`
	// Severity - Severity level of the event.
	Severity *NsSeverityLevel `json:"severity,omitempty"`
	// Summary - Summary of the event.
	Summary *string `json:"summary,omitempty"`
	// Activity - Description of the event.
	Activity *string `json:"activity,omitempty"`
	// AlarmId - The alarm ID if the event is related to an alarm.
	AlarmId *string `json:"alarm_id,omitempty"`
	// Params - Arguments provided for event creation in key-value structure. Until KV implementation for events, will ignore keys (though keys must be non-empty) and take values positionally.
	Params []*NsKeyValue `json:"params,omitempty"`
	// TenantId - Tenant ID of the event. This is used to determine what tenant context the event belongs to.
	TenantId *string `json:"tenant_id,omitempty"`
}

// EncodeEvent - Transform Event to event type
func EncodeEvent(request interface{}) (*event, error) {
	reqEvent := request.(*Event)
	byte, err := json.Marshal(reqEvent)
	objPtr := &event{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeEvent Transform event to Event type
func DecodeEvent(request interface{}) (*Event, error) {
	reqEvent := request.(*event)
	byte, err := json.Marshal(reqEvent)
	obj := &Event{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
