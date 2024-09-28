package models

type AlcatelSrosShowPort struct {
	Port_id	string	`json:"PORT_ID"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Link	string	`json:"LINK"`
	Port_state	string	`json:"PORT_STATE"`
	Cfg_mtu	string	`json:"CFG_MTU"`
	Oper_mtu	string	`json:"OPER_MTU"`
	Lag	string	`json:"LAG"`
	Port_mode	string	`json:"PORT_MODE"`
	Port_encp	string	`json:"PORT_ENCP"`
	Port_type	string	`json:"PORT_TYPE"`
	C_qs_s_xfp_mdimdx	string	`json:"C_QS_S_XFP_MDIMDX"`
}