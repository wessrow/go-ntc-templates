package brocade_fastiron

type ShowVersion struct {
	Sw_version           []string `json:"SW_VERSION"`
	Poe                  []string `json:"POE"`
	Flash_measurement    string   `json:"FLASH_MEASUREMENT"`
	Start_time           string   `json:"START_TIME"`
	Serial               []string `json:"SERIAL"`
	Ports                []string `json:"PORTS"`
	License              []string `json:"LICENSE"`
	Ram_measurement      string   `json:"RAM_MEASUREMENT"`
	Uptime_days          []string `json:"UPTIME_DAYS"`
	Ram                  string   `json:"RAM"`
	Uptime_hours         []string `json:"UPTIME_HOURS"`
	Uptime_minutes       []string `json:"UPTIME_MINUTES"`
	Sw_bin               []string `json:"SW_BIN"`
	Module_type          []string `json:"MODULE_TYPE"`
	Start_month          string   `json:"START_MONTH"`
	Start_timezone       string   `json:"START_TIMEZONE"`
	Start_year           string   `json:"START_YEAR"`
	Switch_id            []string `json:"SWITCH_ID"`
	Boot_monitor_version string   `json:"BOOT_MONITOR_VERSION"`
	Hw                   string   `json:"HW"`
	Flash                string   `json:"FLASH"`
	Model                []string `json:"MODEL"`
	Reload_reason        string   `json:"RELOAD_REASON"`
	Start_day            string   `json:"START_DAY"`
	Start_date           string   `json:"START_DATE"`
	Slot                 []string `json:"SLOT"`
	Uptime_seconds       []string `json:"UPTIME_SECONDS"`
}

var ShowVersion_Template string = `Value List SWITCH_ID (\d+)
Value List SW_BIN (\S+)
Value List SW_VERSION (\S+)
Value BOOT_MONITOR_VERSION (.*)
Value HW (.*)
Value List SLOT (\d+)
Value List MODEL (\S+)
Value List POE (POE)
Value List PORTS (\d+)
Value List MODULE_TYPE (\S+)
Value List SERIAL (\S+$)
Value List LICENSE (\S+)
Value FLASH (\d+)
Value FLASH_MEASUREMENT (\S+)
Value RAM (\d+)
Value RAM_MEASUREMENT (\S+)
Value List UPTIME_DAYS (\d+)
Value List UPTIME_HOURS (\d+)
Value List UPTIME_MINUTES (\d+)
Value List UPTIME_SECONDS (\d+)
Value START_TIME (\d+:\d+:\d+)
Value START_TIMEZONE (\S+)
Value START_DAY (\S+)
Value START_MONTH (\S+)
Value START_DATE (\d+)
Value START_YEAR (\d+)
Value RELOAD_REASON (.+?)

Start
  ^show\s+version
  ^\s+UNIT\s+${SWITCH_ID}
  ^\s+\(.+?\)\s+from\s+\S+\s+${SW_BIN}\s*$$
  ^\s*SW:\s+Version\s+${SW_VERSION}
  ^\s*Boot-Monitor.*,\s+Version:${BOOT_MONITOR_VERSION}\s*$$
  ^\s*HW:\s+${HW}\s*$$
  ^\s+Copyright
  ^\s*$$
  ^=+\s*$$ -> Hardware
  ^. -> Error

Hardware
  ^UNIT\s+\d+:\s+SL\s+${SLOT}:\s+${MODEL}\s+(?:${POE}\s+|)${PORTS}(?:-|)port\s+${MODULE_TYPE}\s+Module\s*$$
  ^\s+Serial\s+#:\s+${SERIAL}\s*$$
  ^\s+License:\s+${LICENSE}
  ^.*ENGINE
  ^=
  ^.+uptime\s+is\s+(?:${UPTIME_DAYS}\s+day\(s\)\s+|)(?:${UPTIME_HOURS}\s+hour\(s\)\s+|)(?:${UPTIME_MINUTES}\s+minute\(s\)\s+|)(?:${UPTIME_SECONDS}\s+second\(s\)|)
  ^${FLASH}\s+${FLASH_MEASUREMENT}\s+flash
  ^\s+${RAM}\s+${RAM_MEASUREMENT}.+RAM
  ^.+started\s+at\s+${START_TIME}\s+${START_TIMEZONE}\s+${START_DAY}\s+${START_MONTH}\s+${START_DATE}\s+${START_YEAR}
  ^\s+The\s+system\s+:\s+started=${RELOAD_REASON}\s*$$
  ^\s*$$
  ^.+processor
  ^My\s+stack
  ^. -> Error
`
