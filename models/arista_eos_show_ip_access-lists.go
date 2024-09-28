package models

type AristaEosShowIpAccessLists struct {
	Name	string	`json:"NAME"`
	Sn	string	`json:"SN"`
	Action	string	`json:"ACTION"`
	Protocol	string	`json:"PROTOCOL"`
	Source	string	`json:"SOURCE"`
	Port_modifier	string	`json:"PORT_MODIFIER"`
	Port_range	string	`json:"PORT_RANGE"`
	Destination	string	`json:"DESTINATION"`
	Modifier	string	`json:"MODIFIER"`
}