package models

type BrocadeFastironShowArp struct {
	Protocol	string	`json:"PROTOCOL"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Age	string	`json:"AGE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
	Port	string	`json:"PORT"`
}