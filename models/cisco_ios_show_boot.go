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