package models

type ArubaAoscxShowMacAddressTable struct {
	Mac_address	string	`json:"MAC_ADDRESS"`
	Vlan_id	string	`json:"VLAN_ID"`
	Type	string	`json:"TYPE"`
	Port	string	`json:"PORT"`
}