package cisco_ios 

type ShowPortSecurityInterfaceInterface struct {
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

var ShowPortSecurityInterfaceInterface_Template = `Value PORT_SECURITY (\w+)
Value PORT_STATUS (\S+)
Value VIOLATION_MODE ([Pp]rotect|[Rr]estrict|[Ss]hutdown)
Value AGING_TIME (\d{1,4}\smins)
Value AGING_TYPE ([Aa]bsolute|[Ii]nactivity)
Value SS_ADDR_AGING (\w+)
Value MAX_MAC_ADDRS (\d+)
Value TOTAL_MAC_ADDRS (\d+)
Value CONFIG_MAC_ADDRS (\d+)
Value STICKY_MAC_ADDRS (\d+)
Value LAST_SRC_MAC_ADDR_VLAN (\w.+:\d+)
Value VIOLATION_COUNT (\d+)

Start
  ^Port Security\s+:\s+${PORT_SECURITY}
  ^Port Status\s+:\s+${PORT_STATUS}
  ^Violation Mode\s+:\s+${VIOLATION_MODE}
  ^Aging Time\s+:\s+${AGING_TIME}
  ^Aging Type\s+:\s+${AGING_TYPE}
  ^SecureStatic Address Aging\s+:\s+${SS_ADDR_AGING}
  ^Maximum MAC Addresses\s+:\s+${MAX_MAC_ADDRS}
  ^Total MAC Addresses\s+:\s+${TOTAL_MAC_ADDRS}
  ^Configured MAC Addresses\s+:\s+${CONFIG_MAC_ADDRS}
  ^Sticky MAC Addresses\s+:\s+${STICKY_MAC_ADDRS}
  ^Last Source Address:Vlan\s+:\s+${LAST_SRC_MAC_ADDR_VLAN}
  ^Security Violation Count\s+:\s+${VIOLATION_COUNT}
  ^\s*$$
  ^. -> Error
`