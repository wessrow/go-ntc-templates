package cisco_nvfis 

type ShowRunningConfigSnmpGroup struct {
	Groupname	string	`json:"GROUPNAME"`
	Version	string	`json:"VERSION"`
	Priviledge	string	`json:"PRIVILEDGE"`
	Write	string	`json:"WRITE"`
	Read	string	`json:"READ"`
	Notify	string	`json:"NOTIFY"`
}

var ShowRunningConfigSnmpGroup_Template = `Value GROUPNAME ([-\w]+)
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