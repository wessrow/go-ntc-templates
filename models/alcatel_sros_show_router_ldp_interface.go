package models

type AlcatelSrosShowRouterLdpInterface struct {
	Interface	string	`json:"INTERFACE"`
	Adm	string	`json:"ADM"`
	Opr	string	`json:"OPR"`
	Sub_interface	string	`json:"SUB_INTERFACE"`
	Sub_interface_adm	string	`json:"SUB_INTERFACE_ADM"`
	Sub_interface_opr	string	`json:"SUB_INTERFACE_OPR"`
	Hello_fctr	string	`json:"HELLO_FCTR"`
	Holdtime	string	`json:"HOLDTIME"`
	Ka_fctr	string	`json:"KA_FCTR"`
	Ka_time	string	`json:"KA_TIME"`
	Transport_address	string	`json:"TRANSPORT_ADDRESS"`
}

var AlcatelSrosShowRouterLdpInterface_Template = `Value INTERFACE (\S+)
Value ADM (Up|Dwn)
Value OPR (Up|Dwn)
Value SUB_INTERFACE (ipv\d)
Value SUB_INTERFACE_ADM (Up|Dwn)
Value SUB_INTERFACE_OPR (Up|Dwn)
Value HELLO_FCTR (\d+)
Value HOLDTIME (\d+)
Value KA_FCTR (\d+)
Value KA_TIME (\d+)
Value TRANSPORT_ADDRESS (\S+)

Start
  ^----------- -> Interface

Interface
  ^${INTERFACE}\s+${ADM}/${OPR}(\s|$$) -> SubInterface

SubInterface
  ^\s+${SUB_INTERFACE}\s+${SUB_INTERFACE_ADM}/${SUB_INTERFACE_OPR}\s+${HELLO_FCTR}\s+${HOLDTIME}\s+${KA_FCTR}\s+${KA_TIME}\s+${TRANSPORT_ADDRESS}(\s|$$) -> Record Interface
  ^No.\s+of\s+Interfaces
  ^-----------
  ^\s*$$
  ^. -> Error

`