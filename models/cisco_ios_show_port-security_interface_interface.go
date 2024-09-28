package models

type CiscoIosShowPortSecurityInterfaceInterface struct {
	Port_security	string	`json:"PORT_SECURITY"`
	Port_status	string	`json:"PORT_STATUS"`
	Violation_mode	string	`json:"VIOLATION_MODE"`
	Aging_time	string	`json:"AGING_TIME"`
	Aging_type	string	`json:"AGING_TYPE"`
	Ss_addr_aging	string	`json:"SS_ADDR_AGING"`
	Max_mac_addrs	string	`json:"MAX_MAC_ADDRS"`
	Total_mac_addrs	string	`json:"TOTAL_MAC_ADDRS"`
	Config_mac_addrs	string	`json:"CONFIG_MAC_ADDRS"`
	Sticky_mac_addrs	string	`json:"STICKY_MAC_ADDRS"`
	Last_src_mac_addr_vlan	string	`json:"LAST_SRC_MAC_ADDR_VLAN"`
	Violation_count	string	`json:"VIOLATION_COUNT"`
}