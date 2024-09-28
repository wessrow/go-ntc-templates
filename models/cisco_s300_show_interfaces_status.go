package models

type CiscoS300ShowInterfacesStatus struct {
	Port	string	`json:"PORT"`
	Type	string	`json:"TYPE"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Neg	string	`json:"NEG"`
	Flowctrl	string	`json:"FLOWCTRL"`
	Linkstate	string	`json:"LINKSTATE"`
	Backpressure	string	`json:"BACKPRESSURE"`
	Mdixmode	string	`json:"MDIXMODE"`
}