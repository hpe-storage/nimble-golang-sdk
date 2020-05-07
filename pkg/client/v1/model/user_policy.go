/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/UserPolicy


// UserPolicy :
type UserPolicy struct {
   // ID
   ID string `json:"id,omitempty"`
   // AllowedAttempts
   AllowedAttempts  int32 `json:"allowed_attempts,omitempty"`
   // MinLength
   MinLength  int32 `json:"min_length,omitempty"`
   // Upper
   Upper  int32 `json:"upper,omitempty"`
   // Lower
   Lower  int32 `json:"lower,omitempty"`
   // Digit
   Digit  int32 `json:"digit,omitempty"`
   // Special
   Special  int32 `json:"special,omitempty"`
   // PreviousDiff
   PreviousDiff  int32 `json:"previous_diff,omitempty"`
   // NoReuse
   NoReuse  int32 `json:"no_reuse,omitempty"`
   // MaxSessions
   MaxSessions  int32 `json:"max_sessions,omitempty"`
}
