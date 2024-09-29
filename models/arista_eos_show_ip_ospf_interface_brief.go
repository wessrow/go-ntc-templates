package models

type AristaEosShowIpOspfInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Instance	string	`json:"INSTANCE"`
	Vrf	string	`json:"VRF"`
	Area	string	`json:"AREA"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Cost	string	`json:"COST"`
	State	string	`json:"STATE"`
	Neighbors	string	`json:"NEIGHBORS"`
}

var AristaEosShowIpOspfInterfaceBrief_Template = `Value INTERFACE (\S+)
Value INSTANCE (\d+)
Value VRF (\S+)
Value AREA (\d+\.\d+\.\d+\.\d+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value PREFIX_LENGTH (\d+)
Value COST (\d+)
Value STATE (\S+)
Value NEIGHBORS (\d+)

Start
  ^\s*Interface\s+Instance VRF\s+Area\s+IP Address\s+Cost\s+State\s+Nbrs\s*$$
  ^\s*${INTERFACE}\s+${INSTANCE}\s+${VRF}\s+${AREA}\s+${IP_ADDRESS}\/${PREFIX_LENGTH}\s+${COST}\s+${STATE}\s+${NEIGHBORS}\s*$$ -> Record

`