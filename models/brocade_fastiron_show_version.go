package models

type BrocadeFastironShowVersion struct {
	Switch_id	[]string	`json:"SWITCH_ID"`
	Sw_bin	[]string	`json:"SW_BIN"`
	Sw_version	[]string	`json:"SW_VERSION"`
	Boot_monitor_version	string	`json:"BOOT_MONITOR_VERSION"`
	Hw	string	`json:"HW"`
	Slot	[]string	`json:"SLOT"`
	Model	[]string	`json:"MODEL"`
	Poe	[]string	`json:"POE"`
	Ports	[]string	`json:"PORTS"`
	Module_type	[]string	`json:"MODULE_TYPE"`
	Serial	[]string	`json:"SERIAL"`
	License	[]string	`json:"LICENSE"`
	Flash	string	`json:"FLASH"`
	Flash_measurement	string	`json:"FLASH_MEASUREMENT"`
	Ram	string	`json:"RAM"`
	Ram_measurement	string	`json:"RAM_MEASUREMENT"`
	Uptime_days	[]string	`json:"UPTIME_DAYS"`
	Uptime_hours	[]string	`json:"UPTIME_HOURS"`
	Uptime_minutes	[]string	`json:"UPTIME_MINUTES"`
	Uptime_seconds	[]string	`json:"UPTIME_SECONDS"`
	Start_time	string	`json:"START_TIME"`
	Start_timezone	string	`json:"START_TIMEZONE"`
	Start_day	string	`json:"START_DAY"`
	Start_month	string	`json:"START_MONTH"`
	Start_date	string	`json:"START_DATE"`
	Start_year	string	`json:"START_YEAR"`
	Reload_reason	string	`json:"RELOAD_REASON"`
}