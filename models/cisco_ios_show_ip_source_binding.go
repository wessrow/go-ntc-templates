package models

type CiscoIosShowIpSourceBinding struct {
	Mac_address	string	`json:"MAC_ADDRESS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Lease	string	`json:"LEASE"`
	Type	string	`json:"TYPE"`
	Vlan_id	string	`json:"VLAN_ID"`
	Interface	string	`json:"INTERFACE"`
}