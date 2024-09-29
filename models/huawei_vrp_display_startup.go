package models

type HuaweiVrpDisplayStartup struct {
	Startup_system_software	string	`json:"STARTUP_SYSTEM_SOFTWARE"`
	Next_startup_system_software	string	`json:"NEXT_STARTUP_SYSTEM_SOFTWARE"`
	Backup_system_software_for_next_startup	string	`json:"BACKUP_SYSTEM_SOFTWARE_FOR_NEXT_STARTUP"`
	Startup_saved_configuration_file	string	`json:"STARTUP_SAVED_CONFIGURATION_FILE"`
	Next_startup_saved_configuration_file	string	`json:"NEXT_STARTUP_SAVED_CONFIGURATION_FILE"`
	Startup_license_file	string	`json:"STARTUP_LICENSE_FILE"`
	Next_startup_license_file	string	`json:"NEXT_STARTUP_LICENSE_FILE"`
	Startup_patch_package	string	`json:"STARTUP_PATCH_PACKAGE"`
	Next_startup_patch_package	string	`json:"NEXT_STARTUP_PATCH_PACKAGE"`
	Startup_voice_files	string	`json:"STARTUP_VOICE_FILES"`
	Next_startup_voice_files	string	`json:"NEXT_STARTUP_VOICE_FILES"`
}

var HuaweiVrpDisplayStartup_Template = `Value STARTUP_SYSTEM_SOFTWARE (.+)
Value NEXT_STARTUP_SYSTEM_SOFTWARE (.+)
Value BACKUP_SYSTEM_SOFTWARE_FOR_NEXT_STARTUP (.+)
Value STARTUP_SAVED_CONFIGURATION_FILE (.+)
Value NEXT_STARTUP_SAVED_CONFIGURATION_FILE (.+)
Value STARTUP_LICENSE_FILE (.+)
Value NEXT_STARTUP_LICENSE_FILE (.+)
Value STARTUP_PATCH_PACKAGE (.+)
Value NEXT_STARTUP_PATCH_PACKAGE (.+)
Value STARTUP_VOICE_FILES (.+)
Value NEXT_STARTUP_VOICE_FILES (.+)

Start
  ^\S+.*:\s*$$ -> Continue.Record
  ^\S+.*:\s*$$
  ^\s+Startup\ssystem\ssoftware:\s+${STARTUP_SYSTEM_SOFTWARE}\s*$$\s*$$
  ^\s+Next\sstartup\s+system\s+software:\s+${NEXT_STARTUP_SYSTEM_SOFTWARE}\s*$$
  ^\s+Backup\s+system\s+software\s+for\s+next\sstartup:\s+${BACKUP_SYSTEM_SOFTWARE_FOR_NEXT_STARTUP}\s*$$
  ^\s+Startup\s+saved-configuration\s+file:\s+${STARTUP_SAVED_CONFIGURATION_FILE}\s*$$
  ^\s+Next\sstartup\s+saved-configuration\s+file:\s+${NEXT_STARTUP_SAVED_CONFIGURATION_FILE}\s*$$
  ^\s+Startup\s+license\s+file:\s+${STARTUP_LICENSE_FILE}\s*$$
  ^\s+Next\sstartup\s+license\s+file:\s+${NEXT_STARTUP_LICENSE_FILE}\s*$$
  ^\s+Startup\s+patch\s+package:\s+${STARTUP_PATCH_PACKAGE}\s*$$
  ^\s+Next\sstartup\s+patch\s+package:\s+${NEXT_STARTUP_PATCH_PACKAGE}\s*$$
  ^\s+Startup\s+voice-files:\s+${STARTUP_VOICE_FILES}\s*$$
  ^\s+Next\sstartup\s+voice-files:\s+${NEXT_STARTUP_VOICE_FILES}\s*$$
  ^. -> Error

`