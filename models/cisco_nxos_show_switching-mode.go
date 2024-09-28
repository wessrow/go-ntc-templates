package models

type CiscoNxosShowSwitchingMode struct {
	Configured_switching_mode	string	`json:"CONFIGURED_SWITCHING_MODE"`
	Module_number	string	`json:"MODULE_NUMBER"`
	Operational_mode	string	`json:"OPERATIONAL_MODE"`
}