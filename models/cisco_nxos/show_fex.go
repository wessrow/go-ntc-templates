package cisco_nxos

type ShowFex struct {
	Model  string `json:"MODEL"`
	Serial string `json:"SERIAL"`
	Number string `json:"NUMBER"`
	Descr  string `json:"DESCR"`
	State  string `json:"STATE"`
}

var ShowFex_Template string = `Value NUMBER (\d+)
Value DESCR ((\S+)(\s+\S+)*)
Value STATE (\S+)
Value MODEL (\S+)
Value SERIAL (\S+)

Start
  ^${NUMBER}\s+${DESCR}\s+${STATE}\s+${MODEL}\s+${SERIAL} -> Record
`
