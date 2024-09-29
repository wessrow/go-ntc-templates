package models

type CiscoIosShowVrrpAll struct {
	Interface	string	`json:"INTERFACE"`
	Group	string	`json:"GROUP"`
	Addr_family	string	`json:"ADDR_FAMILY"`
	State	string	`json:"STATE"`
	State_duration	string	`json:"STATE_DURATION"`
	Virtual_ip_address	string	`json:"VIRTUAL_IP_ADDRESS"`
	Virtual_mac_address	string	`json:"VIRTUAL_MAC_ADDRESS"`
	Adv_interval	string	`json:"ADV_INTERVAL"`
	Preempt	string	`json:"PREEMPT"`
	Priority	string	`json:"PRIORITY"`
	Priority_configured	string	`json:"PRIORITY_CONFIGURED"`
	Track_obj	string	`json:"TRACK_OBJ"`
	Track_status	string	`json:"TRACK_STATUS"`
	Track_action	string	`json:"TRACK_ACTION"`
	Master_ip_address	string	`json:"MASTER_IP_ADDRESS"`
	Master_priority	string	`json:"MASTER_PRIORITY"`
	Master_adv_interval	string	`json:"MASTER_ADV_INTERVAL"`
	Master_down_interval	string	`json:"MASTER_DOWN_INTERVAL"`
}

var CiscoIosShowVrrpAll_Template = `Value INTERFACE (\S+)
Value GROUP (\d+)
Value ADDR_FAMILY (\S+)
Value STATE (\w+)
Value STATE_DURATION (.+)
Value VIRTUAL_IP_ADDRESS ([a-fA-F\d\.\:]+)
Value VIRTUAL_MAC_ADDRESS ([a-fA-F\d\.\:]+)
Value ADV_INTERVAL (\S+)
Value PREEMPT (\w+)
Value PRIORITY (\d+)
Value PRIORITY_CONFIGURED (\d+)
Value TRACK_OBJ (\d+)
Value TRACK_STATUS (\S+)
Value TRACK_ACTION (\S+(\s+\S+)?)
Value MASTER_IP_ADDRESS (\S+)
Value MASTER_PRIORITY (\S+)
Value MASTER_ADV_INTERVAL (\S+)
Value MASTER_DOWN_INTERVAL (\S+)

Start
  ^\S+\s+-\s+Group\s+\d+.*$$ -> Continue.Record
  ^\s*${INTERFACE}\s+-\s+Group\s+${GROUP}(\s+-\s+Address-Family\s+${ADDR_FAMILY})?\s*$$
  ^\s+State is\s+${STATE}.*\s*$$
  ^\s*State duration\s+${STATE_DURATION}\s*$$
  ^\s+Virtual IP address is\s+${VIRTUAL_IP_ADDRESS}\s*$$
  ^\s+Virtual MAC address is\s+${VIRTUAL_MAC_ADDRESS}\s*$$
  ^\s+Advertisement interval is\s+${ADV_INTERVAL}
  ^\s+Preemption\s+${PREEMPT}\s*$$
  ^\s+Priority is\s+${PRIORITY}(\s+\((cfgd|[cC]onfigured)\s+${PRIORITY_CONFIGURED}\))?\s*$$
  ^\s*Track object\s+${TRACK_OBJ}\s+state\s+${TRACK_STATUS}\s+${TRACK_ACTION}\s*$$
  # ignore authentication output since later VRRP versions do not even support it
  ^\s+Authentication\s
  # ignoring the "local" designation since we can use the STATE value to determine who the master is
  ^\s+Master Router is\s+${MASTER_IP_ADDRESS}(\s+\(local\))?, priority is\s+${MASTER_PRIORITY}\s*$$
  ^\s+Master Advertisement interval is\s+${MASTER_ADV_INTERVAL}
  ^\s+Master Down interval is\s+${MASTER_DOWN_INTERVAL}
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^\s*$$
  ^. -> Error

`