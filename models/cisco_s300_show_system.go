package models

type CiscoS300ShowSystem struct {
	Description	string	`json:"DESCRIPTION"`
	Up_time	string	`json:"UP_TIME"`
	Contact	string	`json:"CONTACT"`
	Hostname	string	`json:"HOSTNAME"`
	Location	[]string	`json:"LOCATION"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Oid	string	`json:"OID"`
	Fans_status	string	`json:"FANS_STATUS"`
}