package models

type AvayaErsShowSysInfo struct {
	Operation_mode	string	`json:"OPERATION_MODE"`
	Size_of_stack	string	`json:"SIZE_OF_STACK"`
	Base_unit	string	`json:"BASE_UNIT"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Poe_module_fw	string	`json:"POE_MODULE_FW"`
	Reset_count	string	`json:"RESET_COUNT"`
	Last_reset_type	string	`json:"LAST_RESET_TYPE"`
	Autotopology	string	`json:"AUTOTOPOLOGY"`
	Base_unit_selection	string	`json:"BASE_UNIT_SELECTION"`
	Serial_number	string	`json:"SERIAL_NUMBER"`
	Operational_firmware	string	`json:"OPERATIONAL_FIRMWARE"`
	Operational_software	string	`json:"OPERATIONAL_SOFTWARE"`
	Installed_firmware	string	`json:"INSTALLED_FIRMWARE"`
	Installed_software	string	`json:"INSTALLED_SOFTWARE"`
	Operational_license	string	`json:"OPERATIONAL_LICENSE"`
	Installed_license	string	`json:"INSTALLED_LICENSE"`
	Sys_object_id	string	`json:"SYS_OBJECT_ID"`
	Sys_up_time	string	`json:"SYS_UP_TIME"`
	Sys_ntp_time	string	`json:"SYS_NTP_TIME"`
	Sys_services	string	`json:"SYS_SERVICES"`
	Sys_contact	string	`json:"SYS_CONTACT"`
	Sys_name	string	`json:"SYS_NAME"`
	Sys_location	string	`json:"SYS_LOCATION"`
	Stack_sysassetid	string	`json:"STACK_SYSASSETID"`
	Unit_sysassetid	string	`json:"UNIT_SYSASSETID"`
}