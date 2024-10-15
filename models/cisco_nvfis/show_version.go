package cisco_nvfis

type ShowVersion struct {
	Version     string `json:"VERSION"`
	Build_date  string `json:"BUILD_DATE"`
	Last_reboot string `json:"LAST_REBOOT"`
}

var ShowVersion_Template string = `Value VERSION (\S+)
Value BUILD_DATE (.+)
Value LAST_REBOOT (.+)


Start
  ^Version\s${VERSION}
  ^Build\sdate\s${BUILD_DATE}
  ^Last\sReboot\s${LAST_REBOOT} -> Record
`
