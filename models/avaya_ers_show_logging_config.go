package models

type AvayaErsShowLoggingConfig struct {
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