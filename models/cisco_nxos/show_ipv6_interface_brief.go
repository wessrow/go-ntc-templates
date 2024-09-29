package cisco_nxos 

type ShowIpv6InterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Linkipaddr	string	`json:"LINKIPADDR"`
	Status	string	`json:"STATUS"`
	Proto	string	`json:"PROTO"`
}

var ShowIpv6InterfaceBrief_Template = `Value Required INTERFACE (\S+)
Value Required IPV6_ADDRESS ([a-zA-Z0-9:/]+)
Value LINKIPADDR ([a-zA-Z0-9:/]+)
Value STATUS (\S+/\S+)
Value PROTO (\S+)

Start
  ^${INTERFACE}\s+${IPV6_ADDRESS}\s+${STATUS}/${PROTO}
  ^\s+${LINKIPADDR} -> Record

`