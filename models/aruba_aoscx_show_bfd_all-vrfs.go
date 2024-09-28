package models

type ArubaAoscxShowBfdAllVrfs struct {
	Admin_status	string	`json:"ADMIN_STATUS"`
	Src_ip	string	`json:"SRC_IP"`
	Session	string	`json:"SESSION"`
	Interface	string	`json:"INTERFACE"`
	Vrf	string	`json:"VRF"`
	Source_ip	string	`json:"SOURCE_IP"`
	Destination_ip	string	`json:"DESTINATION_IP"`
	Echo	string	`json:"ECHO"`
	State	string	`json:"STATE"`
	Application	string	`json:"APPLICATION"`
}