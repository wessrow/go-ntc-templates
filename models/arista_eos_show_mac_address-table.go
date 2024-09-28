package models

type AristaEosShowMacAddressTable struct {
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
	Vlan_id	string	`json:"VLAN_ID"`
	Destination_port	[]string	`json:"DESTINATION_PORT"`
	Moves	string	`json:"MOVES"`
	Last_move	string	`json:"LAST_MOVE"`
}