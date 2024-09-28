package models

type ArubaOsShowArp struct {
	Protocol	string	`json:"PROTOCOL"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Age	string	`json:"AGE"`
}