package models

type CiscoIosShowStandbyBrief struct {
	Interface	string	`json:"INTERFACE"`
	Group	string	`json:"GROUP"`
	Priority	string	`json:"PRIORITY"`
	Preempt	string	`json:"PREEMPT"`
	State	string	`json:"STATE"`
	Active	string	`json:"ACTIVE"`
	Standby	string	`json:"STANDBY"`
	Virtual_ip_address	string	`json:"VIRTUAL_IP_ADDRESS"`
}

var CiscoIosShowStandbyBrief_Template = `Value INTERFACE (\S+)
Value GROUP (\d+)
Value PRIORITY (\d+)
Value PREEMPT (.)
Value STATE (\w+)
Value ACTIVE (\S+)
Value STANDBY (\S+)
Value VIRTUAL_IP_ADDRESS (\S+)

Start
  ^Interface\s+Grp\s+Pri\s+P\s+State\s+Active\s+Standby\s+Virtual\s+IP\s*$$ -> Standby
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

Standby
  ^${INTERFACE}\s+${GROUP}\s+${PRIORITY}\s+${PREEMPT}\s+${STATE}\s+${ACTIVE}\s+${STANDBY}\s+${VIRTUAL_IP_ADDRESS} -> Record
  ^${INTERFACE}\s*$$
  ^\s*${GROUP}\s+${PRIORITY}\s+${PREEMPT}\s+${STATE}\s+${ACTIVE}\s+${STANDBY}\s+${VIRTUAL_IP_ADDRESS} -> Record
  ^\s*$$
  ^. -> Error

`