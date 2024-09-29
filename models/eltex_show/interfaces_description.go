package eltex_show 

type InterfacesDescription struct {
	Interface	string	`json:"INTERFACE"`
	Interface_mode	string	`json:"INTERFACE_MODE"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Link_state	string	`json:"LINK_STATE"`
	Description	[]string	`json:"DESCRIPTION"`
}

var InterfacesDescription_Template = `Value INTERFACE ((?!Port|Ch|-)\S+)
Value INTERFACE_MODE (Trunk|Access(?:\s*\(\d+\))?|Customer(?:\s*\(\d+\))?)
Value ADMIN_STATE (Up|Down)
Value LINK_STATE (Up|Down|Not Present)
Value List DESCRIPTION (.*?)

Start
  ^\s*Admin\s+Link\s*$$
  ^\s*Port\s+Description\s*$$ -> Port2
  ^\s*Port\s+State\s+State\s+Description\s*$$ -> Port4
  ^\s*Port\s+Port\s+Mode\s+\(VLAN\)\s+State\s+State\s+Description\s*$$ -> Port5
  ^\s*Ch\s+Description\s*$$ -> Ch2
  ^\s*Ch\s+State\s+State\s+Description\s*$$ -> Ch4
  ^\s*Ch\s+Port\s+Mode\s+\(VLAN\)\s+State\s+State\s+Description\s*$$ -> Ch5
  ^\s*Vlan\s+Admin\s+State\s+Link\s+State\s+Description\s*$$ -> Vlan
  ^\s*Loopback\s+State\s+State\s+Description\s*$$ -> Loopback
  ^\s*Oob-eth\s+State\s+State\s+Description\s*$$ -> Oob
  ^\s*$$
  ^. -> Error

Port2
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Port4
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}\s+${ADMIN_STATE}\s+${LINK_STATE}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Port5
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}(?:\s+${INTERFACE_MODE})?\s+${ADMIN_STATE}\s+${LINK_STATE}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Ch2
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Ch4
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}\s+${ADMIN_STATE}\s+${LINK_STATE}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Ch5
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}(?:\s+${INTERFACE_MODE})?\s+${ADMIN_STATE}\s+${LINK_STATE}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Vlan
  ^(?:\s*-+)+\s*$$
  ^(?:\s*(?!Port|Ch|-)\S+\s+(?:Up|Down).*|\s*)$$ -> Continue.Record
  ^\s*$$ -> Start
  ^\s*${INTERFACE}\s+${ADMIN_STATE}\s+${LINK_STATE}(?:\s+${DESCRIPTION})?\s*$$
  ^\s*${DESCRIPTION}\s*$$
  ^. -> Error

Loopback
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}\s+${ADMIN_STATE}\s+${LINK_STATE}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Oob
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}\s+${ADMIN_STATE}\s+${LINK_STATE}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error
`