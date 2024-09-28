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