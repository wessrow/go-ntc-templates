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