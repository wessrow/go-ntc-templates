package models

type CiscoIosShowRepTopology struct {
	Switch	string	`json:"SWITCH"`
	Interface	string	`json:"INTERFACE"`
	Edge	string	`json:"EDGE"`
	Role	string	`json:"ROLE"`
}

var CiscoIosShowRepTopology_Template = `Value SWITCH (\S+)
Value INTERFACE (\S+)
Value EDGE (Pri|Sec)
Value ROLE (Alt|Open)

Start
  # HEADER CHECK
  ^REP\s+Segment\s+\d+$$
  ^BridgeName\s+PortName\s+Edge\s+Role$$
  ^(-+(\s+)?)+$$ -> Data
  ^. -> Error

Data
  # OUTPUT
  ^${SWITCH}\s+${INTERFACE}\s+${EDGE}\s+${ROLE}$$ -> Record
  ^\s*$$
  ^. -> Error
`