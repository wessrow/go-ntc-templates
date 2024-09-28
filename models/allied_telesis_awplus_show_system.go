package models

type AlliedTelesisAwplusShowSystem struct {
	Hardware	[]string	`json:"HARDWARE"`
	Serial	[]string	`json:"SERIAL"`
	Version	string	`json:"VERSION"`
	Hostname	string	`json:"HOSTNAME"`
}