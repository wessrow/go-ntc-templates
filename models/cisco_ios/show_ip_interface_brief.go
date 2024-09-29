package cisco_ios 

type ShowIpInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Proto	string	`json:"PROTO"`
}

var ShowIpInterfaceBrief_Template = `Value INTERFACE (\S+)
Value IP_ADDRESS (\S+)
Value STATUS (up|down|administratively down)
Value PROTO (up|down)

Start
  ^${INTERFACE}\s+${IP_ADDRESS}\s+\w+\s+\w+\s+${STATUS}\s+${PROTO} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`