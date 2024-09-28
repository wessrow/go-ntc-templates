package models

type CiscoNvfisShowRunningConfigSystemTime struct {
	Timezone	string	`json:"TIMEZONE"`
	Preferred_ntp_server	string	`json:"PREFERRED_NTP_SERVER"`
	Preferred_ntp_server_type	string	`json:"PREFERRED_NTP_SERVER_TYPE"`
	Backup_ntp_server	string	`json:"BACKUP_NTP_SERVER"`
	Backup_ntp_server_type	string	`json:"BACKUP_NTP_SERVER_TYPE"`
}