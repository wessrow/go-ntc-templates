package models

type AristaEosShowEnvironmentTemperature struct {
	Sensor	string	`json:"SENSOR"`
	Description	string	`json:"DESCRIPTION"`
	Current_value	string	`json:"CURRENT_VALUE"`
	Setpoint	string	`json:"SETPOINT"`
	Alert_limit	string	`json:"ALERT_LIMIT"`
	Critical_limit	string	`json:"CRITICAL_LIMIT"`
	Module	string	`json:"MODULE"`
}