package models

type CiscoIosShowVrfDetail struct {
	Vrf	string	`json:"VRF"`
	Rd	string	`json:"RD"`
	Rt_import	[]string	`json:"RT_IMPORT"`
	Rt_export	[]string	`json:"RT_EXPORT"`
	Vpn_id	string	`json:"VPN_ID"`
	Description	string	`json:"DESCRIPTION"`
	Interfaces	[]string	`json:"INTERFACES"`
}