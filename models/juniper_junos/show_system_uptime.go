package juniper_junos

type ShowSystemUptime struct {
	Last_config_year     string `json:"LAST_CONFIG_YEAR"`
	Last_config_user     string `json:"LAST_CONFIG_USER"`
	Load_average_5min    string `json:"LOAD_AVERAGE_5MIN"`
	Timezone             string `json:"TIMEZONE"`
	Boot_time            string `json:"BOOT_TIME"`
	Boot_month           string `json:"BOOT_MONTH"`
	Protocols_month      string `json:"PROTOCOLS_MONTH"`
	Last_config_day      string `json:"LAST_CONFIG_DAY"`
	Last_config_month    string `json:"LAST_CONFIG_MONTH"`
	Time                 string `json:"TIME"`
	Month                string `json:"MONTH"`
	Day                  string `json:"DAY"`
	Protocols_time       string `json:"PROTOCOLS_TIME"`
	Protocols_year       string `json:"PROTOCOLS_YEAR"`
	Year                 string `json:"YEAR"`
	Boot_day             string `json:"BOOT_DAY"`
	Protocols_timezone   string `json:"PROTOCOLS_TIMEZONE"`
	Last_config_timezone string `json:"LAST_CONFIG_TIMEZONE"`
	Load_average_1min    string `json:"LOAD_AVERAGE_1MIN"`
	Users                string `json:"USERS"`
	Load_average_15min   string `json:"LOAD_AVERAGE_15MIN"`
	Time_source          string `json:"TIME_SOURCE"`
	Boot_timezone        string `json:"BOOT_TIMEZONE"`
	Boot_year            string `json:"BOOT_YEAR"`
	Protocols_day        string `json:"PROTOCOLS_DAY"`
	Last_config_time     string `json:"LAST_CONFIG_TIME"`
}

var ShowSystemUptime_Template string = `Value TIME (\d+:\d+:\d+)
Value TIMEZONE (\S+)
Value MONTH (\w+)
Value DAY (\d+)
Value YEAR (\d+)
Value TIME_SOURCE (.*)
Value BOOT_TIME (\d+:\d+:\d+)
Value BOOT_TIMEZONE (\S+)
Value BOOT_MONTH (\w+)
Value BOOT_DAY (\d+)
Value BOOT_YEAR (\d+)
Value PROTOCOLS_TIME (\d+:\d+:\d+)
Value PROTOCOLS_TIMEZONE (\S+)
Value PROTOCOLS_MONTH (\w+)
Value PROTOCOLS_DAY (\d+)
Value PROTOCOLS_YEAR (\d+)
Value LAST_CONFIG_TIME (\d+:\d+:\d+)
Value LAST_CONFIG_TIMEZONE (\S+)
Value LAST_CONFIG_MONTH (\w+)
Value LAST_CONFIG_DAY (\d+)
Value LAST_CONFIG_YEAR (\d+)
Value LAST_CONFIG_USER (\w+)
Value USERS (\d+)
Value LOAD_AVERAGE_1MIN (\d+\.\d{2})
Value LOAD_AVERAGE_5MIN (\d+\.\d{2})
Value LOAD_AVERAGE_15MIN (\d+\.\d{2})

Start
  ^Current\s+time:\s*${YEAR}-${MONTH}-${DAY}\s+${TIME}\s+${TIMEZONE}
  ^Time\s+[Ss]ource:\s*${TIME_SOURCE}
  ^System\s+booted:\s*${BOOT_YEAR}-${BOOT_MONTH}-${BOOT_DAY}\s+${BOOT_TIME}\s+${BOOT_TIMEZONE}\s+\([^)]+\)
  ^Protocols\s+started:\s*${PROTOCOLS_YEAR}-${PROTOCOLS_MONTH}-${PROTOCOLS_DAY}\s+${PROTOCOLS_TIME}\s+${PROTOCOLS_TIMEZONE}\s+\([^)]+\)
  ^Last\s+configured:\s*${LAST_CONFIG_YEAR}-${LAST_CONFIG_MONTH}-${LAST_CONFIG_DAY}\s+${LAST_CONFIG_TIME}\s+${LAST_CONFIG_TIMEZONE}\s+\([^)]+\)\s+by\s+${LAST_CONFIG_USER}
  ^.*,\s*${USERS}\s+users?,\s*load\s+averages:\s*${LOAD_AVERAGE_1MIN},\s*${LOAD_AVERAGE_5MIN},\s*${LOAD_AVERAGE_15MIN}
  ^\s*$$
  ^. -> Error`
