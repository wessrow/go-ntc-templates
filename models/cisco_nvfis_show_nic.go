package models

type CiscoNvfisShowNic struct {
	Slotid	string	`json:"SLOTID"`
	Adapter	string	`json:"ADAPTER"`
	Vendor	string	`json:"VENDOR"`
	Devid	string	`json:"DEVID"`
	Mode	string	`json:"MODE"`
	Devno	string	`json:"DEVNO"`
	Pnics	string	`json:"PNICS"`
}