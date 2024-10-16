package alcatel_sros

type ShowRouterOspfInterface struct {
	Area           string `json:"AREA"`
	Desig_rtr      string `json:"DESIG_RTR"`
	Bkup_desig_rtr string `json:"BKUP_DESIG_RTR"`
	Admin_state    string `json:"ADMIN_STATE"`
	Oper_state     string `json:"OPER_STATE"`
	Interface      string `json:"INTERFACE"`
}

var ShowRouterOspfInterface_Template string = `Value Required INTERFACE (\S+)
Value Required AREA (\S+)
Value Required DESIG_RTR (\S+)
Value Required BKUP_DESIG_RTR (\S+)
Value Required ADMIN_STATE (\S+)
Value Required OPER_STATE (\S+)

Start
  ^=+
  ^Rtr\s+Base
  ^If\s+Name\s+Area\s+Id\s+Designated\s+Rtr\s+Bkup\s+Desig\s+Rtr\s+Adm\s+Oper\s*$$
  ^-+ -> Interface
  ^\s*$$
  ^. -> Error

Interface
  ^${INTERFACE}\s+${AREA}\s+${DESIG_RTR}\s+${BKUP_DESIG_RTR}\s+${ADMIN_STATE}\s+${OPER_STATE} -> Record
  ^-+ -> Total
  ^\s*$$
  ^. -> Error

Total
  ^No.\s+
  ^=+ -> Done
  ^. -> Error
  
Done
`
