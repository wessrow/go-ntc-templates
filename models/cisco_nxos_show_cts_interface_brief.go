package models

type CiscoNxosShowCtsInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Mode	string	`json:"MODE"`
	Ifc_state	string	`json:"IFC_STATE"`
	Sgt_assignment	string	`json:"SGT_ASSIGNMENT"`
	Sgt_propagate	string	`json:"SGT_PROPAGATE"`
}

var CiscoNxosShowCtsInterfaceBrief_Template = `Value INTERFACE (\S+)
Value MODE (\S+)
Value IFC_STATE (\S+)
Value SGT_ASSIGNMENT ((\S+(\s\S+)*))
Value SGT_PROPAGATE (Enabled|Disabled)

Start
  ^${INTERFACE}\s+${MODE}\s+${IFC_STATE}\s+${SGT_ASSIGNMENT}\s+${SGT_PROPAGATE} -> Record
`