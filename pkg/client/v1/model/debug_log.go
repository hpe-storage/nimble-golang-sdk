// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// DebugLog - Method to help log events from outside of storage array to provide context for troubleshooting host-side or array-side issues.
// Export DebugLogFields for advance operations like search filter etc.
var DebugLogFields *DebugLog

func init() {

	DebugLogFields = &DebugLog{
		Tag:     "tag",
		Message: "message",
	}
}

type DebugLog struct {
	// Level - Log level.
	Level *NsTraceLevel `json:"level,omitempty"`
	// Tag - Specifies the context of the message.
	Tag string `json:"tag,omitempty"`
	// Message - The message to log.
	Message string `json:"message,omitempty"`
}

// sdk internal struct
type debugLog struct {
	// Level - Log level.
	Level *NsTraceLevel `json:"level,omitempty"`
	// Tag - Specifies the context of the message.
	Tag *string `json:"tag,omitempty"`
	// Message - The message to log.
	Message *string `json:"message,omitempty"`
}

// EncodeDebugLog - Transform DebugLog to debugLog type
func EncodeDebugLog(request interface{}) (*debugLog, error) {
	reqDebugLog := request.(*DebugLog)
	byte, err := json.Marshal(reqDebugLog)
	objPtr := &debugLog{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeDebugLog Transform debugLog to DebugLog type
func DecodeDebugLog(request interface{}) (*DebugLog, error) {
	reqDebugLog := request.(*debugLog)
	byte, err := json.Marshal(reqDebugLog)
	obj := &DebugLog{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
