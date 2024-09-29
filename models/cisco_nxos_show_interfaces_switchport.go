package models

type CiscoNxosShowInterfacesSwitchport struct {
	Interface	string	`json:"INTERFACE"`
	Switchport	string	`json:"SWITCHPORT"`
	Switchport_monitor	string	`json:"SWITCHPORT_MONITOR"`
	Mode	string	`json:"MODE"`
	Access_vlan	string	`json:"ACCESS_VLAN"`
	Access_vlan_name	string	`json:"ACCESS_VLAN_NAME"`
	Native_vlan	string	`json:"NATIVE_VLAN"`
	Native_vlan_name	string	`json:"NATIVE_VLAN_NAME"`
	Trunking_vlans	string	`json:"TRUNKING_VLANS"`
	Voice_vlan	string	`json:"VOICE_VLAN"`
}

var CiscoNxosShowInterfacesSwitchport_Template = `Value Required INTERFACE (\S+)
Value SWITCHPORT (.+)
Value SWITCHPORT_MONITOR (\S+(?:\s+\S+)?)
Value MODE (.+)
Value ACCESS_VLAN (\d+)
Value ACCESS_VLAN_NAME (.+)
Value NATIVE_VLAN (\d+)
Value NATIVE_VLAN_NAME (.+)
Value TRUNKING_VLANS (\S+)
Value VOICE_VLAN (\S+)

Start
  ^Name: -> Continue.Record
  ^Name:\s+${INTERFACE}
  ^\s*Switchport:\s+${SWITCHPORT}
  ^\s*Switchport\s+Monitor:\s+${SWITCHPORT_MONITOR}
  ^\s*Operational\s+Mode:\s+${MODE}
  ^\s*Access\s+Mode\s+VLAN:\s+${ACCESS_VLAN}\s+\(${ACCESS_VLAN_NAME}\)
  ^\s*Access\s+Mode\s+VLAN:\s+${ACCESS_VLAN}
  ^\s*Trunking\s+Native\s+Mode\s+VLAN:\s+${NATIVE_VLAN}\s+\(${NATIVE_VLAN_NAME}\)
  ^\s*Trunking\s+Native\s+Mode\s+VLAN:\s+${NATIVE_VLAN}
  ^\s*Trunking\s+VLANs\s+Allowed:\s+${TRUNKING_VLANS}
  ^\s*Voice\s+VLAN:\s+${VOICE_VLAN}
  ^.+$$
  ^. -> Error

`