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

var CienaSaosRstpShow_Template = `Value NAME ([0-9A-Za-z\.]+)
Value PORT_OP (.{2}|\w+)
Value PORT_STP (\w+)
Value PORT_ROLE (\w+)
Value RSTP_ST (\w+)
Value STP_ST (\w+)
Value PR (\w+)
Value PATH_COST_OPER (\w+|.{6})
Value PATH_COST_D (\w+)
Value EDGEP_AD (\w+)
Value EDGEP_OP (\w+)
Value P2P_MAC_ADM (\w+)
Value P2P_MAC_OP (\w+)
Value DOMAIN_ADM (\w+)
Value DOMAIN_EF (\w+)
Value PORT_UPTIME (.{13})


Start
  ^\+---------------------Bridge-------------------\+BridgeTimer\+-------Bridge------\+
  ^\+-.*
  ^\|Rst.*
  ^\|Adm.*
  ^\+-.*
  ^\|\s*ENA.*
  ^\+-.*
  ^\|Rst.*
  ^\|Mode.*
  ^\+-.*
  ^\|\s*RSTP.*
  ^\+-.*
  ^\+-.*
  ^\+-.*
  ^\|\s*Port.*
  ^\|\s*Name.*
  ^\+-.*
  ^\|${NAME}\s*\|${PORT_OP}\|${PORT_STP}\|\s*${PORT_ROLE}\|${RSTP_ST}\s*\|\s*${STP_ST}\s*\|\s*${PR}\|\s*${PATH_COST_OPER}\s*\|${PATH_COST_D}\s*\|\s*${EDGEP_AD}\|\s*${EDGEP_OP}\|${P2P_MAC_ADM}\|\s*${P2P_MAC_OP}\s*\|\s*${DOMAIN_ADM}\|\s*${DOMAIN_EF}\|${PORT_UPTIME}\| -> Record
  ^\+-.*
  ^\s*$$
  ^. -> Error
`