package ciena_saos

type RstpShow struct {
	Port_stp       string `json:"PORT_STP"`
	Rstp_st        string `json:"RSTP_ST"`
	Port_uptime    string `json:"PORT_UPTIME"`
	Port_op        string `json:"PORT_OP"`
	Stp_st         string `json:"STP_ST"`
	Path_cost_oper string `json:"PATH_COST_OPER"`
	Path_cost_d    string `json:"PATH_COST_D"`
	Edgep_op       string `json:"EDGEP_OP"`
	P2p_mac_adm    string `json:"P2P_MAC_ADM"`
	P2p_mac_op     string `json:"P2P_MAC_OP"`
	Name           string `json:"NAME"`
	Domain_adm     string `json:"DOMAIN_ADM"`
	Domain_ef      string `json:"DOMAIN_EF"`
	Port_role      string `json:"PORT_ROLE"`
	Pr             string `json:"PR"`
	Edgep_ad       string `json:"EDGEP_AD"`
}

var RstpShow_Template string = `Value NAME ([0-9A-Za-z\.]+)
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
  ^. -> Error`
