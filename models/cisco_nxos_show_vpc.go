package models

type CiscoNxosShowVpc struct {
	Id	string	`json:"ID"`
	Port	string	`json:"PORT"`
	Status	string	`json:"STATUS"`
}

var CiscoNxosShowVpc_Template = `Value ID (\d+)
Value PORT (\S+)
Value STATUS (\S+)

Start
  ^${ID}\s+${PORT}\s+${STATUS} -> Record

`