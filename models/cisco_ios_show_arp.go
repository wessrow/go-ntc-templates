package models

type CiscoIosShowArp struct {
	Protocol	string	`json:"PROTOCOL"`
	Address	string	`json:"ADDRESS"`
	Age_min	string	`json:"AGE_MIN"`
	Hardware_address	string	`json:"HARDWARE_ADDRESS"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
}