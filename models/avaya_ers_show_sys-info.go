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

var AvayaErsShowSysInfo_Template = `Value OPERATION_MODE (.*\w)
Value SIZE_OF_STACK (\d)
Value BASE_UNIT (\d)
Value MAC_ADDRESS (.*\w)
Value POE_MODULE_FW (.*\w)
Value RESET_COUNT (\d+)
Value LAST_RESET_TYPE (.*\w)
Value AUTOTOPOLOGY (.*\w)
Value BASE_UNIT_SELECTION (.*\w)
Value SERIAL_NUMBER (.*\w)
Value OPERATIONAL_FIRMWARE (.*\w)
Value OPERATIONAL_SOFTWARE (.*\w)
Value INSTALLED_FIRMWARE (.*\w)
Value INSTALLED_SOFTWARE (.*\w)
Value OPERATIONAL_LICENSE (.*\w)
Value INSTALLED_LICENSE (.*\w)
Value SYS_OBJECT_ID (.*\w)
Value SYS_UP_TIME (.*\w)
Value SYS_NTP_TIME (.*\w)
Value SYS_SERVICES (\d+)
Value SYS_CONTACT (.*\w)
Value SYS_NAME (.*\w)
Value SYS_LOCATION (.*\w)
Value STACK_SYSASSETID (.*\w)
Value UNIT_SYSASSETID (.*\w)

Start
  ^Operation Mode:\s+${OPERATION_MODE}(\s+|$$)
  ^Size Of Stack:\s+${SIZE_OF_STACK}(\s+|$$)
  ^Base Unit:\s+${BASE_UNIT}(\s+|$$)
  ^MAC Address:\s+${MAC_ADDRESS}(\s+|$$)
  ^PoE Module FW:\s+${POE_MODULE_FW}(\s+|$$)
  ^Reset Count:\s+${RESET_COUNT}(\s+|$$)
  ^Last Reset Type:\s+${LAST_RESET_TYPE}(\s+|$$)
  ^Autotopology:\s+${AUTOTOPOLOGY}(\s+|$$)
  ^Base Unit Selection:\s+${BASE_UNIT_SELECTION}(\s+|$$)
  ^Serial #:\s+${SERIAL_NUMBER}(\s+|$$)
  ^Operational Software:\s+FW:${OPERATIONAL_FIRMWARE}\s+SW:${OPERATIONAL_SOFTWARE}(\s+|$$)
  ^Installed software:\s+FW:${INSTALLED_FIRMWARE}\s+SW:${INSTALLED_SOFTWARE}(\s+|$$)
  ^Operational license:\s+${OPERATIONAL_LICENSE}(\s+|$$)
  ^Installed license:\s+${INSTALLED_LICENSE}(\s+|$$)
  ^sysObjectID:\s+${SYS_OBJECT_ID}(\s+|$$)
  ^sysUpTime:\s+${SYS_UP_TIME}(\s+|$$)
  ^sysNtpTime:\s+${SYS_NTP_TIME}(\s+|$$)
  ^sysServices:\s+${SYS_SERVICES}(\s+|$$)
  ^sysContact:\s+${SYS_CONTACT}(\s+|$$)
  ^sysName:\s+${SYS_NAME}(\s+|$$)
  ^sysLocation:\s+${SYS_LOCATION}(\s+|$$)
  ^Stack sysAssetId:\s+${STACK_SYSASSETID}(\s+|$$)
  ^Unit sysAssetId:\s+${UNIT_SYSASSETID}(\s+|$$) -> Record

`