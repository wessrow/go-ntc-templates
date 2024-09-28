package models

type CiscoIosShowAuthenticationSessionsMethodDetails struct {
	Interface	string	`json:"INTERFACE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Ipv4_address	string	`json:"IPV4_ADDRESS"`
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Username	string	`json:"USERNAME"`
	Device_type	string	`json:"DEVICE_TYPE"`
	Device_name	string	`json:"DEVICE_NAME"`
	Status	string	`json:"STATUS"`
	Domain	string	`json:"DOMAIN"`
	Operational_host_mode	string	`json:"OPERATIONAL_HOST_MODE"`
	Operational_control_dir	string	`json:"OPERATIONAL_CONTROL_DIR"`
	Session_timeout	string	`json:"SESSION_TIMEOUT"`
	Timeout_action	string	`json:"TIMEOUT_ACTION"`
	Accounting_update_seconds	string	`json:"ACCOUNTING_UPDATE_SECONDS"`
	Accounting_update_remaining_seconds	string	`json:"ACCOUNTING_UPDATE_REMAINING_SECONDS"`
	Current_policy	string	`json:"CURRENT_POLICY"`
	Server_policy_vlan_group	string	`json:"SERVER_POLICY_VLAN_GROUP"`
	Server_policy_sgt	[]string	`json:"SERVER_POLICY_SGT"`
	Server_policy_vn	string	`json:"SERVER_POLICY_VN"`
	Server_session_timeout	string	`json:"SERVER_SESSION_TIMEOUT"`
	Server_template	string	`json:"SERVER_TEMPLATE"`
	Resultant_policy_vlan_group	string	`json:"RESULTANT_POLICY_VLAN_GROUP"`
	Resultant_policy_sgt	[]string	`json:"RESULTANT_POLICY_SGT"`
	Resultant_policy_vn	string	`json:"RESULTANT_POLICY_VN"`
	Method_type_list	[]string	`json:"METHOD_TYPE_LIST"`
	Method_state_list	[]string	`json:"METHOD_STATE_LIST"`
}