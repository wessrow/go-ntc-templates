package models

type CiscoNxosShowVrfInterface struct {
	Interface	string	`json:"INTERFACE"`
	Name	string	`json:"NAME"`
	Id	string	`json:"ID"`
	Origin	string	`json:"ORIGIN"`
}