package oneaccess_oneos

type ShowSystemSecureCrashlog struct {
	Reload_reason     string   `json:"RELOAD_REASON"`
	Uptime            string   `json:"UPTIME"`
	Version           string   `json:"VERSION"`
	Coredump_file     string   `json:"COREDUMP_FILE"`
	Crash_time        string   `json:"CRASH_TIME"`
	Software          string   `json:"SOFTWARE"`
	Boot_version      string   `json:"BOOT_VERSION"`
	Core_generated_by string   `json:"CORE_GENERATED_BY"`
	Crash_caused_by   string   `json:"CRASH_CAUSED_BY"`
	Recovery_version  string   `json:"RECOVERY_VERSION"`
	Uptime_minutes    string   `json:"UPTIME_MINUTES"`
	Uptime_hours      string   `json:"UPTIME_HOURS"`
	Uptime_days       string   `json:"UPTIME_DAYS"`
	Crash_filename    string   `json:"CRASH_FILENAME"`
	Serial            []string `json:"SERIAL"`
	Deviceid          string   `json:"DEVICEID"`
	Restarted         string   `json:"RESTARTED"`
	Uptime_seconds    string   `json:"UPTIME_SECONDS"`
}

var ShowSystemSecureCrashlog_Template string = `Value COREDUMP_FILE (coredump.*txt)
Value CRASH_CAUSED_BY (.*\S)
Value CRASH_TIME (.*\S)
Value CRASH_FILENAME (.*\S)
Value List SERIAL (\w+)
Value DEVICEID (\S+)
Value SOFTWARE (.+)
Value BOOT_VERSION (.+)
Value RECOVERY_VERSION (.+)
Value RESTARTED (.*)
Value RELOAD_REASON (.*)
Value UPTIME (.*)
Value UPTIME_SECONDS (\d+)
Value UPTIME_MINUTES (\d+)
Value UPTIME_HOURS (\d+)
Value UPTIME_DAYS (\d+)
Value VERSION (\d)
Value CORE_GENERATED_BY (\S+)

Start
  ^${COREDUMP_FILE}$$
  ^\-+
  ^\s*=+\s*$$
  ^Last\sRaw\sLog\ssaved\sin\sFlash
  ^Crash\scaused\sby\s+:\s+${CRASH_CAUSED_BY}
  ^Crash\stime\s+:\s+${CRASH_TIME}
  ^Crash\s(filename|file\sname)\s+:\s+${CRASH_FILENAME}
  ^Device\sidentifier\s+:\s+${DEVICEID}\sS\/N\s${SERIAL}
  ^Software\Wversion\W+:\W+${SOFTWARE} -> Continue
  ^Software\Wversion\W+:\W+.*\-V?${VERSION}\..*
  ^Boot\Wversion\W+:\W${BOOT_VERSION}
  ^Recovery\Wversion\W+:\W+${RECOVERY_VERSION}
  ^System\Wstarted\W+:\W+${RESTARTED}
  ^Start\Wcaused\Wby\W+:\W+${RELOAD_REASON}
  ^Sys\WUp\Wtime\W+:\W+(?:${UPTIME_DAYS}d)? ?(?:${UPTIME_HOURS}h)? ?(?:${UPTIME_MINUTES}m)? ?(?:${UPTIME_SECONDS}s)? -> Continue
  ^Sys\sUp\s[tT]ime\s+:\s+${UPTIME}
  ^Core\swas\sgenerated\sby\s'${CORE_GENERATED_BY}'
  ^(Software|Boot|Recovery)\s+created\s+on\s+:
  ^License\s+token\s+:
  ^[A-Z\s]+\s*$$
  ^(code|message)\s+:
  ^file
  # Not capturing anything outside of the first section
  ^(\*|=+\s+\S+) -> EOF
  # not capturing anything indented
  ^\s
  ^. -> Error
`
