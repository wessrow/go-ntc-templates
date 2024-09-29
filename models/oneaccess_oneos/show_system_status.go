package oneaccess_oneos 

type ShowSystemStatus struct {
	Serial	[]string	`json:"SERIAL"`
	Device_id	string	`json:"DEVICE_ID"`
	Software	string	`json:"SOFTWARE"`
	Boot_version	string	`json:"BOOT_VERSION"`
	Recovery_version	string	`json:"RECOVERY_VERSION"`
	Boot_flags	string	`json:"BOOT_FLAGS"`
	System_time	string	`json:"SYSTEM_TIME"`
	Restarted	string	`json:"RESTARTED"`
	Reload_reason	string	`json:"RELOAD_REASON"`
	Uptime	string	`json:"UPTIME"`
	Uptime_seconds	string	`json:"UPTIME_SECONDS"`
	Uptime_minutes	string	`json:"UPTIME_MINUTES"`
	Uptime_hours	string	`json:"UPTIME_HOURS"`
	Uptime_days	string	`json:"UPTIME_DAYS"`
	Version	string	`json:"VERSION"`
}

var ShowSystemStatus_Template = `Value List SERIAL (\w+)
Value DEVICE_ID ([^ ]+)
Value SOFTWARE (.+)
Value BOOT_VERSION (.+)
Value RECOVERY_VERSION (.+)
Value BOOT_FLAGS (\dx\d+)
Value SYSTEM_TIME (.*)
Value RESTARTED (.*)
Value RELOAD_REASON (.*)
Value UPTIME (.*)
Value UPTIME_SECONDS (\d+)
Value UPTIME_MINUTES (\d+)
Value UPTIME_HOURS (\d+)
Value UPTIME_DAYS (\d+)
Value VERSION (\d)


Start
  ^System\WInformation\Wfor\Wdevice\W${DEVICE_ID}\W+S\/N\W${SERIAL}
  ^Software\Wversion\W+:\W+${SOFTWARE} -> Continue
  ^Software\Wversion\W+:\W+.*\-V?${VERSION}\..*
  ^Boot\Wversion\W+:\W${BOOT_VERSION}
  ^Boot\WFlags\W+:\W+${BOOT_FLAGS}
  ^Recovery\Wversion\W+:\W+${RECOVERY_VERSION}
  ^Current\Wsystem\Wtime\W+:\W+${SYSTEM_TIME}
  ^System\Wstarted\W+:\W+${RESTARTED}
  ^Start\Wcaused\Wby\W+:\W+${RELOAD_REASON}
  ^Sys\WUp\Wtime\W+:\W+(?:${UPTIME_DAYS}d)? ?(?:${UPTIME_HOURS}h)? ?(?:${UPTIME_MINUTES}m)? ?(?:${UPTIME_SECONDS}s)? -> Continue
  ^Sys\WUp\Wtime\W+:\W+${UPTIME} -> Record
  ^\s*(Software|Boot|Recovery)\s+created\s+on\s+:
  ^License\s+token\s+:
  ^System\s+clock\s+ticks\s+:
  ^(Current|Average)\s+.*CPU\s+load
  ^Free\s+/\s+Max\s+RAM\s+:
  ^PSU\s+:
  # Do not care about data after first section
  ^(Core\s+Type\s+last\s+sec|Temperatures\s*:) -> EOF
  ^\s*$$
  ^. -> Error
`