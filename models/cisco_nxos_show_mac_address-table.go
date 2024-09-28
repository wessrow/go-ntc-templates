package models

type CiscoNxosShowMacAddressTable struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
	Age	string	`json:"AGE"`
	Secure	string	`json:"SECURE"`
	Ntfy	string	`json:"NTFY"`
	Ports	string	`json:"PORTS"`
}