// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// AuditLog - View audit log.
// Export AuditLogFields for advance operations like search filter etc.
var AuditLogFields *AuditLog

func init() {

	AuditLogFields = &AuditLog{
		ID:               "id",
		ObjectId:         "object_id",
		ObjectName:       "object_name",
		Scope:            "scope",
		ErrorCode:        "error_code",
		UserId:           "user_id",
		UserName:         "user_name",
		UserFullName:     "user_full_name",
		SourceIp:         "source_ip",
		ExtUserId:        "ext_user_id",
		ExtUserGroupId:   "ext_user_group_id",
		ExtUserGroupName: "ext_user_group_name",
		AppName:          "app_name",
		AccessType:       "access_type",
		Activity:         "activity",
	}
}

type AuditLog struct {
	// ID - Identifier for the audit log record.
	ID string `json:"id,omitempty"`
	// Type - Identifier for type of audit log record.
	Type int64 `json:"type,omitempty"`
	// ObjectId - Identifier of object operated upon.
	ObjectId string `json:"object_id,omitempty"`
	// ObjectName - Name of object operated upon.
	ObjectName string `json:"object_name,omitempty"`
	// ObjectType - Type of the object being operated upon.
	ObjectType *NsObjectType `json:"object_type,omitempty"`
	// Scope - Scope within which object exists, for example, name of the array for a NIC.
	Scope string `json:"scope,omitempty"`
	// Time - Time when this operation was performed.
	Time int64 `json:"time,omitempty"`
	// Status - Status of the operation -- success or failure.
	Status *NsOperationStatus `json:"status,omitempty"`
	// ErrorCode - If the operation has failed, this indicates the error code corresponding to the failure.
	ErrorCode string `json:"error_code,omitempty"`
	// UserId - Identifier of the user who performed the operation.
	UserId string `json:"user_id,omitempty"`
	// UserName - Username of the user who performed the operation.
	UserName string `json:"user_name,omitempty"`
	// UserFullName - Full name of the user who performed the operation.
	UserFullName string `json:"user_full_name,omitempty"`
	// SourceIp - IP address from where the operation request originated.
	SourceIp string `json:"source_ip,omitempty"`
	// ExtUserId - The user id of an external user.
	ExtUserId string `json:"ext_user_id,omitempty"`
	// ExtUserGroupId - The group ID of an external user.
	ExtUserGroupId string `json:"ext_user_group_id,omitempty"`
	// ExtUserGroupName - The group name of an external user.
	ExtUserGroupName string `json:"ext_user_group_name,omitempty"`
	// AppName - Name of application from where the operation request was issued, for example, pam, VSS Agent, etc.
	AppName string `json:"app_name,omitempty"`
	// AccessType - Name of access on how the operation request was issued, for example, GUI, CLI or API.
	AccessType string `json:"access_type,omitempty"`
	// Category - Category of the audit log record.
	Category *NsAuditCategory `json:"category,omitempty"`
	// ActivityType - Type of activity performed, for example, create, update or delete.
	ActivityType *NsAuditOperationType `json:"activity_type,omitempty"`
	// Activity - Description of activity performed and recorded in audit log.
	Activity string `json:"activity,omitempty"`
}

// sdk internal struct
type auditLog struct {
	// ID - Identifier for the audit log record.
	ID *string `json:"id,omitempty"`
	// Type - Identifier for type of audit log record.
	Type *int64 `json:"type,omitempty"`
	// ObjectId - Identifier of object operated upon.
	ObjectId *string `json:"object_id,omitempty"`
	// ObjectName - Name of object operated upon.
	ObjectName *string `json:"object_name,omitempty"`
	// ObjectType - Type of the object being operated upon.
	ObjectType *NsObjectType `json:"object_type,omitempty"`
	// Scope - Scope within which object exists, for example, name of the array for a NIC.
	Scope *string `json:"scope,omitempty"`
	// Time - Time when this operation was performed.
	Time *int64 `json:"time,omitempty"`
	// Status - Status of the operation -- success or failure.
	Status *NsOperationStatus `json:"status,omitempty"`
	// ErrorCode - If the operation has failed, this indicates the error code corresponding to the failure.
	ErrorCode *string `json:"error_code,omitempty"`
	// UserId - Identifier of the user who performed the operation.
	UserId *string `json:"user_id,omitempty"`
	// UserName - Username of the user who performed the operation.
	UserName *string `json:"user_name,omitempty"`
	// UserFullName - Full name of the user who performed the operation.
	UserFullName *string `json:"user_full_name,omitempty"`
	// SourceIp - IP address from where the operation request originated.
	SourceIp *string `json:"source_ip,omitempty"`
	// ExtUserId - The user id of an external user.
	ExtUserId *string `json:"ext_user_id,omitempty"`
	// ExtUserGroupId - The group ID of an external user.
	ExtUserGroupId *string `json:"ext_user_group_id,omitempty"`
	// ExtUserGroupName - The group name of an external user.
	ExtUserGroupName *string `json:"ext_user_group_name,omitempty"`
	// AppName - Name of application from where the operation request was issued, for example, pam, VSS Agent, etc.
	AppName *string `json:"app_name,omitempty"`
	// AccessType - Name of access on how the operation request was issued, for example, GUI, CLI or API.
	AccessType *string `json:"access_type,omitempty"`
	// Category - Category of the audit log record.
	Category *NsAuditCategory `json:"category,omitempty"`
	// ActivityType - Type of activity performed, for example, create, update or delete.
	ActivityType *NsAuditOperationType `json:"activity_type,omitempty"`
	// Activity - Description of activity performed and recorded in audit log.
	Activity *string `json:"activity,omitempty"`
}

// EncodeAuditLog - Transform AuditLog to auditLog type
func EncodeAuditLog(request interface{}) (*auditLog, error) {
	reqAuditLog := request.(*AuditLog)
	byte, err := json.Marshal(reqAuditLog)
	objPtr := &auditLog{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeAuditLog Transform auditLog to AuditLog type
func DecodeAuditLog(request interface{}) (*AuditLog, error) {
	reqAuditLog := request.(*auditLog)
	byte, err := json.Marshal(reqAuditLog)
	obj := &AuditLog{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
