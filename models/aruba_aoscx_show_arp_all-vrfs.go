package models

type ArubaAoscxShowArpAllVrfs struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Port_id	string	`json:"PORT_ID"`
	Physical_port	string	`json:"PHYSICAL_PORT"`
	State	string	`json:"STATE"`
	Vrf	string	`json:"VRF"`
}