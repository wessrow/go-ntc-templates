package models

type CiscoXrShowOspfVrfAllInterfaceBrief struct {
	Vrf	string	`json:"VRF"`
	Interface	string	`json:"INTERFACE"`
	Process	string	`json:"PROCESS"`
	Area	string	`json:"AREA"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Cost	string	`json:"COST"`
	State	string	`json:"STATE"`
	Neighbors	string	`json:"NEIGHBORS"`
}

var CiscoXrShowOspfVrfAllInterfaceBrief_Template = `Value Filldown VRF (\S+)
Value Required INTERFACE ([\w\./-]+)
Value PROCESS (\d+)
Value AREA ([\d\.]+)
Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value PREFIX_LENGTH (\d+)
Value COST (\d+)
Value STATE (\S+)
Value NEIGHBORS (\d+/\d+)

Start
  ^\s*Interfaces for OSPF \d+, VRF ${VRF}\s*$$
  ^\s+Interface\s+PID\s+Area\s+IP Address/Mask\s+Cost\s+State\s+Nbrs\s+F/C\s*$$
  ^\s*${INTERFACE}\s+${PROCESS}\s+${AREA}\s+${IP_ADDRESS}\/${PREFIX_LENGTH}\s+${COST}\s+${STATE}\s+${NEIGHBORS}\s*$$ -> Record

`