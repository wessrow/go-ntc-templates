package models

type CiscoNxosShowEnvironmentTemperature struct {
	Current_value	string	`json:"CURRENT_VALUE"`
	Major_threshold	string	`json:"MAJOR_THRESHOLD"`
	Minor_threshold	string	`json:"MINOR_THRESHOLD"`
	Status	string	`json:"STATUS"`
	Module	string	`json:"MODULE"`
	Sensor	string	`json:"SENSOR"`
}