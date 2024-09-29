package cisco_nxos 

type ShowFex struct {
	Number	string	`json:"NUMBER"`
	Descr	string	`json:"DESCR"`
	State	string	`json:"STATE"`
	Model	string	`json:"MODEL"`
	Serial	string	`json:"SERIAL"`
}

var ShowFex_Template = `Value NUMBER (\d+)
Value DESCR ((\S+)(\s+\S+)*)
Value STATE (\S+)
Value MODEL (\S+)
Value SERIAL (\S+)

Start
  ^${NUMBER}\s+${DESCR}\s+${STATE}\s+${MODEL}\s+${SERIAL} -> Record
`