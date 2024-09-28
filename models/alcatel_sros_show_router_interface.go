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