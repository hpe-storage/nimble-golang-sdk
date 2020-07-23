// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// ProtectionTemplate - Manage protection templates. Protection templates are sets of snapshot schedules, replication schedules, and retention limits that can be used to prefill the protection information when creating new volume collections. A volume collection, once created, is not affected by edits to the protection template that was used to create it. All the volumes assigned to a volume collection use the same settings. You cannot edit or delete the predefined protection templates provided by storage array, but you can create custom protection templates as needed.
// Export ProtectionTemplateFields for advance operations like search filter etc.
var ProtectionTemplateFields *ProtectionTemplate

func init() {
	IDfield := "id"
	Namefield := "name"
	FullNamefield := "full_name"
	SearchNamefield := "search_name"
	Descriptionfield := "description"
	AppServerfield := "app_server"
	AppClusterNamefield := "app_cluster_name"
	AppServiceNamefield := "app_service_name"
	VcenterHostnamefield := "vcenter_hostname"
	VcenterUsernamefield := "vcenter_username"
	VcenterPasswordfield := "vcenter_password"
	AgentHostnamefield := "agent_hostname"
	AgentUsernamefield := "agent_username"
	AgentPasswordfield := "agent_password"

	ProtectionTemplateFields = &ProtectionTemplate{
		ID:              &IDfield,
		Name:            &Namefield,
		FullName:        &FullNamefield,
		SearchName:      &SearchNamefield,
		Description:     &Descriptionfield,
		AppServer:       &AppServerfield,
		AppClusterName:  &AppClusterNamefield,
		AppServiceName:  &AppServiceNamefield,
		VcenterHostname: &VcenterHostnamefield,
		VcenterUsername: &VcenterUsernamefield,
		VcenterPassword: &VcenterPasswordfield,
		AgentHostname:   &AgentHostnamefield,
		AgentUsername:   &AgentUsernamefield,
		AgentPassword:   &AgentPasswordfield,
	}
}

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
