package models

type CienaSaosLldpShowNeighbors struct {
	Remote_ipv4_address	string	`json:"REMOTE_IPV4_ADDRESS"`
	Remote_ipv6_address	string	`json:"REMOTE_IPV6_ADDRESS"`
	System_name	string	`json:"SYSTEM_NAME"`
	System_description	string	`json:"SYSTEM_DESCRIPTION"`
	Local_port	string	`json:"LOCAL_PORT"`
	Remote_port	string	`json:"REMOTE_PORT"`
	Remote_chassis_id	string	`json:"REMOTE_CHASSIS_ID"`
	Remote_mgmt_ipv4_address	[]string	`json:"REMOTE_MGMT_IPV4_ADDRESS"`
	Remote_mgmt_ipv6_address	[]string	`json:"REMOTE_MGMT_IPV6_ADDRESS"`
	Remote_system_name	string	`json:"REMOTE_SYSTEM_NAME"`
	Remote_system_description	string	`json:"REMOTE_SYSTEM_DESCRIPTION"`
	Remote_synce_support	string	`json:"REMOTE_SYNCE_SUPPORT"`
	Remote_synce_config	string	`json:"REMOTE_SYNCE_CONFIG"`
}