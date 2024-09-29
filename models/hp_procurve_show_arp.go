package models

type HpProcurveShowArp struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
	Port	string	`json:"PORT"`
}

var HpProcurveShowArp_Template = `Value IP_ADDRESS (\S+)
Value MAC_ADDRESS ([0-9a-fA-F]{6}-[0-9a-fA-F]{6})
Value TYPE (\S+)
Value PORT (\S+)

Start
  ^.*IP ARP table
  ^.*IP Address -> ARP
  ^\s*$$
  ^. -> Error

ARP
  ^\s+${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${TYPE}\s+${PORT} -> Record
  ^\s+${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${TYPE} -> Record
  ^\s*------
  ^\s*$$
  ^. -> Error

`