package models

type CiscoNxosShowHsrpAll struct {
	Interface	string	`json:"INTERFACE"`
	Version	string	`json:"VERSION"`
	Group_number	string	`json:"GROUP_NUMBER"`
	Priority	string	`json:"PRIORITY"`
	Hsrp_router_state	string	`json:"HSRP_ROUTER_STATE"`
	Configured_priority	string	`json:"CONFIGURED_PRIORITY"`
	Preempt	string	`json:"PREEMPT"`
	Lower_fwd_treshold	string	`json:"LOWER_FWD_TRESHOLD"`
	Upper_fwd_treshold	string	`json:"UPPER_FWD_TRESHOLD"`
	Timers_hello_sec	string	`json:"TIMERS_HELLO_SEC"`
	Timers_hold_sec	string	`json:"TIMERS_HOLD_SEC"`
	Primary_ipv4_address	string	`json:"PRIMARY_IPV4_ADDRESS"`
	Secondary_ipv4_address	string	`json:"SECONDARY_IPV4_ADDRESS"`
	Active_router	string	`json:"ACTIVE_ROUTER"`
	Active_expire	string	`json:"ACTIVE_EXPIRE"`
	Active_ip_address	string	`json:"ACTIVE_IP_ADDRESS"`
	Active_priority	string	`json:"ACTIVE_PRIORITY"`
	Standby_router	string	`json:"STANDBY_ROUTER"`
	Standby_expire	string	`json:"STANDBY_EXPIRE"`
	Standby_ip_address	string	`json:"STANDBY_IP_ADDRESS"`
	Standby_priority	string	`json:"STANDBY_PRIORITY"`
	Authentication	string	`json:"AUTHENTICATION"`
	Virtual_mac_address	string	`json:"VIRTUAL_MAC_ADDRESS"`
	Virtual_mac_address_status	string	`json:"VIRTUAL_MAC_ADDRESS_STATUS"`
	Num_state_changes	string	`json:"NUM_STATE_CHANGES"`
	Last_state_change	string	`json:"LAST_STATE_CHANGE"`
	Session_name	string	`json:"SESSION_NAME"`
}