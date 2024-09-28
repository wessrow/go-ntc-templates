package models

type CiscoIosShowApSummary struct {
	Ap_name	string	`json:"AP_NAME"`
	Slot	string	`json:"SLOT"`
	Ap_model	string	`json:"AP_MODEL"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Radio_mac	string	`json:"RADIO_MAC"`
	Location	string	`json:"LOCATION"`
	Country	string	`json:"COUNTRY"`
	Regulatory_domain	string	`json:"REGULATORY_DOMAIN"`
	Ip_address	string	`json:"IP_ADDRESS"`
	State	string	`json:"STATE"`
}