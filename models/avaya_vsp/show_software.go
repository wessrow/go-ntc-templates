package avaya_vsp 

type ShowSoftware struct {
	Release_location	string	`json:"RELEASE_LOCATION"`
	Backup_release	string	`json:"BACKUP_RELEASE"`
	Primary_release	string	`json:"PRIMARY_RELEASE"`
	Auto_commit	string	`json:"AUTO_COMMIT"`
	Commit_timeout	string	`json:"COMMIT_TIMEOUT"`
}

var ShowSoftware_Template = `Value RELEASE_LOCATION (\S+)
Value BACKUP_RELEASE (\S+)
Value PRIMARY_RELEASE (\S+)
Value AUTO_COMMIT (\S+)
Value COMMIT_TIMEOUT (.*\w)

Start
  ^==+
  ^\s+software releases in ${RELEASE_LOCATION}
  ^${BACKUP_RELEASE}\s+\(Backup Release\)
  ^${PRIMARY_RELEASE}\s+\(Primary Release\)
  ^Auto Commit\s+:\s+${AUTO_COMMIT}
  ^Commit Timeout\s+:\s+${COMMIT_TIMEOUT}

Done
`