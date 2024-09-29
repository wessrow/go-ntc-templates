package models

type CiscoXrShowInterfacesDescription struct {
	Interface	string	`json:"INTERFACE"`
	Status	string	`json:"STATUS"`
	Protocol	string	`json:"PROTOCOL"`
	Description	string	`json:"DESCRIPTION"`
}

var CiscoXrShowInterfacesDescription_Template = `Value INTERFACE (\S+)
Value STATUS (up|down|deleted|admin\-down|reset)
Value PROTOCOL (up|down|deleted|admin\-down)
Value DESCRIPTION (.+?)

Start
  ^${INTERFACE}\s+${STATUS}\s+${PROTOCOL}\s*(?:${DESCRIPTION})?\s*$$ -> Record
  ^\S+\s+\S+\s+\d+
  ^Interface\s+Status\s+Protocol\s+Description\s*$$
  ^-+\s*$$
  ^\s*$$
  ^. -> Error

`