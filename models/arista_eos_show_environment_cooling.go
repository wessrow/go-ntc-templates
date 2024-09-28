package models

type AristaEosShowEnvironmentCooling struct {
	Fan	string	`json:"FAN"`
	Status	string	`json:"STATUS"`
	Config_speed	string	`json:"CONFIG_SPEED"`
	Actual_speed	string	`json:"ACTUAL_SPEED"`
	Sys_cooling_status	string	`json:"SYS_COOLING_STATUS"`
	Ambient_temp	string	`json:"AMBIENT_TEMP"`
	Airflow	string	`json:"AIRFLOW"`
}