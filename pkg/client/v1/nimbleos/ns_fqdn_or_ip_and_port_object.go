// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsFqdnOrIpAndPortObjectFields provides field names to use in filter parameters, for example.
var NsFqdnOrIpAndPortObjectFields *NsFqdnOrIpAndPortObjectFieldHandles

func init() {
	NsFqdnOrIpAndPortObjectFields = &NsFqdnOrIpAndPortObjectFieldHandles{
		SyslogServer: "syslog_server",
		SyslogPort:   "syslog_port",
	}
}

// NsFqdnOrIpAndPortObject - Object wrapper of Fqdn/IP Address and Port numbers.
type NsFqdnOrIpAndPortObject struct {
	// SyslogServer - A Fqdn/IP Address.
	SyslogServer *string `json:"syslog_server,omitempty"`
	// SyslogPort - Syslog port number.
	SyslogPort *int64 `json:"syslog_port,omitempty"`
}

// NsFqdnOrIpAndPortObjectFieldHandles provides a string representation for each AccessControlRecord field.
type NsFqdnOrIpAndPortObjectFieldHandles struct {
	SyslogServer string
	SyslogPort   string
}
