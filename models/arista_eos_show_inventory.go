package models

type AristaEosShowInventory struct {
	Port	string	`json:"PORT"`
	Pid	string	`json:"PID"`
	Sn	string	`json:"SN"`
	Descr	string	`json:"DESCR"`
	Vid	string	`json:"VID"`
}