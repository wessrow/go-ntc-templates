package models

type CiscoIosShowAccessList struct {
	Name	string	`json:"NAME"`
	Type	string	`json:"TYPE"`
	Sn	string	`json:"SN"`
	Action	string	`json:"ACTION"`
	Protocol	string	`json:"PROTOCOL"`
	Source	string	`json:"SOURCE"`
	Operator_source_port	string	`json:"OPERATOR_SOURCE_PORT"`
	Operator_destination_port	string	`json:"OPERATOR_DESTINATION_PORT"`
	Source_port	string	`json:"SOURCE_PORT"`
	Destination_port	string	`json:"DESTINATION_PORT"`
	Destination	string	`json:"DESTINATION"`
	Modifier	string	`json:"MODIFIER"`
	Matches	string	`json:"MATCHES"`
}