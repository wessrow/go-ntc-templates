package models

type CiscoNxosShowCtsInterfaceAll struct {
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