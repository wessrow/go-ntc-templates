package cisco_ios 

type ShowInterfacesStatus struct {
	Port	string	`json:"PORT"`
	Name	string	`json:"NAME"`
	Status	string	`json:"STATUS"`
	Vlan_id	string	`json:"VLAN_ID"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Type	string	`json:"TYPE"`
	Fc_mode	string	`json:"FC_MODE"`
}

var ShowInterfacesStatus_Template = `Value PORT (\S+)
Value NAME (.+?)
Value STATUS (err-disabled|disabled|connected|notconnect|inactive|up|down|monitoring|suspended)
Value VLAN_ID (\d+|trunk|routed|unassigned)
Value DUPLEX (\S+)
Value SPEED (\S+)
Value TYPE (.*)
Value FC_MODE (\S+)

Start
  ^Load\s+for\s+
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Time\s+source\s+is
  ^-+\s*$$
  ^Port\s+Name\s+Status\s+Vlan\s+Duplex\s+Speed\s+Type -> Interfaces
  ^\s*$$
  ^. -> Error

Interfaces
  #Match fc...
  ^\s*${PORT}\s+is\s+${STATUS}\s+Port\s+mode\s+is\s+${FC_MODE}\s*$$ -> Record
  ^\s*${PORT}\s+is\s+${STATUS}\s+\(${TYPE}\)\s*$$ -> Record
  ^\s*${PORT}\s+${STATUS}\s+${VLAN_ID}\s+${DUPLEX}\s+${SPEED}\s*${TYPE}$$ -> Record
  ^\s*${PORT}\s+${NAME}\s+${STATUS}:?\s+${VLAN_ID}\s+${DUPLEX}\s+${SPEED}\s*${TYPE}\s*$$ -> Record
  ^-+
  ^\s*$$
  ^. -> Error
`