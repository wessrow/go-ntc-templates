package cisco_ios

type ShowVersion struct {
	Uptime_years    string   `json:"UPTIME_YEARS"`
	Serial          []string `json:"SERIAL"`
	Mac_address     []string `json:"MAC_ADDRESS"`
	Restarted       string   `json:"RESTARTED"`
	Rommon          string   `json:"ROMMON"`
	Reload_reason   string   `json:"RELOAD_REASON"`
	Release         string   `json:"RELEASE"`
	Uptime_hours    string   `json:"UPTIME_HOURS"`
	Running_image   string   `json:"RUNNING_IMAGE"`
	Config_register string   `json:"CONFIG_REGISTER"`
	Software_image  string   `json:"SOFTWARE_IMAGE"`
	Version         string   `json:"VERSION"`
	Uptime_weeks    string   `json:"UPTIME_WEEKS"`
	Uptime_days     string   `json:"UPTIME_DAYS"`
	Uptime_minutes  string   `json:"UPTIME_MINUTES"`
	Hardware        []string `json:"HARDWARE"`
	Hostname        string   `json:"HOSTNAME"`
	Uptime          string   `json:"UPTIME"`
}

var ShowVersion_Template string = `Value SOFTWARE_IMAGE (\S+)
Value VERSION (.+?)
Value RELEASE (\S+)
Value ROMMON (\S+)
Value HOSTNAME (\S+)
Value UPTIME (.+)
Value UPTIME_YEARS (\d+)
Value UPTIME_WEEKS (\d+)
Value UPTIME_DAYS (\d+)
Value UPTIME_HOURS (\d+)
Value UPTIME_MINUTES (\d+)
Value RELOAD_REASON (.+?)
Value RUNNING_IMAGE (\S+)
Value List HARDWARE (\S+|\S+\d\S+)
Value List SERIAL (\w+)
Value CONFIG_REGISTER (\S+)
Value List MAC_ADDRESS ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5})
Value RESTARTED (.+)

Start
  ^.*Software,*\s+\(${SOFTWARE_IMAGE}\),\sVersion\s${VERSION},*\s+RELEASE.*\(${RELEASE}\)
  # CUST-SPECIAL:
  ^.*Software,*\s+\(${SOFTWARE_IMAGE}\),\sVersion\s${VERSION},*\s+\S+.*:${RELEASE}
  ^.*Software,*\s+\(${SOFTWARE_IMAGE}\),\sVersion\s${VERSION},
  ^ROM:\s+${ROMMON}
  ^\s*${HOSTNAME}\s+uptime\s+is\s+${UPTIME} -> Continue
  ^.*\s+uptime\s+is.*\s+${UPTIME_YEARS}\syear -> Continue
  ^.*\s+uptime\s+is.*\s+${UPTIME_WEEKS}\sweek -> Continue
  ^.*\s+uptime\s+is.*\s+${UPTIME_DAYS}\sday -> Continue
  ^.*\s+uptime\s+is.*\s+${UPTIME_HOURS}\shour -> Continue
  ^.*\s+uptime\s+is.*\s+${UPTIME_MINUTES}\sminute
  ^[sS]ystem\s+image\s+file\s+is\s+"(.*?):${RUNNING_IMAGE}"
  ^(?:[lL]ast\s+reload\s+reason:|System\s+returned\s+to\s+ROM\s+by)\s+${RELOAD_REASON}\s*$$
  ^[Pp]rocessor\s+board\s+ID\s+${SERIAL}
  ^[Cc]isco\s+${HARDWARE}\s+\(.+\).+
  ^[Cc]onfiguration\s+register\s+is\s+${CONFIG_REGISTER}
  ^Base\s+[Ee]thernet\s+MAC\s+[Aa]ddress\s+:\s+${MAC_ADDRESS}
  ^System\s+restarted\s+at\s+${RESTARTED}$$
  ^Switch\s+Port -> Stack
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Switch\s\d+ -> Stack
  ^Load\s+for\s+
  ^Time\s+source\s+is


Stack
  ^[Ss]ystem\s+[Ss]erial\s+[Nn]umber\s+:\s+${SERIAL}
  ^[Mm]odel\s+[Nn]umber\s+:\s+${HARDWARE}\s*
  ^[Cc]onfiguration\s+register\s+is\s+${CONFIG_REGISTER}
  ^Base [Ee]thernet MAC [Aa]ddress\s+:\s+${MAC_ADDRESS}
`
