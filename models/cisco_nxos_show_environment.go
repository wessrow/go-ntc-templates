package models

type CiscoNxosShowEnvironment struct {
	Power_supply	string	`json:"POWER_SUPPLY"`
	Power_supply_model	string	`json:"POWER_SUPPLY_MODEL"`
	Power_supply_output	string	`json:"POWER_SUPPLY_OUTPUT"`
	Power_supply_input	string	`json:"POWER_SUPPLY_INPUT"`
	Power_supply_capacity	string	`json:"POWER_SUPPLY_CAPACITY"`
	Power_supply_status	string	`json:"POWER_SUPPLY_STATUS"`
	Fan	string	`json:"FAN"`
	Fan_status	string	`json:"FAN_STATUS"`
	Temperature_module	string	`json:"TEMPERATURE_MODULE"`
	Temperature_sensor	string	`json:"TEMPERATURE_SENSOR"`
	Temperature_major_thresh	string	`json:"TEMPERATURE_MAJOR_THRESH"`
	Temperature_minor_thresh	string	`json:"TEMPERATURE_MINOR_THRESH"`
	Temperature_current	string	`json:"TEMPERATURE_CURRENT"`
	Temperature_status	string	`json:"TEMPERATURE_STATUS"`
}