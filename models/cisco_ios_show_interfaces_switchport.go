package models

type CiscoIosShowInterfacesSwitchport struct {
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