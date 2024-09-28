package models

type CiscoIosShowPowerSupplies struct {
	Ps_needed	string	`json:"PS_NEEDED"`
	Ps_avail	string	`json:"PS_AVAIL"`
}