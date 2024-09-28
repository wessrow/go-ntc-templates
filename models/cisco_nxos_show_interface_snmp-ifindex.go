package models

type CiscoNxosShowInterfaceSnmpIfindex struct {
	Port	string	`json:"PORT"`
	Ifmib	string	`json:"IFMIB"`
	Ifindex_hex	string	`json:"IFINDEX_HEX"`
}