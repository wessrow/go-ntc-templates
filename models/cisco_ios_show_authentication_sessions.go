package models

type CiscoIosShowAuthenticationSessions struct {
	Interface	string	`json:"INTERFACE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Method	string	`json:"METHOD"`
	Domain	string	`json:"DOMAIN"`
	Status	string	`json:"STATUS"`
	Session	string	`json:"SESSION"`
}