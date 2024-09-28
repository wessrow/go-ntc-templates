package models

type CiscoNxosShowCtsInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Mode	string	`json:"MODE"`
	Ifc_state	string	`json:"IFC_STATE"`
	Sgt_assignment	string	`json:"SGT_ASSIGNMENT"`
	Sgt_propagate	string	`json:"SGT_PROPAGATE"`
}