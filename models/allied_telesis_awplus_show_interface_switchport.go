package models

type AlliedTelesisAwplusShowInterfaceSwitchport struct {
	Interface	string	`json:"INTERFACE"`
	Mode	string	`json:"MODE"`
	Native_vlan	string	`json:"NATIVE_VLAN"`
	Trunking_vlans	[]string	`json:"TRUNKING_VLANS"`
}