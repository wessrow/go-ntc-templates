package oneaccess_oneos 

type ShowRebootCounters struct {
	Device	string	`json:"DEVICE"`
	Serial	string	`json:"SERIAL"`
	Reboot_cause	string	`json:"REBOOT_CAUSE"`
	Hardware_reset	string	`json:"HARDWARE_RESET"`
	Power_fail_detection	string	`json:"POWER_FAIL_DETECTION"`
	Total_software_reboots	string	`json:"TOTAL_SOFTWARE_REBOOTS"`
	System_defense	string	`json:"SYSTEM_DEFENSE"`
	Generic_software_reboot	string	`json:"GENERIC_SOFTWARE_REBOOT"`
	Admin_requested_reboot	string	`json:"ADMIN_REQUESTED_REBOOT"`
	Admin_delayed_reboot	string	`json:"ADMIN_DELAYED_REBOOT"`
	Spurious_power_fails	string	`json:"SPURIOUS_POWER_FAILS"`
}

var ShowRebootCounters_Template = `Value Required DEVICE (\S+)
Value SERIAL (\S+)
Value REBOOT_CAUSE (.*)
Value HARDWARE_RESET (\d+)
Value POWER_FAIL_DETECTION (\d+)
Value TOTAL_SOFTWARE_REBOOTS (\d+)
Value SYSTEM_DEFENSE (\d+)
Value GENERIC_SOFTWARE_REBOOT (\d+)
Value ADMIN_REQUESTED_REBOOT (\d+)
Value ADMIN_DELAYED_REBOOT (\d+)
Value SPURIOUS_POWER_FAILS (\d+)

Start
  ^Reboot\sstatus\sfor\sdevice\s${DEVICE}\sS\/N\s${SERIAL}
  ^Last\sReboot\sCause\s+:\s+${REBOOT_CAUSE}
  ^Reboot\sCounters
  ^Reboot\son\shardware\sreset\s+:\s+${HARDWARE_RESET}
  ^Power\sFail\sdetection\s+:\s+${POWER_FAIL_DETECTION}
  ^Total\sSoftware\sRequested\sReboots\s+:\s+${TOTAL_SOFTWARE_REBOOTS}
  ^\s*Generic\ssoftware\sreboot\srequest\s+:\s+${GENERIC_SOFTWARE_REBOOT}
  ^\s*System\sdefense\s-\sreboot\safter\scrash\s+:\s+${SYSTEM_DEFENSE}
  ^\s*Administrator\srequested\sreboot\s+:\s+${ADMIN_REQUESTED_REBOOT}
  ^\s*Administrator\srequested\sdelayed\sreboot\s+:\s+${ADMIN_DELAYED_REBOOT}
  ^Since\sstartup,\sspurious\sPower\sFail\shas\soccur+ed\s${SPURIOUS_POWER_FAILS}\stime
  ^\s*$$
  ^. -> Error
`