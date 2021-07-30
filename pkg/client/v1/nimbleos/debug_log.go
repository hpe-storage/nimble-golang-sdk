// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// DebugLog - Method to help log events from outside of storage array to provide context for troubleshooting host-side or array-side issues.
// Export DebugLogFields for advance operations like search filter etc.
var DebugLogFields *DebugLogStringFields

func init() {
	Levelfield := "level"
	Tagfield := "tag"
	Messagefield := "message"

	DebugLogFields = &DebugLogStringFields{
		Level:   &Levelfield,
		Tag:     &Tagfield,
		Message: &Messagefield,
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

// Struct for DebugLogFields
type DebugLogStringFields struct {
	Level   *string
	Tag     *string
	Message *string
}
