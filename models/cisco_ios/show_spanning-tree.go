package cisco_ios

type ShowSpanningTree struct {
	Status        string `json:"STATUS"`
	Cost          string `json:"COST"`
	Port_priority string `json:"PORT_PRIORITY"`
	Port_id       string `json:"PORT_ID"`
	Type          string `json:"TYPE"`
	Vlan_id       string `json:"VLAN_ID"`
	Interface     string `json:"INTERFACE"`
	Role          string `json:"ROLE"`
}

var ShowSpanningTree_Template string = `Value Filldown VLAN_ID (\d+)
Value Required INTERFACE (\S+)
Value ROLE (\w+)
Value STATUS (\w+)
Value COST (\d+)
Value PORT_PRIORITY (\w+)
Value PORT_ID (\w+)
Value TYPE (.*)

Start
  ^VLAN(0*)?${VLAN_ID}
  ^${INTERFACE}\s+${ROLE}\s+${STATUS}\s+${COST}\s+${PORT_PRIORITY}.${PORT_ID}\s+${TYPE} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`
