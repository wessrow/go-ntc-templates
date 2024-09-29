package models

type CiscoIosShowApCdpNeighbors struct {
	Ap_name	string	`json:"AP_NAME"`
	Ap_ip	string	`json:"AP_IP"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Neighbor_ip	string	`json:"NEIGHBOR_IP"`
	Neighbor_port	string	`json:"NEIGHBOR_PORT"`
}

var CiscoIosShowApCdpNeighbors_Template = `Value AP_NAME (\S+)
Value AP_IP ([a-fA-F0-9:\.]+)
Value NEIGHBOR_NAME (\S+)
Value NEIGHBOR_IP ([a-fA-F0-9:\.]+)
Value NEIGHBOR_PORT ([a-zA-Z0-9\/\.]+)

Start
  ^[a-zA-Z]+\s[a-z]+\s[a-zA-z]+:\s\d+$$
  ^AP\sName\s+AP\sIP\s+Neighbor\sName\s+Neighbor\sIP\s+Neighbor\sPort\s*$$
  ^-+\s*$$
  ^${AP_NAME}\s+${AP_IP}\s+${NEIGHBOR_NAME}\s+${NEIGHBOR_IP}\s+${NEIGHBOR_PORT}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`