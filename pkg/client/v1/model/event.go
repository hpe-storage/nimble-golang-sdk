/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Event - View events.
// Export EventFields for advance operations like search filter etc.
var EventFields *Event

func init(){
	IDfield:= "id"
	Namefield:= "name"
	Scopefield:= "scope"
	Targetfield:= "target"
	Summaryfield:= "summary"
	Activityfield:= "activity"
	AlarmIDfield:= "alarm_id"
	TenantIDfield:= "tenant_id"
		
	EventFields= &Event{
		ID: &IDfield,
		Name: &Namefield,
		Scope: &Scopefield,
		Target: &Targetfield,
		Summary: &Summaryfield,
		Activity: &Activityfield,
		AlarmID: &AlarmIDfield,
		TenantID: &TenantIDfield,
		
	}
}

type Event struct {
   
    // Identifier for the event record.
    
 	ID *string `json:"id,omitempty"`
   
    // Type of the event record.
    
   	Type *int64 `json:"type,omitempty"`
   
    // Name of alert macro to generate.
    
 	Name *string `json:"name,omitempty"`
   
    // The array name for array level event.
    
 	Scope *string `json:"scope,omitempty"`
   
    // Name of object upon which the event occurred.
    
 	Target *string `json:"target,omitempty"`
   
    // Target type of the event record.
    
   	TargetType *NsEventTargetType `json:"target_type,omitempty"`
   
    // Time when this event happened.
    
   	Timestamp *int64 `json:"timestamp,omitempty"`
   
    // Category of the event record.
    
   	Category *NsEventCategory `json:"category,omitempty"`
   
    // Severity level of the event.
    
   	Severity *NsSeverityLevel `json:"severity,omitempty"`
   
    // Summary of the event.
    
 	Summary *string `json:"summary,omitempty"`
   
    // Description of the event.
    
 	Activity *string `json:"activity,omitempty"`
   
    // The alarm ID if the event is related to an alarm.
    
 	AlarmID *string `json:"alarm_id,omitempty"`
   
    // Arguments provided for event creation in key-value structure. Until KV implementation for events, will ignore keys (though keys must be non-empty) and take values positionally.
    
   	Params []*NsKeyValue `json:"params,omitempty"`
   
    // Tenant ID of the event. This is used to determine what tenant context the event belongs to.
    
 	TenantID *string `json:"tenant_id,omitempty"`
}
