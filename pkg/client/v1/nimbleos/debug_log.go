// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export DebugLogFields provides field names to use in filter parameters, for example.
var DebugLogFields *DebugLogFieldHandles

func init() {
	DebugLogFields = &DebugLogFieldHandles{
		Level:   "level",
		Tag:     "tag",
		Message: "message",
	}
}

// DebugLog - Method to help log events from outside of storage array to provide context for troubleshooting host-side or array-side issues.
type DebugLog struct {
	// Level - Log level.
	Level *NsTraceLevel `json:"level,omitempty"`
	// Tag - Specifies the context of the message.
	Tag *string `json:"tag,omitempty"`
	// Message - The message to log.
	Message *string `json:"message,omitempty"`
}

// DebugLogFieldHandles provides a string representation for each AccessControlRecord field.
type DebugLogFieldHandles struct {
	Level   string
	Tag     string
	Message string
}
