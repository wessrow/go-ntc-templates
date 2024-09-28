package models

type CiscoIosShowIpv6Neighbors struct {
	Address	string	`json:"ADDRESS"`
	Age	string	`json:"AGE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
}