package cisco_nvfis

type ShowRunningConfigSnmpGroup struct {
	Priviledge string `json:"PRIVILEDGE"`
	Write      string `json:"WRITE"`
	Read       string `json:"READ"`
	Notify     string `json:"NOTIFY"`
	Groupname  string `json:"GROUPNAME"`
	Version    string `json:"VERSION"`
}

var ShowRunningConfigSnmpGroup_Template string = `Value GROUPNAME ([-\w]+)
Value VERSION (\d)
Value PRIVILEDGE (\w+)
Value WRITE (\w+)
Value READ (\w+)
Value NOTIFY (\w+)


Start
  ^snmp\sgroup\s${GROUPNAME}\s\S+\s${VERSION}\s${PRIVILEDGE}
  ^\swrite\s+${WRITE}
  ^\sread\s+${READ}
  ^\snotify\s+${NOTIFY}
  ^! -> Record
  
EOF`
