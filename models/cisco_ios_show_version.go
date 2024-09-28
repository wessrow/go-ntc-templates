package models

type CiscoIosShowVersion struct {
	Software_image	string	`json:"SOFTWARE_IMAGE"`
	Version	string	`json:"VERSION"`
	Release	string	`json:"RELEASE"`
	Rommon	string	`json:"ROMMON"`
	Hostname	string	`json:"HOSTNAME"`
	Uptime	string	`json:"UPTIME"`
	Uptime_years	string	`json:"UPTIME_YEARS"`
	Uptime_weeks	string	`json:"UPTIME_WEEKS"`
	Uptime_days	string	`json:"UPTIME_DAYS"`
	Uptime_hours	string	`json:"UPTIME_HOURS"`
	Uptime_minutes	string	`json:"UPTIME_MINUTES"`
	Reload_reason	string	`json:"RELOAD_REASON"`
	Running_image	string	`json:"RUNNING_IMAGE"`
	Hardware	[]string	`json:"HARDWARE"`
	Serial	[]string	`json:"SERIAL"`
	Config_register	string	`json:"CONFIG_REGISTER"`
	Mac_address	[]string	`json:"MAC_ADDRESS"`
	Restarted	string	`json:"RESTARTED"`
}