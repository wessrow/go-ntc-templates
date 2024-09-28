package models

type AlcatelSrosShowRouterRsvpInterface struct {
	Interface	string	`json:"INTERFACE"`
	Total_sessions	string	`json:"TOTAL_SESSIONS"`
	Active_sessions	string	`json:"ACTIVE_SESSIONS"`
	Total_bw	string	`json:"TOTAL_BW"`
	Resv_bw	string	`json:"RESV_BW"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Oper_state	string	`json:"OPER_STATE"`
}