package models

type CiscoIosShowBoot struct {
	Switch_number	string	`json:"SWITCH_NUMBER"`
	Boot_path	string	`json:"BOOT_PATH"`
	Config_file	string	`json:"CONFIG_FILE"`
	Priv_config_file	string	`json:"PRIV_CONFIG_FILE"`
	Enable_break	string	`json:"ENABLE_BREAK"`
	Manual_boot	string	`json:"MANUAL_BOOT"`
	Allow_dev_key	string	`json:"ALLOW_DEV_KEY"`
	Helper_path_list	string	`json:"HELPER_PATH_LIST"`
	Auto_upgrade	string	`json:"AUTO_UPGRADE"`
	Auto_upgrade_path	string	`json:"AUTO_UPGRADE_PATH"`
	Buffer_size	string	`json:"BUFFER_SIZE"`
	Timeout_config_download	string	`json:"TIMEOUT_CONFIG_DOWNLOAD"`
	Config_download_dhcp	string	`json:"CONFIG_DOWNLOAD_DHCP"`
	Config_download_dhcp_next_boot	string	`json:"CONFIG_DOWNLOAD_DHCP_NEXT_BOOT"`
	Boot_mode	string	`json:"BOOT_MODE"`
	Boot_optimization	string	`json:"BOOT_OPTIMIZATION"`
	Current_boot_variables	string	`json:"CURRENT_BOOT_VARIABLES"`
}

var CiscoIosShowBoot_Template = `Value SWITCH_NUMBER (\d+)
Value BOOT_PATH (\S+)
Value CONFIG_FILE (\S+)
Value PRIV_CONFIG_FILE (\S+)
Value ENABLE_BREAK (yes|no)
Value MANUAL_BOOT (yes|no)
Value ALLOW_DEV_KEY (yes|no)
Value HELPER_PATH_LIST (\S+)
Value AUTO_UPGRADE (yes|no)
Value AUTO_UPGRADE_PATH (\S+)
Value BUFFER_SIZE (\d+)
Value TIMEOUT_CONFIG_DOWNLOAD (\d+)
Value CONFIG_DOWNLOAD_DHCP (enabled|disabled)
Value CONFIG_DOWNLOAD_DHCP_NEXT_BOOT (enabled|disabled)
Value BOOT_MODE (\S+)
Value BOOT_OPTIMIZATION (\S+)
Value CURRENT_BOOT_VARIABLES (BOOT\s*variable\s*does\s*not\s*exist|BOOT\s*variable\s*=\s*\S+)

Start
  ^.+show\s+boot
  ^Current\s*Boot\s*Variables: -> CurrentBootVariables
  ^Switch\s+${SWITCH_NUMBER}
  ^BOOT\s+(path-list|variable)\s+(:|=)\s+${BOOT_PATH}
  ^Config\s+file\s+:\s+${CONFIG_FILE}
  ^Private\s+Config\s+file\s+:\s+${PRIV_CONFIG_FILE}
  ^Enable\s+Break\s+(:|=)\s+${ENABLE_BREAK}
  ^Manual\s+Boot\s+(:|=)\s+${MANUAL_BOOT}
  ^Allow\s+Dev\s+Key\s+:\s+${ALLOW_DEV_KEY}
  ^HELPER\s+path-list\s+:.*
  ^Auto\s+upgrade\s+:\s+${AUTO_UPGRADE}
  ^Auto\s+upgrade\s+path\s+:.*
  ^NVRAM/Config\s+file
  ^\s*buffer\s+size:\s+${BUFFER_SIZE}
  ^iPXE\s*Timeout\s*=\s*${TIMEOUT_CONFIG_DOWNLOAD}
  ^Timeout\s+for\s+Config
  ^\s*Download:\s+${TIMEOUT_CONFIG_DOWNLOAD}
  ^\s*Config\s+Download
  ^\s*via\s+DHCP:\s+${CONFIG_DOWNLOAD_DHCP}\s\(next\sboot:\s${CONFIG_DOWNLOAD_DHCP_NEXT_BOOT}
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^Boot\s*Mode\s*=\s*${BOOT_MODE}
  ^Boot\s+optimization\s+:\s+${BOOT_OPTIMIZATION}
  ^\s*$$
  ^-+
  ^. -> Error
  
CurrentBootVariables
  ^${CURRENT_BOOT_VARIABLES}
  ^Boot\s*Variables\s*on\s*next\s*reload: -> Start
  ^. -> Error

`