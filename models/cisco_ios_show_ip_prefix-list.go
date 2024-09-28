package models

type CiscoIosShowIpPrefixList struct {
	Name	string	`json:"NAME"`
	Seq	string	`json:"SEQ"`
	Action	string	`json:"ACTION"`
	Network	string	`json:"NETWORK"`
	Netmask	string	`json:"NETMASK"`
	Le	string	`json:"LE"`
	Ge	string	`json:"GE"`
}