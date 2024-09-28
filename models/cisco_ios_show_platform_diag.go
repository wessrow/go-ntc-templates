package models

type CiscoIosShowPlatformDiag struct {
	Chassis_type	string	`json:"CHASSIS_TYPE"`
	Slot_number	string	`json:"SLOT_NUMBER"`
	Module_sku	string	`json:"MODULE_SKU"`
	State	string	`json:"STATE"`
	Running_state	string	`json:"RUNNING_STATE"`
	Internal_state	string	`json:"INTERNAL_STATE"`
	Internal_operational_state	string	`json:"INTERNAL_OPERATIONAL_STATE"`
	Insert_time	string	`json:"INSERT_TIME"`
	Uptime	string	`json:"UPTIME"`
	Hardware_signal	string	`json:"HARDWARE_SIGNAL"`
	Packet_signal	string	`json:"PACKET_SIGNAL"`
	Cpld_version	string	`json:"CPLD_VERSION"`
	Firmware_version	string	`json:"FIRMWARE_VERSION"`
}