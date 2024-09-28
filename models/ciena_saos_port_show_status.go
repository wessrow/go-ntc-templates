package models

type CienaSaosPortShowStatus struct {
	Name	string	`json:"NAME"`
	Description	string	`json:"DESCRIPTION"`
	Link	string	`json:"LINK"`
	Duration	string	`json:"DURATION"`
	Transceiver	string	`json:"TRANSCEIVER"`
	Stp	string	`json:"STP"`
	Speed_duplex	string	`json:"SPEED_DUPLEX"`
	Mtu	string	`json:"MTU"`
	Flow_control	string	`json:"FLOW_CONTROL"`
}