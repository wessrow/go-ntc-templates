package models

type AristaEosShowEnvironmentPower struct {
	Power_supply_location	string	`json:"POWER_SUPPLY_LOCATION"`
	Power_supply_pid	string	`json:"POWER_SUPPLY_PID"`
	Power_supply_watts	string	`json:"POWER_SUPPLY_WATTS"`
	Power_supply_input	string	`json:"POWER_SUPPLY_INPUT"`
	Power_supply_output	string	`json:"POWER_SUPPLY_OUTPUT"`
	Power_supply_watts_using	string	`json:"POWER_SUPPLY_WATTS_USING"`
	Power_supply_status	string	`json:"POWER_SUPPLY_STATUS"`
	Power_supply_uptime	string	`json:"POWER_SUPPLY_UPTIME"`
	System_total_power_watts	string	`json:"SYSTEM_TOTAL_POWER_WATTS"`
	System_total_power_used_watts	string	`json:"SYSTEM_TOTAL_POWER_USED_WATTS"`
}