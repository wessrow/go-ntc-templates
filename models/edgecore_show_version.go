package models

type EdgecoreShowVersion struct {
	Unit	string	`json:"UNIT"`
	Serial	string	`json:"SERIAL"`
	Service_tag	string	`json:"SERVICE_TAG"`
	Hardware_version	string	`json:"HARDWARE_VERSION"`
	Chip_id	string	`json:"CHIP_ID"`
	Epld_version	string	`json:"EPLD_VERSION"`
	Num_of_ports	string	`json:"NUM_OF_PORTS"`
	Power_status	string	`json:"POWER_STATUS"`
	Redundant_power_status	string	`json:"REDUNDANT_POWER_STATUS"`
	Loader_version	string	`json:"LOADER_VERSION"`
	Rom_version	string	`json:"ROM_VERSION"`
	Operation_code_version	string	`json:"OPERATION_CODE_VERSION"`
}

var EdgecoreShowVersion_Template = `Value UNIT (\d+)
Value SERIAL (.*)
Value SERVICE_TAG (.*)
Value HARDWARE_VERSION (.*)
Value CHIP_ID (.*)
Value EPLD_VERSION (.*)
Value NUM_OF_PORTS (.*)
Value POWER_STATUS (.*)
Value REDUNDANT_POWER_STATUS (.*)
Value LOADER_VERSION (.*)
Value ROM_VERSION (.*)
Value OPERATION_CODE_VERSION (.*)

Start
  ^\s*Unit\s+\d+$$ -> Continue.Record
  ^\s*Unit\s+${UNIT}$$
  ^\s*Serial\s+Number:\s*${SERIAL}\s*$$
  ^\s*Service\s+Tag:(?:\s*${SERVICE_TAG})?\s*$$
  ^\s*Hardware\s+Version:\s*${HARDWARE_VERSION}\s*$$
  ^\s*Chip\s+Device\s+ID:\s*${CHIP_ID}\s*$$
  ^\s*EPLD\s+Version:\s*${EPLD_VERSION}\s*$$
  ^\s*Number\s+of\s+Ports:\s*${NUM_OF_PORTS}\s*$$
  ^\s*Main\s+Power\s+Status:\s*${POWER_STATUS}\s*$$
  ^\s*Redundant\s+Power\s+Status:\s*${REDUNDANT_POWER_STATUS}\s*$$
  ^\s*Agent.*$$
  ^\s*Unit\s+ID:\s*\d+\s*$$
  ^\s*Loader\s+Version:\s*${LOADER_VERSION}\s*$$
  ^\s*Boot\s+ROM\s+Version:\s*${ROM_VERSION}\s*$$
  ^\s*Operation\s+Code\s+Version:\s*${OPERATION_CODE_VERSION}\s*$$
  ^\s*$$
  ^. -> Error

`