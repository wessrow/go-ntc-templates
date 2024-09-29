package models

type CiscoAsaShowInventory struct {
	Name	string	`json:"NAME"`
	Descr	string	`json:"DESCR"`
	Pid	string	`json:"PID"`
	Vid	string	`json:"VID"`
	Sn	string	`json:"SN"`
}

var CiscoAsaShowInventory_Template = `Value NAME (.*)
Value DESCR (.*)
Value PID ([\w\d\s/-]+[\S]+)
Value VID (\S+)
Value SN (\S+)

Start
  ^Name:\s+"${NAME}"\s*,\s+DESCR:\s+"${DESCR}"
  ^PID:\s+${PID}\s*,\s+VID:\s+${VID}\s*,\s+SN:\s+${SN} -> Record
  ^PID:\s+${PID}\s*,\s+VID:\s+${VID}\s*,\s+SN: -> Record
  ^\s*$$
  ^show_inventory_all\s+\S+ -> NoRecord
  ^.+ -> Error

`