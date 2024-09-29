package cisco_nxos 

type ShowVpc struct {
	Id	string	`json:"ID"`
	Port	string	`json:"PORT"`
	Status	string	`json:"STATUS"`
}

var ShowVpc_Template = `Value ID (\d+)
Value PORT (\S+)
Value STATUS (\S+)

Start
  ^${ID}\s+${PORT}\s+${STATUS} -> Record
`