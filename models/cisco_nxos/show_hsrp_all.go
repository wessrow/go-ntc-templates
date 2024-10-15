package cisco_nxos

type ShowHsrpAll struct {
	Group_number               string `json:"GROUP_NUMBER"`
	Active_ip_address          string `json:"ACTIVE_IP_ADDRESS"`
	Authentication             string `json:"AUTHENTICATION"`
	Version                    string `json:"VERSION"`
	Standby_priority           string `json:"STANDBY_PRIORITY"`
	Interface                  string `json:"INTERFACE"`
	Timers_hold_sec            string `json:"TIMERS_HOLD_SEC"`
	Primary_ipv4_address       string `json:"PRIMARY_IPV4_ADDRESS"`
	Timers_hello_sec           string `json:"TIMERS_HELLO_SEC"`
	Active_priority            string `json:"ACTIVE_PRIORITY"`
	Virtual_mac_address        string `json:"VIRTUAL_MAC_ADDRESS"`
	Num_state_changes          string `json:"NUM_STATE_CHANGES"`
	Preempt                    string `json:"PREEMPT"`
	Upper_fwd_treshold         string `json:"UPPER_FWD_TRESHOLD"`
	Active_router              string `json:"ACTIVE_ROUTER"`
	Virtual_mac_address_status string `json:"VIRTUAL_MAC_ADDRESS_STATUS"`
	Last_state_change          string `json:"LAST_STATE_CHANGE"`
	Hsrp_router_state          string `json:"HSRP_ROUTER_STATE"`
	Lower_fwd_treshold         string `json:"LOWER_FWD_TRESHOLD"`
	Active_expire              string `json:"ACTIVE_EXPIRE"`
	Secondary_ipv4_address     string `json:"SECONDARY_IPV4_ADDRESS"`
	Standby_router             string `json:"STANDBY_ROUTER"`
	Standby_expire             string `json:"STANDBY_EXPIRE"`
	Standby_ip_address         string `json:"STANDBY_IP_ADDRESS"`
	Priority                   string `json:"PRIORITY"`
	Configured_priority        string `json:"CONFIGURED_PRIORITY"`
	Session_name               string `json:"SESSION_NAME"`
}

var ShowHsrpAll_Template string = `# Object names are based on pyATS/Genie parser (boolean values are left out)
Value INTERFACE (\S+)
Value VERSION (\d+)
Value GROUP_NUMBER (\d+)
Value PRIORITY (\d+)
Value HSRP_ROUTER_STATE (\S+)
Value CONFIGURED_PRIORITY (\d+)
Value PREEMPT (.*)
Value LOWER_FWD_TRESHOLD (\d+)
Value UPPER_FWD_TRESHOLD (\d+)
# Value TIMERS_HELLO_MSEC_FLAG
# Value TIMERS_HOLD_MSEC_FLAG
Value TIMERS_HELLO_SEC (\d+)
Value TIMERS_HOLD_SEC (\d+)
Value PRIMARY_IPV4_ADDRESS (\S+)
# Value VIRTUAL_IP_LEARN
Value SECONDARY_IPV4_ADDRESS (\S+)
Value ACTIVE_ROUTER (\S+)
Value ACTIVE_EXPIRE (\d+\.\d+)
Value ACTIVE_IP_ADDRESS (\S+)
Value ACTIVE_PRIORITY (\d+)
Value STANDBY_ROUTER (\S+)
Value STANDBY_EXPIRE (\d+\.\d+)
Value STANDBY_IP_ADDRESS (\S+)
Value STANDBY_PRIORITY (\d+)
Value AUTHENTICATION (\S+)
Value VIRTUAL_MAC_ADDRESS (\S+)
Value VIRTUAL_MAC_ADDRESS_STATUS (.*)
Value NUM_STATE_CHANGES (\d+)
Value LAST_STATE_CHANGE (\S+)
Value SESSION_NAME (\S+)

Start
  ^${INTERFACE}\s+-\s+Group\s${GROUP_NUMBER}\s+\(HSRP-V${VERSION}\)\s\(IPv[46]\)\s*$$
  ^\s+Local\s+state\s+is\s+${HSRP_ROUTER_STATE},\s+priority\s+${PRIORITY}\s+\(Cfged\s+${CONFIGURED_PRIORITY}\),\s+${PREEMPT}\s*$$  
  ^\s+.*lower:\s+${LOWER_FWD_TRESHOLD}\s+upper:\s+${UPPER_FWD_TRESHOLD}\s*$$
  ^\s+Preemption\s+Delay.*$$
  ^\s+Hellotime\s+${TIMERS_HELLO_SEC}.*holdtime\s+${TIMERS_HOLD_SEC}.*$$
  ^\s+Next\s+hello\s+sent.*$$
  ^\s+Virtual\s+IP\s+address\s+is\s+${PRIMARY_IPV4_ADDRESS}\s+.*$$
  ^\s+Secondary\s+Virtual\s+IP\s+address\s+is\s+${SECONDARY_IPV4_ADDRESS}\s*$$
  ^\s+Active\s+router\s+is\s+${ACTIVE_ROUTER}\s*$$
  ^\s+Active\s+router\s+is\s+${ACTIVE_IP_ADDRESS}\s?,\s+priority\s+${ACTIVE_PRIORITY}\s+expires\s+in\s+${ACTIVE_EXPIRE}.*$$
  ^\s+Standby\s+router\s+is\s+${STANDBY_ROUTER}\s*$$
  ^\s+Standby\s+router\s+is\s+${STANDBY_IP_ADDRESS}\s?,\s+priority\s+${STANDBY_PRIORITY}\s+expires\s+in\s+${STANDBY_EXPIRE}.*$$
  # catch clear text authetication:
  ^\s+Authentication\stext\s+.*$$
  ^\s+Authentication\s+${AUTHENTICATION},.*$$
  ^\s+Virtual\s+mac\s+address\s+is\s+${VIRTUAL_MAC_ADDRESS} \(${VIRTUAL_MAC_ADDRESS_STATUS}\)\s*$$
  ^\s+${NUM_STATE_CHANGES}\s+state\s+changes,\s+last\s+state\s+change\s+${LAST_STATE_CHANGE}\s*$$
  ^\s+IP\s+redundancy\s+name\s+is\s+${SESSION_NAME}\s.*$$
  ^\s+Secondary VIP.*$$
  # Catch secondary VIPs:
  ^\s+\d+\.\d+\.\d+\.\d+$$
  # Record data when empty line is found
  ^\s*$$ -> Record
  ^. -> Error
`
