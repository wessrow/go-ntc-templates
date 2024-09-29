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

var CiscoIosShowPlatformDiag_Template = `Value Filldown CHASSIS_TYPE (.*?)
Value Required SLOT_NUMBER (.*?)
Value MODULE_SKU (.+?)
Value STATE (.+)
Value RUNNING_STATE (.+)
Value INTERNAL_STATE (.+)
Value INTERNAL_OPERATIONAL_STATE (.+)
Value INSERT_TIME (.*)
Value UPTIME (.*)
Value HARDWARE_SIGNAL (.*)
Value PACKET_SIGNAL (.*)
Value CPLD_VERSION (.*)
Value FIRMWARE_VERSION (.*)

Start
  ^Chassis type: ${CHASSIS_TYPE}(?:\s|$$)
  ^.*(?:Sub-slot|Slot):\s${SLOT_NUMBER}\,(?:\s${MODULE_SKU}|\s+$$)(?:\s+|$$)
  ^\s+State\s*\:\s${STATE}
  ^.*Running state\s*\:\s${RUNNING_STATE}
  ^.*Internal state\s*\:\s${INTERNAL_STATE}
  ^.*Internal operational state\s*\:\s${INTERNAL_OPERATIONAL_STATE}
  ^.*Physical insert detect time\s*\:\s${INSERT_TIME}
  ^.*Software declared up time\s*\:\s${UPTIME}
  ^.*Hardware ready signal time\s*\:\s${HARDWARE_SIGNAL}
  ^.*Packet ready signal time\s*\:\s${PACKET_SIGNAL}
  ^.*CPLD version\s*\:\s${CPLD_VERSION}
  ^.*Firmware version\s*\:\s${FIRMWARE_VERSION}
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^$$ -> Record
`