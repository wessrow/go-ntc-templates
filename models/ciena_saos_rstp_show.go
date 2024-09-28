package models

type CienaSaosRstpShow struct {
	Name	string	`json:"NAME"`
	Port_op	string	`json:"PORT_OP"`
	Port_stp	string	`json:"PORT_STP"`
	Port_role	string	`json:"PORT_ROLE"`
	Rstp_st	string	`json:"RSTP_ST"`
	Stp_st	string	`json:"STP_ST"`
	Pr	string	`json:"PR"`
	Path_cost_oper	string	`json:"PATH_COST_OPER"`
	Path_cost_d	string	`json:"PATH_COST_D"`
	Edgep_ad	string	`json:"EDGEP_AD"`
	Edgep_op	string	`json:"EDGEP_OP"`
	P2p_mac_adm	string	`json:"P2P_MAC_ADM"`
	P2p_mac_op	string	`json:"P2P_MAC_OP"`
	Domain_adm	string	`json:"DOMAIN_ADM"`
	Domain_ef	string	`json:"DOMAIN_EF"`
	Port_uptime	string	`json:"PORT_UPTIME"`
}