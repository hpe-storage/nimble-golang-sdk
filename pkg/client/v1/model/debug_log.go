// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// DebugLog - Method to help log events from outside of storage array to provide context for troubleshooting host-side or array-side issues.
// Export DebugLogFields for advance operations like search filter etc.
var DebugLogFields *DebugLog

func init(){
	Tagfield:= "tag"
	Messagefield:= "message"
		
	DebugLogFields= &DebugLog{
		Tag:    &Tagfield,
		Message:&Messagefield,
	}
}

type DebugLog struct {
	// Level - Log level.
   	Level *NsTraceLevel `json:"level,omitempty"`
	// Tag - Specifies the context of the message.
 	Tag *string `json:"tag,omitempty"`
	// Message - The message to log.
 	Message *string `json:"message,omitempty"`
}
