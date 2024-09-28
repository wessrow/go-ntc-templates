package models

type CiscoNxosShowAccessLists struct {
	Name	string	`json:"NAME"`
	Sn	string	`json:"SN"`
	Action	string	`json:"ACTION"`
	Protocol	string	`json:"PROTOCOL"`
	Source	string	`json:"SOURCE"`
	Range	string	`json:"RANGE"`
	Port	string	`json:"PORT"`
	Destination	string	`json:"DESTINATION"`
	Modifier	string	`json:"MODIFIER"`
	Remark	string	`json:"REMARK"`
}