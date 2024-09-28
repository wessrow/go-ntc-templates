package models

type CiscoIosShowModuleSubmodule struct {
	Module	string	`json:"MODULE"`
	Submodule	string	`json:"SUBMODULE"`
	Submodule_model	string	`json:"SUBMODULE_MODEL"`
	Submodule_serial	string	`json:"SUBMODULE_SERIAL"`
	Submodule_hw	string	`json:"SUBMODULE_HW"`
	Submodule_status	string	`json:"SUBMODULE_STATUS"`
}