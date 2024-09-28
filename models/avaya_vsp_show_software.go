package models

type AvayaVspShowSoftware struct {
	Release_location	string	`json:"RELEASE_LOCATION"`
	Backup_release	string	`json:"BACKUP_RELEASE"`
	Primary_release	string	`json:"PRIMARY_RELEASE"`
	Auto_commit	string	`json:"AUTO_COMMIT"`
	Commit_timeout	string	`json:"COMMIT_TIMEOUT"`
}