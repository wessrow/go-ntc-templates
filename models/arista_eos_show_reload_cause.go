package models

type AristaEosShowReloadCause struct {
	Reload_cause	string	`json:"RELOAD_CAUSE"`
	Reload_time	string	`json:"RELOAD_TIME"`
	Recommended_action	string	`json:"RECOMMENDED_ACTION"`
	Debug_info	string	`json:"DEBUG_INFO"`
}