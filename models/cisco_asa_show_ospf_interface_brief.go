package models

type CiscoAsaShowOspfInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Process	string	`json:"PROCESS"`
	Area	string	`json:"AREA"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Cost	string	`json:"COST"`
	State	string	`json:"STATE"`
	Neighbors_fc	string	`json:"NEIGHBORS_FC"`
}

var CiscoAsaShowOspfInterfaceBrief_Template = `Value INTERFACE (\S+)
Value PROCESS (\d+)
Value AREA (\S+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value NETMASK (\d+\.\d+\.\d+\.\d+)
Value COST (\d+)
Value STATE (\S+)
Value NEIGHBORS_FC (\d+/\d+)

Start
  ^Interface\s+PID\s+Area\s+IP\s+Address/Mask\s+Cost\s+State\s+Nbrs\s+F/C\s*$$
  ^${INTERFACE}\s+${PROCESS}\s+${AREA}\s+${IP_ADDRESS}\/${NETMASK}\s+${COST}\s+${STATE}\s+${NEIGHBORS_FC} -> Record
  ^\s*$$
  ^. -> Error

`