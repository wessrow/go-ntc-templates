package models

type AlcatelSrosShowRouterMplsInterface struct {
	Interface	string	`json:"INTERFACE"`
	Port	string	`json:"PORT"`
	Admin_status	string	`json:"ADMIN_STATUS"`
	Oper_status_v4	string	`json:"OPER_STATUS_V4"`
	Oper_status_v6	string	`json:"OPER_STATUS_V6"`
	Te_metric	string	`json:"TE_METRIC"`
}