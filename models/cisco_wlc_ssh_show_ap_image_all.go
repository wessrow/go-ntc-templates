package models

type CiscoWlcSshShowApImageAll struct {
	Ap_name	string	`json:"AP_NAME"`
	Primary_image	string	`json:"PRIMARY_IMAGE"`
	Backup_image	string	`json:"BACKUP_IMAGE"`
	Predownload_status	string	`json:"PREDOWNLOAD_STATUS"`
	Predownload_version	string	`json:"PREDOWNLOAD_VERSION"`
	Next_retry_time	string	`json:"NEXT_RETRY_TIME"`
	Retry_count	string	`json:"RETRY_COUNT"`
	Flexcon_predown	string	`json:"FLEXCON_PREDOWN"`
}

var CiscoWlcSshShowApImageAll_Template = `Value AP_NAME (\S+)
Value PRIMARY_IMAGE ([0-9\.]+)
Value BACKUP_IMAGE ([0-9\.]+)
Value PREDOWNLOAD_STATUS (\S+)
Value PREDOWNLOAD_VERSION (\S+)
Value NEXT_RETRY_TIME (\S+)
Value RETRY_COUNT (\S+)
Value FLEXCON_PREDOWN (\s+)


Start
  ^${AP_NAME}\s+${PRIMARY_IMAGE}\s+${BACKUP_IMAGE}\s+${PREDOWNLOAD_STATUS}\s+${PREDOWNLOAD_VERSION}\s+${NEXT_RETRY_TIME}\s+${RETRY_COUNT}\s*$$ -> Record
  ^Total\s+number\s+of\s+APs
  ^Number\s+of\s+APs
  ^.+\.+
  ^\s*$$
  ^\s*Predownload\s+Predownload\s+Flexconnect
  ^AP\s+Name\s+Primary\s+Image\s+Backup\s+Image\s+Status\s+Version\s+Next\s+Retry\s+Time\s+Retry\s+Count\s+Predownload\s*$$
  ^-+
  ^. -> Error

`