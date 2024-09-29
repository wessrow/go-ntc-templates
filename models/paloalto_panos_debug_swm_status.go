package models

type PaloaltoPanosDebugSwmStatus struct {
	Part	string	`json:"PART"`
	State	string	`json:"STATE"`
	Version	string	`json:"VERSION"`
}

var PaloaltoPanosDebugSwmStatus_Template = `Value PART (\S+)
Value STATE (\S+)
Value VERSION (\S+)

Start
  ^Partition\s+State\s+Version
  ^-+$$
  ^${PART}\s+${STATE}\s+${VERSION} -> Record
  ^\s*$$
  ^. -> Error

`