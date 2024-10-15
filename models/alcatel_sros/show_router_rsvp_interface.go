package alcatel_sros

type ShowRouterRsvpInterface struct {
	Active_sessions string `json:"ACTIVE_SESSIONS"`
	Total_bw        string `json:"TOTAL_BW"`
	Resv_bw         string `json:"RESV_BW"`
	Admin_state     string `json:"ADMIN_STATE"`
	Oper_state      string `json:"OPER_STATE"`
	Interface       string `json:"INTERFACE"`
	Total_sessions  string `json:"TOTAL_SESSIONS"`
}

var ShowRouterRsvpInterface_Template string = `Value Required INTERFACE (\S+)
Value Required TOTAL_SESSIONS (\d+|-)
Value Required ACTIVE_SESSIONS (\d+|-)
Value Required TOTAL_BW (\d+|-)
Value Required RESV_BW (\d+|-)
Value Required ADMIN_STATE (Up|Dwn|Down)
Value Required OPER_STATE (Up|Dwn|Down)

Start
  ^=+
  ^RSVP\s+Interfaces
  ^Interface\s+Total\s+Active\s+Total\s+BW\s+Resv\s+BW\s+Adm\s+Opr\s*$$
  ^\s+Sessions\s+Sessions\s+\(Mbps\)\s+\(Mbps\)
  ^-+ -> Interface
  ^\s*$$
  ^. -> Error

Interface
  ^${INTERFACE}\s*${TOTAL_SESSIONS}\s*${ACTIVE_SESSIONS}\s*${TOTAL_BW}\s*${RESV_BW}\s*${ADMIN_STATE}\s*${OPER_STATE} -> Record
  ^Interfaces
  ^=+
  ^-+ -> Done
  ^\s*$$
  ^. -> Error

Done
`
