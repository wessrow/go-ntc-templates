package alcatel_sros 

type ShowRouterPimInterface struct {
	Interface	string	`json:"INTERFACE"`
	Adm	string	`json:"ADM"`
	Opr	string	`json:"OPR"`
	Dr	string	`json:"DR"`
	Dr_prty	string	`json:"DR_PRTY"`
	Hello_intvl	string	`json:"HELLO_INTVL"`
	Mcast_send	string	`json:"MCAST_SEND"`
}

var ShowRouterPimInterface_Template = `Value INTERFACE (\S+)
Value ADM (Up|Down)
Value OPR (Up|Down)
Value DR ((\d+\.\d+\.\d+\.\d+|[0-9a-f:]*)|N/A)
Value DR_PRTY (\d+|N/A)
Value HELLO_INTVL (\d+)
Value MCAST_SEND (\S+)

Start
  ^----------- -> Interface

Interface
  ^${INTERFACE}\s+${ADM}\s+${OPR}\s+${DR_PRTY}\s+${HELLO_INTVL}\s+${MCAST_SEND}(\s|$$) -> Dr

Dr
  ^\s+${DR}(\s|$$) -> Record Interface
  ^No.\s+of\s+Interfaces
  ^-----------
  ^\s*$$
  ^. -> Error
`