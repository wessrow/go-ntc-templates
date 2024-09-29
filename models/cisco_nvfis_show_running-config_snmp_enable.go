package models

type CiscoNvfisShowRunningConfigSnmpEnable struct {
	Trap	string	`json:"TRAP"`
	Status	string	`json:"STATUS"`
}

var CiscoNvfisShowRunningConfigSnmpEnable_Template = `Value TRAP (linkUp|linkDown)
Value STATUS (\w+)


Start
  ^snmp\s${STATUS}\straps\s${TRAP} -> Record
  ^snmp\s${STATUS}\straps\s${TRAP}-> Record
  
EOF
`