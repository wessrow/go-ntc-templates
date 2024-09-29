package alcatel_sros 

type ShowPort struct {
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

var ShowPort_Template = `Value PORT_ID (\S+)
Value ADMIN_STATE (Up|Down)
Value LINK (Yes|No)
Value PORT_STATE (Up|Down|Ghost|Link Up)
Value CFG_MTU (\d+)
Value OPER_MTU (\d+)
Value LAG (\d+|-)
Value PORT_MODE (\S+)
Value PORT_ENCP (\S+)
Value PORT_TYPE (\S+)
Value C_QS_S_XFP_MDIMDX (.*)

Start
  ^----------- -> Port

Port
  ^${PORT_ID}\s+${ADMIN_STATE}\s+${PORT_STATE}\s+conn\s*${C_QS_S_XFP_MDIMDX} -> Record
  ^${PORT_ID}\s+${ADMIN_STATE}\s+${LINK}\s+${PORT_STATE}\s+${CFG_MTU}\s+${OPER_MTU}\s+${LAG}\s+${PORT_MODE}\s+${PORT_ENCP}\s+${PORT_TYPE}\s*${C_QS_S_XFP_MDIMDX} -> Record
  ^\s*$$
  ^-----------
  ^===========
  ^Port
  ^Id
  ^\*\sindicates
  ^. -> Error
`