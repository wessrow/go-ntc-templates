package models

type CiscoS300ShowMacAddressTable struct {
	Destination_address	string	`json:"DESTINATION_ADDRESS"`
	Type	string	`json:"TYPE"`
	Vlan_id	string	`json:"VLAN_ID"`
	Destination_port	string	`json:"DESTINATION_PORT"`
}