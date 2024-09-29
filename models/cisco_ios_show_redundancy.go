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

var CiscoIosShowRedundancy_Template = `Value SYSTEM_UPTIME_YEARS (\d+)
Value SYSTEM_UPTIME_WEEKS (\d+)
Value SYSTEM_UPTIME_DAYS (\d+)
Value SYSTEM_UPTIME_HOURS (\d+)
Value SYSTEM_UPTIME_MINUTES (\d+)
Value SWITCHOVERS (\d+)
Value STANDBY_FAILURES (\d+)
Value LAST_SWITCHOVER_REASON (.+?)
Value HARDWARE_MODE (.+?)
Value CONFIGURED_REDUNDANCY_MODE (.+?)
Value OPERATING_REDUNDANCY_MODE (.+?)
Value MAINTENANCE_MODE (.+?)
Value COMMUNICATION_STATUS (.+?)
Value ACTIVE_SLOT (\S+)
Value ACTIVE_SOFTWARE_STATE (.+?)
Value ACTIVE_UPTIME_YEARS (\d+)
Value ACTIVE_UPTIME_WEEKS (\d+)
Value ACTIVE_UPTIME_DAYS (\d+)
Value ACTIVE_UPTIME_HOURS (\d+)
Value ACTIVE_UPTIME_MINUTES (\d+)
Value ACTIVE_MAJOR_RELEASE (\d+)
Value ACTIVE_MINOR_RELEASE (\d+)
Value ACTIVE_MAINTENANCE_RELEASE (\S+)
Value ACTIVE_BOOT_DIRECTORY (\S+)
Value ACTIVE_BOOT_FILE (\S+)
Value ACTIVE_CONFIG_REGISTER (\S+)
Value STANDBY_SLOT (\S+)
Value STANDBY_SOFTWARE_STATE (.+?)
Value STANDBY_UPTIME_YEARS (\d+)
Value STANDBY_UPTIME_WEEKS (\d+)
Value STANDBY_UPTIME_DAYS (\d+)
Value STANDBY_UPTIME_HOURS (\d+)
Value STANDBY_UPTIME_MINUTES (\d+)
Value STANDBY_MAJOR_RELEASE (\d+)
Value STANDBY_MINOR_RELEASE (\d+)
Value STANDBY_MAINTENANCE_RELEASE (\S+)
Value STANDBY_BOOT_DIRECTORY (\S+)
Value STANDBY_BOOT_FILE (\S+)
Value STANDBY_CONFIG_REGISTER (\S+)
Value STANDBY_STATUS (.+?)


Start
  ^Redundant\s+System\s+Information
  ^\s*-+\s*$$
  ^\s*Available\s+system\s+uptime\s*=\s*(${SYSTEM_UPTIME_YEARS}\s*year(s)?,)?\s*(${SYSTEM_UPTIME_WEEKS}\s*week(s)?,)?\s*(${SYSTEM_UPTIME_DAYS}\s*day(s)?,)?\s*(${SYSTEM_UPTIME_HOURS}\s*hour(s)?,)?\s*${SYSTEM_UPTIME_MINUTES}\s*minute(s)?\s*$$
  ^\s*Switchovers\s+system\s+experienced\s*=\s*${SWITCHOVERS}\s*$$
  ^\s*Standby\s+failures\s*=\s*${STANDBY_FAILURES}\s*$$
  ^\s*Last\s+switchover\s+reason\s*=\s*${LAST_SWITCHOVER_REASON}\s*$$
  ^\s*Hardware\s+Mode\s*=\s*${HARDWARE_MODE}\s*$$
  ^\s*Configured\s+Redundancy\s+Mode\s*=\s*${CONFIGURED_REDUNDANCY_MODE}\s*$$
  ^\s*Operating\s+Redundancy\s+Mode\s*=\s*${OPERATING_REDUNDANCY_MODE}\s*$$
  ^\s*Maintenance\s+Mode\s*=\s*${MAINTENANCE_MODE}\s*$$
  ^\s*Communications\s*=\s*${COMMUNICATION_STATUS}\s*$$
  ^\s*Current\s+Processor -> Current
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^\s*$$
  ^. -> Error

Current
  ^\s*-+\s*$$
  ^\s*Active\s+Location\s*=\s*slot\s+${ACTIVE_SLOT}\s*$$
  ^\s*Current\s+Software\s+state\s*=\s*${ACTIVE_SOFTWARE_STATE}\s*$$
  ^\s*Uptime\s+in\s+current\s+state\s*=\s*(${ACTIVE_UPTIME_YEARS}\s*year(s)?,)?\s*(${ACTIVE_UPTIME_WEEKS}\s*week(s)?,)?\s*(${ACTIVE_UPTIME_DAYS}\s*day(s)?,)?\s*(${ACTIVE_UPTIME_HOURS}\s*hour(s)?,)?\s*${ACTIVE_UPTIME_MINUTES}\s*minute(s)?\s*$$
  ^\s*Image\s+Version
  ^.+Version\s+${ACTIVE_MAJOR_RELEASE}\.${ACTIVE_MINOR_RELEASE}\(${ACTIVE_MAINTENANCE_RELEASE}\)
  ^.+[Cc][Ii][Ss][Cc][Oo]
  ^.*(?:Copyright|Compiled|CONFIG_FILE|BOOTLDR)
  ^\s*BOOT\s*=\s*${ACTIVE_BOOT_DIRECTORY}:${ACTIVE_BOOT_FILE}\s*$$
  ^\s*Configuration\s+register\s*=\s*${ACTIVE_CONFIG_REGISTER}\s*$$
  ^Peer\s*\(slot:\s*unavailable\)\s*\w*\s*\w*\s*\w*\s*\w*\s*\w\s*\w*\s*\w*\s*\w*\s*\w*\s*\'${STANDBY_STATUS}\'\s*state$$ -> EOF
  ^\s*Peer\s+Processor -> Peer
  ^\s*$$
  ^. -> Error

Peer
  ^\s*-+\s*$$
  ^\s*Standby\s+Location\s*=\s*slot\s+${STANDBY_SLOT}\s*$$
  ^\s*Current\s+Software\s+state\s*=\s*${STANDBY_SOFTWARE_STATE}\s*$$
  ^\s*Uptime\s+in\s+current\s+state\s*=\s*(${STANDBY_UPTIME_YEARS}\s*year(s)?,)?\s*(${STANDBY_UPTIME_WEEKS}\s*week(s)?,)?\s*(${STANDBY_UPTIME_DAYS}\s*day(s)?,)?\s*(${STANDBY_UPTIME_HOURS}\s*hour(s)?,)?\s*${STANDBY_UPTIME_MINUTES}\s*minute(s)?\s*$$
  ^\s*Image\s+Version
  ^.+Version\s+${STANDBY_MAJOR_RELEASE}\.${STANDBY_MINOR_RELEASE}\(${STANDBY_MAINTENANCE_RELEASE}\)
  ^.+[Cc][Ii][Ss][Cc][Oo]
  ^.*(?:Copyright|Compiled|CONFIG_FILE|BOOTLDR)
  ^\s*BOOT\s*=\s*${STANDBY_BOOT_DIRECTORY}:${STANDBY_BOOT_FILE}\s*$$
  ^\s*Configuration\s+register\s*=\s*${STANDBY_CONFIG_REGISTER}\s*$$
  ^\s*$$
  ^\s*Compiled\s*
  ^. -> Error

`