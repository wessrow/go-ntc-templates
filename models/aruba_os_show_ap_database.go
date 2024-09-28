package models

type ArubaOsShowApDatabase struct {
	Ap_name	string	`json:"AP_NAME"`
	Group	string	`json:"GROUP"`
	Ap_model	string	`json:"AP_MODEL"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Flags	string	`json:"FLAGS"`
	Primary	string	`json:"PRIMARY"`
	Standby	string	`json:"STANDBY"`
}