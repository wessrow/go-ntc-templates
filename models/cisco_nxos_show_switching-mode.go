package models

type CiscoNxosShowSwitchingMode struct {
	Configured_switching_mode	string	`json:"CONFIGURED_SWITCHING_MODE"`
	Module_number	string	`json:"MODULE_NUMBER"`
	Operational_mode	string	`json:"OPERATIONAL_MODE"`
}

var CiscoNxosShowSwitchingMode_Template = `Value Filldown CONFIGURED_SWITCHING_MODE ([\S\s]+?)
Value Required MODULE_NUMBER (\S+)
Value Required OPERATIONAL_MODE ([\S\s]+?)

Start
  ^\s*Configured\s*switching\s*mode:\s*${CONFIGURED_SWITCHING_MODE}\s*$$
  ^\s*Module\s*Number\s*Operational\s*Mode\s*$$
  ^\s*${MODULE_NUMBER}\s*${OPERATIONAL_MODE}\s*$$ -> Record
  ^\s*$$
 ^. -> Error
`