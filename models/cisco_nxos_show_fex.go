package models

type CiscoNxosShowFex struct {
	Number	string	`json:"NUMBER"`
	Descr	string	`json:"DESCR"`
	State	string	`json:"STATE"`
	Model	string	`json:"MODEL"`
	Serial	string	`json:"SERIAL"`
}