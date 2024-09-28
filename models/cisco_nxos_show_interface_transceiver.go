package models

type CiscoNxosShowInterfaceTransceiver struct {
	Interface	string	`json:"INTERFACE"`
	Status	string	`json:"STATUS"`
	Manufacturer	string	`json:"MANUFACTURER"`
	Type	string	`json:"TYPE"`
	Serial	string	`json:"SERIAL"`
	Part_number	string	`json:"PART_NUMBER"`
	Product_id	string	`json:"PRODUCT_ID"`
}