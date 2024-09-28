package models

type CiscoAsaShowModuleDetails struct {
	Module	string	`json:"MODULE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Mod_hw	string	`json:"MOD_HW"`
	Mod_fw	string	`json:"MOD_FW"`
	Mod_sw	string	`json:"MOD_SW"`
}