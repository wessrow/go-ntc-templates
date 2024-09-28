package models

type AlliedTelesisAwplusShowArp struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
	Port	string	`json:"PORT"`
}