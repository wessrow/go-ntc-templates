package models

type AlcatelSrosShowServiceIdInterface struct {
	Interface_name	string	`json:"INTERFACE_NAME"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Oper_state_v4	string	`json:"OPER_STATE_V4"`
	Oper_state_v6	string	`json:"OPER_STATE_V6"`
	Type	string	`json:"TYPE"`
	Port_sap_id	string	`json:"PORT_SAP_ID"`
	Ip_address	[]string	`json:"IP_ADDRESS"`
	Pfx_state	[]string	`json:"PFX_STATE"`
}

var AlcatelSrosShowServiceIdInterface_Template = `Value Required INTERFACE_NAME (\S+)
Value Required ADMIN_STATE (Up|Down)
Value Required OPER_STATE_V4 (Up|Down)
Value Required OPER_STATE_V6 (Up|Down)
Value Required TYPE (\S+)
Value Required PORT_SAP_ID (\S+)
Value List IP_ADDRESS (\S+)
Value List PFX_STATE (\S+)

Start
  ^=+
  ^Interface\s+Table
  ^Interface-Name\s+Adm\s+Opr\(v4\/v6\)\s+Type\s+Port/SapId\s*$$ -> Interface
  ^\s*$$
  ^. -> Error

Interface
  ^\s+IP-Address\s+PfxState
  ^-+
  ^.*?\s+(Up|Down) -> Continue.Record
  ^${INTERFACE_NAME}\s+${ADMIN_STATE}\s+${OPER_STATE_V4}\/${OPER_STATE_V6}\s+${TYPE}\s+${PORT_SAP_ID} 
  ^\s+${IP_ADDRESS}\s+${PFX_STATE}
  ^Interfaces\s*:\s+\d+\s*$$
  ^=+
  ^\*\s+indicates
  ^\s*$$
  ^. -> Error

`