package models

type CiscoIosShowSwitchDetailStackPorts struct {
	Switch	string	`json:"SWITCH"`
	Status1	string	`json:"STATUS1"`
	Status2	string	`json:"STATUS2"`
	Neighbor1	string	`json:"NEIGHBOR1"`
	Neighbor2	string	`json:"NEIGHBOR2"`
}

var CiscoIosShowSwitchDetailStackPorts_Template = `Value Key SWITCH (\d+)
Value STATUS1 (\w+)
Value STATUS2 (\w+)
Value NEIGHBOR1 (\w+)
Value NEIGHBOR2 (\w+)


Start
  ^\s*Stack\s+Port\s+Status\s+Neighbors\s*$$ -> Stack

Stack
  ^\s*Switch#\s+Port\s+1\s+Port\s+2\s+Port\s+1\s+Port\s+2\s*$$
  ^-+$$
  ^\s*${SWITCH}\s+${STATUS1}\s+${STATUS2}\s+${NEIGHBOR1}\s+${NEIGHBOR2}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`