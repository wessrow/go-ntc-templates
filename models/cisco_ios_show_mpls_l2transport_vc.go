package models

type CiscoIosShowMplsL2transportVc struct {
	Interface	string	`json:"INTERFACE"`
	Local_circuit	string	`json:"LOCAL_CIRCUIT"`
	Dest_address	string	`json:"DEST_ADDRESS"`
	Vc_id	string	`json:"VC_ID"`
	Status	string	`json:"STATUS"`
}

var CiscoIosShowMplsL2transportVc_Template = `Value INTERFACE ([a-zA-Z0-9\-/.]+)
Value LOCAL_CIRCUIT (.*)
Value DEST_ADDRESS (\S+)
Value VC_ID (\d+)
Value STATUS (.*)

Start
  ^\s*Local\s+intf\s+Local\s+circuit\s+Dest\s+address\s+VC\s+ID\s+Status\s*$$ -> L2transportTable
  ^\s*$$
  ^. -> Error

L2transportTable
  ^\s*-+(\s+-+)*\s*$$
  ^\s*${INTERFACE}\s+${LOCAL_CIRCUIT}\s+${DEST_ADDRESS}\s+${VC_ID}\s+${STATUS}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`