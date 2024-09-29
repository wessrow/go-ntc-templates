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

var AristaEosShowEnvironmentCooling_Template = `Value Required FAN (\S+)
Value STATUS (\S+\s?\S+)
Value CONFIG_SPEED (\S+)
Value ACTUAL_SPEED (\S+)
Value Filldown SYS_COOLING_STATUS (\S+)
Value Filldown AMBIENT_TEMP (\S+)
Value Filldown AIRFLOW (\S+)

Start
  ^System\scooling\sstatus\sis:\s${SYS_COOLING_STATUS}
  ^Ambient\stemperature:\s${AMBIENT_TEMP}
  ^Airflow:\s${AIRFLOW}
  ^Fan\s+Status\s+Configured\sSpeed\s+Actual\sSpeed
  ^-+\s+-+\s+-+\s+-+
  ^${FAN}\s+${STATUS}\s+${CONFIG_SPEED}\%?\s+${ACTUAL_SPEED}\%? -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
`