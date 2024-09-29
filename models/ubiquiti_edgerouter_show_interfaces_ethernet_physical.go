package models

type UbiquitiEdgerouterShowInterfacesEthernetPhysical struct {
	Interface	string	`json:"INTERFACE"`
	Auto_negotiation	string	`json:"AUTO_NEGOTIATION"`
	Speed	string	`json:"SPEED"`
	Duplex	string	`json:"DUPLEX"`
	Link_detected	string	`json:"LINK_DETECTED"`
}

var UbiquitiEdgerouterShowInterfacesEthernetPhysical_Template = `Value INTERFACE (\S+)
Value AUTO_NEGOTIATION (\S+)
Value SPEED (\S+)
Value DUPLEX (\S+)
Value LINK_DETECTED (\S+)

Start
  ^Settings for ${INTERFACE}:
  ^\s+Auto-negotiation:\s${AUTO_NEGOTIATION}
  ^\s+Speed:\s${SPEED}
  ^\s+Duplex:\s${DUPLEX}
  ^\s+Link\sdetected:\s${LINK_DETECTED}
  ^\s*$$
  ^. -> Error

`