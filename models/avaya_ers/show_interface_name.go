package avaya_ers 

type ShowInterfaceName struct {
	Port	string	`json:"PORT"`
	Name	string	`json:"NAME"`
}

var ShowInterfaceName_Template = `Value PORT (\S+)
Value NAME ([\S ]*)

Start
  ^Port Name|Unit\/Port Name
  ^-+ -+
  ^${PORT}\s+${NAME} -> Record

Done
`