package models

type CiscoIosShowEnvironmentTemperature struct {
	Switch	string	`json:"SWITCH"`
	Switch_temperature	string	`json:"SWITCH_TEMPERATURE"`
	Inlet_temperature_value	string	`json:"INLET_TEMPERATURE_VALUE"`
	Inlet_temperature_state	string	`json:"INLET_TEMPERATURE_STATE"`
	Inlet_yellow_threshold	string	`json:"INLET_YELLOW_THRESHOLD"`
	Inlet_red_threshold	string	`json:"INLET_RED_THRESHOLD"`
	Hotspot_temperature_value	string	`json:"HOTSPOT_TEMPERATURE_VALUE"`
	Hotspot_temperature_state	string	`json:"HOTSPOT_TEMPERATURE_STATE"`
	Hotspot_yellow_threshold	string	`json:"HOTSPOT_YELLOW_THRESHOLD"`
	Hotspot_red_threshold	string	`json:"HOTSPOT_RED_THRESHOLD"`
}