package models

type BrocadeFastironShowMacAddress struct {
	Mac_address	string	`json:"MAC_ADDRESS"`
	Destination_port	string	`json:"DESTINATION_PORT"`
	Age	string	`json:"AGE"`
	Vlan_id	string	`json:"VLAN_ID"`
	Type	string	`json:"TYPE"`
}