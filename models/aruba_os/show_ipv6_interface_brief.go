package aruba_os 

type ShowIpv6InterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ipv6_address	[]string	`json:"IPV6_ADDRESS"`
	Admin	string	`json:"ADMIN"`
	Protocol	string	`json:"PROTOCOL"`
}

var ShowIpv6InterfaceBrief_Template = `Value INTERFACE (\S+\s\S+|\S+)
Value List IPV6_ADDRESS (\S+)
Value ADMIN (\S+)
Value PROTOCOL (\S+)

Start
  ^${INTERFACE}\s+\[\s+${ADMIN}/${PROTOCOL}\s+\] -> Interfaces

Interfaces
  ^(\S+\s\S+|\S+)\s+\[\s+(\S+)/(\S+)\s+\] -> Continue.Record
  ^${INTERFACE}\s+\[\s+${ADMIN}/${PROTOCOL}\s+\]
  ^\s+${IPV6_ADDRESS}
`