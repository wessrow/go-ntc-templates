package cisco_nxos 

type ShowFeature struct {
	Name	string	`json:"NAME"`
	Instance	string	`json:"INSTANCE"`
	State	string	`json:"STATE"`
}

var ShowFeature_Template = `Value NAME (\S+)
Value INSTANCE (\d+)
Value STATE (disabled|enabled\s\(not-running\)|enabled)

Start
  ^${NAME}\s+${INSTANCE}\s+${STATE} -> Record
`