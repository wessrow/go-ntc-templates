package yamaha

type ShowEnvironment struct {
	Cpu_5_min           string `json:"CPU_5_MIN"`
	Default_firmware    string `json:"DEFAULT_FIRMWARE"`
	Security_class      string `json:"SECURITY_CLASS"`
	Inside_temperature  string `json:"INSIDE_TEMPERATURE"`
	Serial              string `json:"SERIAL"`
	Default_config_file string `json:"DEFAULT_CONFIG_FILE"`
	Cpu_1_min           string `json:"CPU_1_MIN"`
	Config_file         string `json:"CONFIG_FILE"`
	Boot_time           string `json:"BOOT_TIME"`
	Current_time        string `json:"CURRENT_TIME"`
	Elasped_time        string `json:"ELASPED_TIME"`
	Telnet              string `json:"TELNET"`
	Bootrom             string `json:"BOOTROM"`
	Version             string `json:"VERSION"`
	Firmware            string `json:"FIRMWARE"`
	Forget              string `json:"FORGET"`
	Cpu_5_sec           string `json:"CPU_5_SEC"`
	Memory_used         string `json:"MEMORY_USED"`
}

var ShowEnvironment_Template string = `# You need to set "console character ascci".
Value BOOTROM (\S+)
Value VERSION (\S+)
Value SERIAL (\w+)
Value CPU_5_SEC (\d+)
Value CPU_1_MIN (\d+)
Value CPU_5_MIN (\d+)
Value MEMORY_USED (\d+)
Value FIRMWARE (\S+)
Value CONFIG_FILE (\S+)
Value DEFAULT_FIRMWARE (\S+)
Value DEFAULT_CONFIG_FILE (\S+)
Value BOOT_TIME (.+)
Value CURRENT_TIME (.+)
Value ELASPED_TIME (.+)
Value SECURITY_CLASS ([1-3])
Value FORGET (ON|OFF)
Value TELNET (ON|OFF)
Value INSIDE_TEMPERATURE (\d+)

Start
  ^.+\sBootROM\sVer\.${BOOTROM}
  ^.+\s+Rev\.${VERSION}\s+\(.+\)
  ^\s+main:.+serial=${SERIAL}\s+MAC-Address.+
  ^.+=([a-f0-9:]{17})
  ^CPU:\s+${CPU_5_SEC}%\(5sec\)\s+${CPU_1_MIN}%\(1min\)\s+${CPU_5_MIN}%\(5min\)\s+Memory:\s+${MEMORY_USED}%\s+used
  ^Packet-buffer:.+
  ^Firmware:\s${FIRMWARE}\s+Config\.\sfile:\s+${CONFIG_FILE}
  ^Default\sfirmware:\s${DEFAULT_FIRMWARE}\s+Default\sconfig\.\sfile:\s${DEFAULT_CONFIG_FILE}
  ^Boot\stime:\s+${BOOT_TIME}
  ^Current\stime:\s+${CURRENT_TIME}
  ^Elapsed\stime\sfrom\sboot:\s+${ELASPED_TIME}
  ^Security\sClass:\s+${SECURITY_CLASS},\s+FORGET:\s+${FORGET},\s+TELNET:\s+${TELNET}
  ^Inside\sTemperature\(C\.\):\s+${INSIDE_TEMPERATURE}
  ^\s*$$
  ^. -> Error
`
