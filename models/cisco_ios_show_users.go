package models

type CiscoIosShowUsers struct {
	Line	string	`json:"LINE"`
	User	string	`json:"USER"`
	Hosts	string	`json:"HOSTS"`
	Idle	string	`json:"IDLE"`
	Location	string	`json:"LOCATION"`
}