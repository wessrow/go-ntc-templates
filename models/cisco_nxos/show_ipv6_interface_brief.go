package cisco_nxos

type ShowIpv6InterfaceBrief struct {
	Ipv6_address string `json:"IPV6_ADDRESS"`
	Linkipaddr   string `json:"LINKIPADDR"`
	Status       string `json:"STATUS"`
	Proto        string `json:"PROTO"`
	Interface    string `json:"INTERFACE"`
}

var ShowIpv6InterfaceBrief_Template string = `Value Required INTERFACE (\S+)
Value Required IPV6_ADDRESS ([a-zA-Z0-9:/]+)
Value LINKIPADDR ([a-zA-Z0-9:/]+)
Value STATUS (\S+/\S+)
Value PROTO (\S+)

Start
  ^${INTERFACE}\s+${IPV6_ADDRESS}\s+${STATUS}/${PROTO}
  ^\s+${LINKIPADDR} -> Record

`
