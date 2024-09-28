package models

type CiscoIosShowMplsL2transportVc struct {
	Interface	string	`json:"INTERFACE"`
	Local_circuit	string	`json:"LOCAL_CIRCUIT"`
	Dest_address	string	`json:"DEST_ADDRESS"`
	Vc_id	string	`json:"VC_ID"`
	Status	string	`json:"STATUS"`
}