package models

type CiscoIosShowInterfacesDescription struct {
	Port	string	`json:"PORT"`
	Status	string	`json:"STATUS"`
	Protocol	string	`json:"PROTOCOL"`
	Description	string	`json:"DESCRIPTION"`
}