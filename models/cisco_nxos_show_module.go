package models

type CiscoNxosShowModule struct {
	Module	string	`json:"MODULE"`
	Ports	string	`json:"PORTS"`
	Type	string	`json:"TYPE"`
	Model	string	`json:"MODEL"`
	Status	string	`json:"STATUS"`
}