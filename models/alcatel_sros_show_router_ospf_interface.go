package models

type AlcatelSrosShowRouterOspfInterface struct {
	Interface	string	`json:"INTERFACE"`
	Area	string	`json:"AREA"`
	Desig_rtr	string	`json:"DESIG_RTR"`
	Bkup_desig_rtr	string	`json:"BKUP_DESIG_RTR"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Oper_state	string	`json:"OPER_STATE"`
}