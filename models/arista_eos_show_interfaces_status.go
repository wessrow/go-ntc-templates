package models

type AristaEosShowInterfacesStatus struct {
	Port	string	`json:"PORT"`
	Name	string	`json:"NAME"`
	Status	string	`json:"STATUS"`
	Vlan_id	string	`json:"VLAN_ID"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Type	string	`json:"TYPE"`
}

var AristaEosShowInterfacesStatus_Template = `Value PORT (\S+)
Value NAME (\S.*(?<!\s))
Value STATUS (connected|notconnect|inactive|errdisabled|disabled)
Value VLAN_ID (\S+)
Value DUPLEX (\S+)
Value SPEED (\S+)
Value TYPE (.+?)

Start
  ^${PORT}\s+${NAME}\s+${STATUS}\s+(?:in\s+)?${VLAN_ID}\s+${DUPLEX}\s+${SPEED}\s+${TYPE}\s*$$ -> Record
  ^${PORT}\s+${STATUS}\s+(?:in\s+)?${VLAN_ID}\s+${DUPLEX}\s+${SPEED}\s+${TYPE}\s*$$ -> Record
  ^\s*
  ^. -> Error

`