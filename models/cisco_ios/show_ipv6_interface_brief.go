package cisco_ios 

type ShowIpv6InterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ipv6_address	[]string	`json:"IPV6_ADDRESS"`
	Admin	string	`json:"ADMIN"`
	Protocol	string	`json:"PROTOCOL"`
}

var ShowIpv6InterfaceBrief_Template = `Value INTERFACE (\S+)
Value List IPV6_ADDRESS (\S+)
Value ADMIN (\S+|\S+\s\S+)
Value PROTOCOL (\S+|\S+\s\S+)


Start
  ^${INTERFACE}\s+\[${ADMIN}/${PROTOCOL}\] -> Interfaces
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

Interfaces
  ^(\S+)\s+\[(\S+|\S+\s\S+)/(\S+|\S+\s\S+)\] -> Continue.Record
  ^${INTERFACE}\s+\[${ADMIN}/${PROTOCOL}\]
  ^\s+${IPV6_ADDRESS}

`