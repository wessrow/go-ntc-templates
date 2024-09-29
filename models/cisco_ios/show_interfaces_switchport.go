package cisco_ios 

type ShowInterfacesSwitchport struct {
	Interface	string	`json:"INTERFACE"`
	Switchport	string	`json:"SWITCHPORT"`
	Switchport_monitor	string	`json:"SWITCHPORT_MONITOR"`
	Switchport_negotiation	string	`json:"SWITCHPORT_NEGOTIATION"`
	Mode	string	`json:"MODE"`
	Admin_mode	string	`json:"ADMIN_MODE"`
	Access_vlan	string	`json:"ACCESS_VLAN"`
	Native_vlan	string	`json:"NATIVE_VLAN"`
	Voice_vlan	string	`json:"VOICE_VLAN"`
	Trunking_vlans	[]string	`json:"TRUNKING_VLANS"`
}

var ShowInterfacesSwitchport_Template = `Value Required INTERFACE (\S+)
Value SWITCHPORT (.+?)
Value SWITCHPORT_MONITOR (.+?)
Value SWITCHPORT_NEGOTIATION (.+?)
Value MODE (.+?)
Value ADMIN_MODE (.+?)
Value ACCESS_VLAN (\d+|unassigned)
Value NATIVE_VLAN (\d+)
Value VOICE_VLAN (\S+)
Value List TRUNKING_VLANS (\S+?)

Start
  ^Name: -> Continue.Record
  ^Name:\s+${INTERFACE}
  ^\s*Switchport:\s+${SWITCHPORT}$$
  ^\s*Switchport\s+Monitor:\s+${SWITCHPORT_MONITOR}$$
  ^\s*Operational\s+Mode:\s+${MODE}$$
  ^\s*Negotiation\s+of\s+Trunking:\s+${SWITCHPORT_NEGOTIATION}$$
  ^\s*Access\s+Mode\s+VLAN:\s+${ACCESS_VLAN}
  ^\s*Trunking\s+Native\s+Mode\s+VLAN:\s+${NATIVE_VLAN}
  ^\s*Voice\s+VLAN:\s+${VOICE_VLAN}
  ^\s*Trunking\s+VLANs\s+Enabled:\s+${TRUNKING_VLANS},\s*$$ -> Trunk
  ^\s*Trunking\s+VLANs\s+Enabled:\s+${TRUNKING_VLANS}$$
  ^\s*Trunking\s+VLANs\s+Active
  ^\s*Administrative\s+Mode:\s+${ADMIN_MODE}$$
  ^\s*(?:Operational|Administrative)\s+(?:Trunking|Native\s+VLAN|private-vlan)
  ^\s*Voice\s+VLAN:
  ^\s*Pruning\s+VLANs
  ^\s*Capture\s+(?:Mode|VLANs)
  ^\s*Autostate\s+mode\s+exclude
  ^\s*Protected
  ^\s*Unknown\s+(unicast|multicast)
  ^\s*Vepa\s+Enabled
  ^\s*App\s+Interface
  ^\s*Appliance\s+trust
  ^\s*Operational\s+Dot1q\s+Ethertype
  ^\s*Priority\s+for\s+untagged\s+frame
  ^\s*Override\s+vlan\s+tag\s+priority
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error

Trunk
  ^\s+${TRUNKING_VLANS},\s*$$
  ^\s+${TRUNKING_VLANS}\s*$$ -> Start
  ^. -> Error
`