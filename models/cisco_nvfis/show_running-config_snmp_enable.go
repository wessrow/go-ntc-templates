package cisco_nvfis

type ShowRunningConfigSnmpEnable struct {
	Trap   string `json:"TRAP"`
	Status string `json:"STATUS"`
}

var ShowRunningConfigSnmpEnable_Template string = `Value TRAP (linkUp|linkDown)
Value STATUS (\w+)


Start
  ^snmp\s${STATUS}\straps\s${TRAP} -> Record
  ^snmp\s${STATUS}\straps\s${TRAP}-> Record
  
EOF`
