package models

type CiscoNxosShowIpAdjacency struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Pref	string	`json:"PREF"`
	Source	string	`json:"SOURCE"`
	Interface	string	`json:"INTERFACE"`
	Flags	string	`json:"FLAGS"`
}