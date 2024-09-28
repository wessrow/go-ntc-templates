package models

type CiscoIosShowEnvironmentPowerAll struct {
	Ps_module	string	`json:"PS_MODULE"`
	Pid	string	`json:"PID"`
	Serial	string	`json:"SERIAL"`
	Status	string	`json:"STATUS"`
	Sys_pwr	string	`json:"SYS_PWR"`
	Poe_pwr	string	`json:"POE_PWR"`
	Watts	string	`json:"WATTS"`
}