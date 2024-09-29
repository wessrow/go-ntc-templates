package models

type AlcatelSrosShowRouterMplsInterface struct {
	Interface	string	`json:"INTERFACE"`
	Port	string	`json:"PORT"`
	Admin_status	string	`json:"ADMIN_STATUS"`
	Oper_status_v4	string	`json:"OPER_STATUS_V4"`
	Oper_status_v6	string	`json:"OPER_STATUS_V6"`
	Te_metric	string	`json:"TE_METRIC"`
}

var AlcatelSrosShowRouterMplsInterface_Template = `Value Required INTERFACE (\S+)
Value Required PORT (\S+)
Value Required ADMIN_STATUS (Up|Down)
Value Required OPER_STATUS_V4 (Up|Down)
Value Required OPER_STATUS_V6 (Up|Down)
Value Required TE_METRIC (\d+|None)


Start
  ^=+
  ^MPLS Interfaces
  ^Interface\s+Port-id\s+Adm\s+Opr\(V4\/V6\)\s+TE- -> Interface
  ^\s+metric
  ^-+
  ^\s*$$
  ^. -> Error

Interface
  ^${INTERFACE}\s+${PORT}\s+${ADMIN_STATUS}\s+${OPER_STATUS_V4}\S${OPER_STATUS_V6}\s+${TE_METRIC} -> Record
  ^\s+Admin\s+Groups
  ^\s+\S+
  ^\s+SRLG\s+Groups
  ^Interfaces
  ^=+
  ^\s+metric
  ^-+
  ^\s*$$
  ^. -> Error
`