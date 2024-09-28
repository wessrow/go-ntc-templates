package models

type ArubaOsShowApRadioDatabase struct {
	Ap_name	string	`json:"AP_NAME"`
	Group	string	`json:"GROUP"`
	Ap_model	string	`json:"AP_MODEL"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Flags	string	`json:"FLAGS"`
	Primary	string	`json:"PRIMARY"`
	Standby	string	`json:"STANDBY"`
	Radio0	string	`json:"RADIO0"`
	Radio1	string	`json:"RADIO1"`
	Radio2	string	`json:"RADIO2"`
}