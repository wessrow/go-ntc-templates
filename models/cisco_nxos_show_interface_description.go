package models

type CiscoNxosShowInterfaceDescription struct {
	Port	string	`json:"PORT"`
	Type	string	`json:"TYPE"`
	Speed	string	`json:"SPEED"`
	Description	string	`json:"DESCRIPTION"`
}