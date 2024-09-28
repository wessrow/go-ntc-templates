package models

type CiscoNvfisShowVersion struct {
	Version	string	`json:"VERSION"`
	Build_date	string	`json:"BUILD_DATE"`
	Last_reboot	string	`json:"LAST_REBOOT"`
}