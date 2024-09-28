package models

type CiscoIosShowRedundancy struct {
	System_uptime_years	string	`json:"SYSTEM_UPTIME_YEARS"`
	System_uptime_weeks	string	`json:"SYSTEM_UPTIME_WEEKS"`
	System_uptime_days	string	`json:"SYSTEM_UPTIME_DAYS"`
	System_uptime_hours	string	`json:"SYSTEM_UPTIME_HOURS"`
	System_uptime_minutes	string	`json:"SYSTEM_UPTIME_MINUTES"`
	Switchovers	string	`json:"SWITCHOVERS"`
	Standby_failures	string	`json:"STANDBY_FAILURES"`
	Last_switchover_reason	string	`json:"LAST_SWITCHOVER_REASON"`
	Hardware_mode	string	`json:"HARDWARE_MODE"`
	Configured_redundancy_mode	string	`json:"CONFIGURED_REDUNDANCY_MODE"`
	Operating_redundancy_mode	string	`json:"OPERATING_REDUNDANCY_MODE"`
	Maintenance_mode	string	`json:"MAINTENANCE_MODE"`
	Communication_status	string	`json:"COMMUNICATION_STATUS"`
	Active_slot	string	`json:"ACTIVE_SLOT"`
	Active_software_state	string	`json:"ACTIVE_SOFTWARE_STATE"`
	Active_uptime_years	string	`json:"ACTIVE_UPTIME_YEARS"`
	Active_uptime_weeks	string	`json:"ACTIVE_UPTIME_WEEKS"`
	Active_uptime_days	string	`json:"ACTIVE_UPTIME_DAYS"`
	Active_uptime_hours	string	`json:"ACTIVE_UPTIME_HOURS"`
	Active_uptime_minutes	string	`json:"ACTIVE_UPTIME_MINUTES"`
	Active_major_release	string	`json:"ACTIVE_MAJOR_RELEASE"`
	Active_minor_release	string	`json:"ACTIVE_MINOR_RELEASE"`
	Active_maintenance_release	string	`json:"ACTIVE_MAINTENANCE_RELEASE"`
	Active_boot_directory	string	`json:"ACTIVE_BOOT_DIRECTORY"`
	Active_boot_file	string	`json:"ACTIVE_BOOT_FILE"`
	Active_config_register	string	`json:"ACTIVE_CONFIG_REGISTER"`
	Standby_slot	string	`json:"STANDBY_SLOT"`
	Standby_software_state	string	`json:"STANDBY_SOFTWARE_STATE"`
	Standby_uptime_years	string	`json:"STANDBY_UPTIME_YEARS"`
	Standby_uptime_weeks	string	`json:"STANDBY_UPTIME_WEEKS"`
	Standby_uptime_days	string	`json:"STANDBY_UPTIME_DAYS"`
	Standby_uptime_hours	string	`json:"STANDBY_UPTIME_HOURS"`
	Standby_uptime_minutes	string	`json:"STANDBY_UPTIME_MINUTES"`
	Standby_major_release	string	`json:"STANDBY_MAJOR_RELEASE"`
	Standby_minor_release	string	`json:"STANDBY_MINOR_RELEASE"`
	Standby_maintenance_release	string	`json:"STANDBY_MAINTENANCE_RELEASE"`
	Standby_boot_directory	string	`json:"STANDBY_BOOT_DIRECTORY"`
	Standby_boot_file	string	`json:"STANDBY_BOOT_FILE"`
	Standby_config_register	string	`json:"STANDBY_CONFIG_REGISTER"`
	Standby_status	string	`json:"STANDBY_STATUS"`
}