package models

type ArubaAoscxShowArpAllVrfs struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Port_id	string	`json:"PORT_ID"`
	Physical_port	string	`json:"PHYSICAL_PORT"`
	State	string	`json:"STATE"`
	Vrf	string	`json:"VRF"`
}

var ArubaAoscxShowArpAllVrfs_Template = `Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value MAC_ADDRESS (\S+)
Value PORT_ID (\S+)
Value PHYSICAL_PORT (\S+)
Value STATE ([a-z]+)
Value VRF (\S+)

Start
  ^IPv4\s+Address\s+MAC\s+Port\s+Physical\s+Port\s+State
  ^${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${PORT_ID}\s+${PHYSICAL_PORT}\s+${STATE}\s+${VRF} -> Record
  ^\s*-*
  ^\s*Total.*$$
  ^. -> Error

`