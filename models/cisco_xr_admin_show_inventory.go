package models

type CiscoXrAdminShowInventory struct {
	Name	string	`json:"NAME"`
	Descr	string	`json:"DESCR"`
	Pid	string	`json:"PID"`
	Vid	string	`json:"VID"`
	Sn	string	`json:"SN"`
}

var CiscoXrAdminShowInventory_Template = `Value Required NAME (.*?)
Value DESCR (.*?)
Value PID (\S+)
Value VID (\S+)
Value SN (\S*)

Start
  ^NAME:\s+"${NAME}",\s+DESCR:\s+"${DESCR}"\s*$$
  ^\s*Name:\s+${NAME}\s+Descr:\s+${DESCR}\s*$$
  ^\s*PID:\s+${PID}\s*,\s+VID:\s+${VID}\s*,\s+SN:\s*${SN}\s*$$ -> Record
  ^\s*PID:\s+${PID}\s+(,\s+)?VID:\s+${VID}\s+(,\s+)?SN:\s+${SN}\s*$$ -> Record
  #
  ^\s*$$
  ^(Mon|Tue|Wed|Thu|Fri|Sat|Sun)\s\S{3}\s+\d
  ^. -> Error

`