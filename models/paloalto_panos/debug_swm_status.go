package paloalto_panos 

type DebugSwmStatus struct {
	Part	string	`json:"PART"`
	State	string	`json:"STATE"`
	Version	string	`json:"VERSION"`
}

var DebugSwmStatus_Template = `Value PART (\S+)
Value STATE (\S+)
Value VERSION (\S+)

Start
  ^Partition\s+State\s+Version
  ^-+$$
  ^${PART}\s+${STATE}\s+${VERSION} -> Record
  ^\s*$$
  ^. -> Error
`