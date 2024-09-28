package models

type CiscoNxosShowInventory struct {
	Name	string	`json:"NAME"`
	Descr	string	`json:"DESCR"`
	Pid	string	`json:"PID"`
	Vid	string	`json:"VID"`
	Sn	string	`json:"SN"`
}