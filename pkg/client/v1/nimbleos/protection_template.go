// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export ProtectionTemplateFields provides field names to use in filter parameters, for example.
var ProtectionTemplateFields *ProtectionTemplateFieldHandles

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldFullName := "full_name"
	fieldSearchName := "search_name"
	fieldDescription := "description"
	fieldReplPriority := "repl_priority"
	fieldAppSync := "app_sync"
	fieldAppServer := "app_server"
	fieldAppId := "app_id"
	fieldAppClusterName := "app_cluster_name"
	fieldAppServiceName := "app_service_name"
	fieldVcenterHostname := "vcenter_hostname"
	fieldVcenterUsername := "vcenter_username"
	fieldVcenterPassword := "vcenter_password"
	fieldAgentHostname := "agent_hostname"
	fieldAgentUsername := "agent_username"
	fieldAgentPassword := "agent_password"
	fieldCreationTime := "creation_time"
	fieldLastModified := "last_modified"
	fieldScheduleList := "schedule_list"

	ProtectionTemplateFields = &ProtectionTemplateFieldHandles{
		ID:              &fieldID,
		Name:            &fieldName,
		FullName:        &fieldFullName,
		SearchName:      &fieldSearchName,
		Description:     &fieldDescription,
		ReplPriority:    &fieldReplPriority,
		AppSync:         &fieldAppSync,
		AppServer:       &fieldAppServer,
		AppId:           &fieldAppId,
		AppClusterName:  &fieldAppClusterName,
		AppServiceName:  &fieldAppServiceName,
		VcenterHostname: &fieldVcenterHostname,
		VcenterUsername: &fieldVcenterUsername,
		VcenterPassword: &fieldVcenterPassword,
		AgentHostname:   &fieldAgentHostname,
		AgentUsername:   &fieldAgentUsername,
		AgentPassword:   &fieldAgentPassword,
		CreationTime:    &fieldCreationTime,
		LastModified:    &fieldLastModified,
		ScheduleList:    &fieldScheduleList,
	}
}

// ProtectionTemplate - Manage protection templates. Protection templates are sets of snapshot schedules, replication schedules, and retention limits that can be used to prefill the protection information when creating new volume collections. A volume collection, once created, is not affected by edits to the protection template that was used to create it. All the volumes assigned to a volume collection use the same settings. You cannot edit or delete the predefined protection templates provided by storage array, but you can create custom protection templates as needed.
type ProtectionTemplate struct {
	// ID - Identifier for protection template.
	ID *string `json:"id,omitempty"`
	// Name - User provided identifier.
	Name *string `json:"name,omitempty"`
	// FullName - Fully qualified name of protection template.
	FullName *string `json:"full_name,omitempty"`
	// SearchName - Name of protection template used for object search.
	SearchName *string `json:"search_name,omitempty"`
	// Description - Text description of protection template.
	Description *string `json:"description,omitempty"`
	// ReplPriority - Replication priority for the protection template with the following choices: {normal | high}.
	ReplPriority *NsReplPriorityType `json:"repl_priority,omitempty"`
	// AppSync - Application synchronization ({none|vss|vmware|generic}).
	AppSync *NsAppSyncType `json:"app_sync,omitempty"`
	// AppServer - Application server hostname.
	AppServer *string `json:"app_server,omitempty"`
	// AppId - Application ID running on the server. Application ID can only be specified if application synchronization is VSS.
	AppId *NsAppIdType `json:"app_id,omitempty"`
	// AppClusterName - If the application is running within a Windows cluster environment then this is the cluster name.
	AppClusterName *string `json:"app_cluster_name,omitempty"`
	// AppServiceName - If the application is running within a Windows cluster environment then this is the instance name of the service running within the cluster environment.
	AppServiceName *string `json:"app_service_name,omitempty"`
	// VcenterHostname - VMware vCenter hostname. Custom port number can be specified with vCenter hostname using :.
	VcenterHostname *string `json:"vcenter_hostname,omitempty"`
	// VcenterUsername - VMware vCenter username.
	VcenterUsername *string `json:"vcenter_username,omitempty"`
	// VcenterPassword - VMware vCenter password.
	VcenterPassword *string `json:"vcenter_password,omitempty"`
	// AgentHostname - Generic Backup agent hostname. Custom port number can be specified with agent hostname using \\":\\".
	AgentHostname *string `json:"agent_hostname,omitempty"`
	// AgentUsername - Generic Backup agent username.
	AgentUsername *string `json:"agent_username,omitempty"`
	// AgentPassword - Generic Backup agent password.
	AgentPassword *string `json:"agent_password,omitempty"`
	// CreationTime - Time when this protection template was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this protection template was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// ScheduleList - List of schedules for this protection policy.
	ScheduleList []*NsSchedule `json:"schedule_list,omitempty"`
}

// ProtectionTemplateFieldHandles provides a string representation for each AccessControlRecord field.
type ProtectionTemplateFieldHandles struct {
	ID              *string
	Name            *string
	FullName        *string
	SearchName      *string
	Description     *string
	ReplPriority    *string
	AppSync         *string
	AppServer       *string
	AppId           *string
	AppClusterName  *string
	AppServiceName  *string
	VcenterHostname *string
	VcenterUsername *string
	VcenterPassword *string
	AgentHostname   *string
	AgentUsername   *string
	AgentPassword   *string
	CreationTime    *string
	LastModified    *string
	ScheduleList    *string
}
