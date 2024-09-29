package models

type AvayaErsShowInterfaceName struct {
	Port	string	`json:"PORT"`
	Name	string	`json:"NAME"`
}

var AvayaErsShowInterfaceName_Template = `Value PORT (\S+)
Value NAME ([\S ]*)

Start
  ^Port Name|Unit\/Port Name
  ^-+ -+
  ^${PORT}\s+${NAME} -> Record

Done

`