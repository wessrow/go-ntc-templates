package edgecore_show 

type IpInterface struct {
	Ip	string	`json:"IP"`
	Netmask	string	`json:"NETMASK"`
	Interface	string	`json:"INTERFACE"`
	Addr_mode	string	`json:"ADDR_MODE"`
}

var IpInterface_Template = `Value IP (\S+)
Value NETMASK (\S+)
Value INTERFACE (.*)
Value ADDR_MODE (.*)

Start
  ^\s*IP\s+Address\s+and\s+Netmask:.*$$ -> Continue.Record
  ^\s*IP\s+Address\s+and\s+Netmask:\s+${IP}\s+${NETMASK}\s+on\s+${INTERFACE},\s*$$
  ^\s*Address\s+Mode:\s+${ADDR_MODE}\s*$$
  ^\s*$$
  ^. -> Error
`