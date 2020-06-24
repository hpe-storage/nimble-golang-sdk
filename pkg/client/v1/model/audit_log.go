/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// AuditLog - View audit log.
// Export AuditLogFields for advance operations like search filter etc.
var AuditLogFields *AuditLog

func init(){
	IDfield:= "id"
	ObjectIDfield:= "object_id"
	ObjectNamefield:= "object_name"
	Scopefield:= "scope"
	ErrorCodefield:= "error_code"
	UserIDfield:= "user_id"
	UserNamefield:= "user_name"
	UserFullNamefield:= "user_full_name"
	SourceIpfield:= "source_ip"
	ExtUserIDfield:= "ext_user_id"
	ExtUserGroupIDfield:= "ext_user_group_id"
	ExtUserGroupNamefield:= "ext_user_group_name"
	AppNamefield:= "app_name"
	AccessTypefield:= "access_type"
	Activityfield:= "activity"
		
	AuditLogFields= &AuditLog{
		ID: &IDfield,
		ObjectID: &ObjectIDfield,
		ObjectName: &ObjectNamefield,
		Scope: &Scopefield,
		ErrorCode: &ErrorCodefield,
		UserID: &UserIDfield,
		UserName: &UserNamefield,
		UserFullName: &UserFullNamefield,
		SourceIp: &SourceIpfield,
		ExtUserID: &ExtUserIDfield,
		ExtUserGroupID: &ExtUserGroupIDfield,
		ExtUserGroupName: &ExtUserGroupNamefield,
		AppName: &AppNamefield,
		AccessType: &AccessTypefield,
		Activity: &Activityfield,
		
	}
}

type AuditLog struct {
   
    // Identifier for the audit log record.
    
 	ID *string `json:"id,omitempty"`
   
    // Identifier for type of audit log record.
    
   	Type *int64 `json:"type,omitempty"`
   
    // Identifier of object operated upon.
    
 	ObjectID *string `json:"object_id,omitempty"`
   
    // Name of object operated upon.
    
 	ObjectName *string `json:"object_name,omitempty"`
   
    // Type of the object being operated upon.
    
   	ObjectType *NsObjectType `json:"object_type,omitempty"`
   
    // Scope within which object exists, for example, name of the array for a NIC.
    
 	Scope *string `json:"scope,omitempty"`
   
    // Time when this operation was performed.
    
   	Time *int64 `json:"time,omitempty"`
   
    // Status of the operation -- success or failure.
    
   	Status *NsOperationStatus `json:"status,omitempty"`
   
    // If the operation has failed, this indicates the error code corresponding to the failure.
    
 	ErrorCode *string `json:"error_code,omitempty"`
   
    // Identifier of the user who performed the operation.
    
 	UserID *string `json:"user_id,omitempty"`
   
    // Username of the user who performed the operation.
    
 	UserName *string `json:"user_name,omitempty"`
   
    // Full name of the user who performed the operation.
    
 	UserFullName *string `json:"user_full_name,omitempty"`
   
    // IP address from where the operation request originated.
    
 	SourceIp *string `json:"source_ip,omitempty"`
   
    // The user id of an external user.
    
 	ExtUserID *string `json:"ext_user_id,omitempty"`
   
    // The group ID of an external user.
    
 	ExtUserGroupID *string `json:"ext_user_group_id,omitempty"`
   
    // The group name of an external user.
    
 	ExtUserGroupName *string `json:"ext_user_group_name,omitempty"`
   
    // Name of application from where the operation request was issued, for example, pam, VSS Agent, etc.
    
 	AppName *string `json:"app_name,omitempty"`
   
    // Name of access on how the operation request was issued, for example, GUI, CLI or API.
    
 	AccessType *string `json:"access_type,omitempty"`
   
    // Category of the audit log record.
    
   	Category *NsAuditCategory `json:"category,omitempty"`
   
    // Type of activity performed, for example, create, update or delete.
    
   	ActivityType *NsAuditOperationType `json:"activity_type,omitempty"`
   
    // Description of activity performed and recorded in audit log.
    
 	Activity *string `json:"activity,omitempty"`
}
