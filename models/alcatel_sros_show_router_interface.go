package models

type AlcatelSrosShowRouterInterface struct {
	Interface	string	`json:"INTERFACE"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Oper_state_v4	string	`json:"OPER_STATE_V4"`
	Oper_state_v6	string	`json:"OPER_STATE_V6"`
	Mode	string	`json:"MODE"`
	Port_sap_id	string	`json:"PORT_SAP_ID"`
	Ip_address	[]string	`json:"IP_ADDRESS"`
	Pfx_state	[]string	`json:"PFX_STATE"`
}

var AlcatelSrosShowRouterInterface_Template = `Value Required INTERFACE (\S+)
Value Required ADMIN_STATE (Up|Down)
Value Required OPER_STATE_V4 (Up|Down)
Value Required OPER_STATE_V6 (Up|Down)
Value Required MODE (\S+)
Value Required PORT_SAP_ID (\S+)
Value List IP_ADDRESS (\S+)
Value List PFX_STATE (\S+)

Start
  ^=+
  ^Interface\s+Table
  ^Interface-Name\s+Adm\s+Opr\(v4\/v6\)\s+Mode\s+Port/SapId\s*$$ -> Interface
  ^\s*$$
  ^. -> Error

Interface
  ^\s+IP-Address\s+PfxState
  ^-+
  ^.*?\s+(Up|Down) -> Continue.Record
  ^${INTERFACE}\s+${ADMIN_STATE}\s+${OPER_STATE_V4}\/${OPER_STATE_V6}\s+${MODE}\s+${PORT_SAP_ID} 
  ^\s+${IP_ADDRESS}\s+${PFX_STATE}
  ^Interfaces -> Done
  ^=+
  ^\s*$$
  ^. -> Error

Done

`