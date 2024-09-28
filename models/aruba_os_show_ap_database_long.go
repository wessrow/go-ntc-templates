package models

type ArubaOsShowApDatabaseLong struct {
	Ap_name	string	`json:"AP_NAME"`
	Group	string	`json:"GROUP"`
	Ap_model	string	`json:"AP_MODEL"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Flags	string	`json:"FLAGS"`
	Primary	string	`json:"PRIMARY"`
	Standby	string	`json:"STANDBY"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Serial	string	`json:"SERIAL"`
	Port	string	`json:"PORT"`
	Fqln	string	`json:"FQLN"`
	Outer_ip	string	`json:"OUTER_IP"`
	User	string	`json:"USER"`
}