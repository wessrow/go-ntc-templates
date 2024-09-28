package models

type AristaEosShowInterfacesTransceiverDetail struct {
	Port	string	`json:"PORT"`
	Type	string	`json:"TYPE"`
	Current_value	string	`json:"CURRENT_VALUE"`
	High_alarm_threshold	string	`json:"HIGH_ALARM_THRESHOLD"`
	High_warn_threshold	string	`json:"HIGH_WARN_THRESHOLD"`
	Low_alarm_threshold	string	`json:"LOW_ALARM_THRESHOLD"`
	Low_warn_threshold	string	`json:"LOW_WARN_THRESHOLD"`
}