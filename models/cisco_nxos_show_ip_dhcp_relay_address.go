package models

type CiscoNxosShowIpDhcpRelayAddress struct {
	Interface	string	`json:"INTERFACE"`
	Relay_address	string	`json:"RELAY_ADDRESS"`
	Vrf	string	`json:"VRF"`
}