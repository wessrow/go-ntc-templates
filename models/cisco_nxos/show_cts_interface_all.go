package cisco_nxos 

type ShowCtsInterfaceAll struct {
	Interface	string	`json:"INTERFACE"`
	Cts_status	string	`json:"CTS_STATUS"`
	Cts_mode	string	`json:"CTS_MODE"`
	Ifc_state	string	`json:"IFC_STATE"`
	Authentication_status	string	`json:"AUTHENTICATION_STATUS"`
	Peer_identity	string	`json:"PEER_IDENTITY"`
	Peer_is	string	`json:"PEER_IS"`
	Dot1x_role	string	`json:"DOT1X_ROLE"`
	Authorization_status	string	`json:"AUTHORIZATION_STATUS"`
	Peer_sgt	string	`json:"PEER_SGT"`
	Peer_sgt_assignment	string	`json:"PEER_SGT_ASSIGNMENT"`
	Sap_status	string	`json:"SAP_STATUS"`
	Sgt_propagate	string	`json:"SGT_PROPAGATE"`
}

var ShowCtsInterfaceAll_Template = `Value Filldown INTERFACE (\S+)
Value Required CTS_STATUS (\w+)
Value CTS_MODE (\S+)
Value IFC_STATE (\S+)
Value AUTHENTICATION_STATUS (\S+)
Value PEER_IDENTITY ((\S+(\s\S+)*))
Value PEER_IS ((\S+(\s\S+)*))
Value DOT1X_ROLE (\S+)
Value AUTHORIZATION_STATUS (\S+)
Value PEER_SGT (\d+)
Value PEER_SGT_ASSIGNMENT ((\S+(\s\S+)*))
Value SAP_STATUS (\S+)
Value SGT_PROPAGATE (\S+)

Start
  ^CTS\s+Information\s+for\s+Interface\s+${INTERFACE}: -> Interface

Interface
  ^\s+CTS\s+is\s+${CTS_STATUS},\s+mode:\s+${CTS_MODE}
  ^\s+IFC\s+state:\s+${IFC_STATE}
  ^\s+Authentication\s+Status:\s+${AUTHENTICATION_STATUS}
  ^\s+Peer\s+Identity:\s+${PEER_IDENTITY}
  ^\s+Peer\s+is:\s+${PEER_IS}
  ^\s+802.1X\s+role:\s+${DOT1X_ROLE}
  ^\s+Authorization\s+Status:\s+${AUTHORIZATION_STATUS}
  ^\s+PEER\s+SGT:\s+${PEER_SGT}
  ^\s+Peer\s+SGT\s+assignment:\s+${PEER_SGT_ASSIGNMENT}
  ^\s+SAP\s+Status:\s+${SAP_STATUS}
  ^\s+Propagate\s+SGT:\s+${SGT_PROPAGATE} -> Record Start

`