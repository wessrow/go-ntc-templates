package models

type CiscoXrShowIpv4VrfAllInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Proto	string	`json:"PROTO"`
	Vrf	string	`json:"VRF"`
}

var CiscoXrShowIpv4VrfAllInterfaceBrief_Template = `Value Required INTERFACE ([\w\./-]+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value STATUS (\w+)
Value PROTO (\w+)
Value VRF (\S+)

Start
  ^\s*Interface\s+IP-Address\s+Status\s+Protocol(\s+Vrf-Name)?\s*$$
  ^\s*${INTERFACE}\s+${IP_ADDRESS}\s+${STATUS}\s+${PROTO}(\s+${VRF})?\s*$$ -> Record

`