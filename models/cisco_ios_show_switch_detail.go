package models

type CiscoIosShowSwitchDetail struct {
	Switch	string	`json:"SWITCH"`
	Role	string	`json:"ROLE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Priority	string	`json:"PRIORITY"`
	Version	string	`json:"VERSION"`
	State	string	`json:"STATE"`
}