package cisco_nvfis

type ShowRunningConfigSystemTime struct {
	Timezone                  string `json:"TIMEZONE"`
	Preferred_ntp_server      string `json:"PREFERRED_NTP_SERVER"`
	Preferred_ntp_server_type string `json:"PREFERRED_NTP_SERVER_TYPE"`
	Backup_ntp_server         string `json:"BACKUP_NTP_SERVER"`
	Backup_ntp_server_type    string `json:"BACKUP_NTP_SERVER_TYPE"`
}

var ShowRunningConfigSystemTime_Template string = `Value TIMEZONE (\w+)
Value PREFERRED_NTP_SERVER ([A-F0-9\.\:]+)
Value PREFERRED_NTP_SERVER_TYPE (\w+)
Value BACKUP_NTP_SERVER ([A-F0-9\.\:]+)
Value BACKUP_NTP_SERVER_TYPE (\w+)


Start
  ^system\stime\stimezone\s${TIMEZONE}
  ^system\stime\sntp\spreferred_server\s${PREFERRED_NTP_SERVER}
  ^system\stime\sntp\sbackup_server\s${BACKUP_NTP_SERVER}
  ^system\stime\sntp\spreferred_server_type\s${PREFERRED_NTP_SERVER_TYPE}
  ^system\stime\sntp\sbackup_server_type\s${BACKUP_NTP_SERVER_TYPE} -> Record

EOF`
