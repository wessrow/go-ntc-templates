package avaya_ers 

type ShowLoggingConfig struct {
	Event_logging	string	`json:"EVENT_LOGGING"`
	Volatile	string	`json:"VOLATILE"`
	Log_types	string	`json:"LOG_TYPES"`
	Nv_storage_types	string	`json:"NV_STORAGE_TYPES"`
	Remote_logging	string	`json:"REMOTE_LOGGING"`
	Remote_address	string	`json:"REMOTE_ADDRESS"`
	Secondary_address	string	`json:"SECONDARY_ADDRESS"`
	Remote_types	string	`json:"REMOTE_TYPES"`
	Facility	string	`json:"FACILITY"`
}

var ShowLoggingConfig_Template = `Value EVENT_LOGGING (.+?)
Value VOLATILE (.+?)
Value LOG_TYPES (.+?)
Value NV_STORAGE_TYPES (.+?)
Value REMOTE_LOGGING (.+?)
Value REMOTE_ADDRESS (.+?)
Value SECONDARY_ADDRESS (.+?)
Value REMOTE_TYPES (.+?)
Value FACILITY (.+?)

Start
  ^Event Logging:\s+${EVENT_LOGGING}\s*$$
  ^Volatile Logging Option:\s+${VOLATILE}\s*$$
  ^Event Types To Log:\s+${LOG_TYPES}\s*$$
  ^Event Types To Log To NV Storage:\s+${NV_STORAGE_TYPES}\s*$$
  ^Remote Logging:\s+${REMOTE_LOGGING}\s*$$
  ^Remote Logging Address:\s+${REMOTE_ADDRESS}\s*$$
  ^Secondary Remote Logging Address:\s+${SECONDARY_ADDRESS}\s*$$
  ^Event Types To Log Remotely:\s+${REMOTE_TYPES}\s*$$
  ^Facility:\s+${FACILITY}\s*$$
  ^\s*$$
  ^. -> Error
`