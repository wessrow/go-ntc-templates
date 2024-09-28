package models

type CiscoIosShowModule struct {
	Module	string	`json:"MODULE"`
	Switch_number	string	`json:"SWITCH_NUMBER"`
	Port	string	`json:"PORT"`
	Cardtype	string	`json:"CARDTYPE"`
	Model	string	`json:"MODEL"`
	Serial	string	`json:"SERIAL"`
}