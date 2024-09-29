package zyxel_os 

type CfgStaticDhcpGet struct {
	Index	string	`json:"INDEX"`
	Status	string	`json:"STATUS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Ip_address	string	`json:"IP_ADDRESS"`
}

var CfgStaticDhcpGet_Template = `Value INDEX (\d+)
Value STATUS (0|1)
Value MAC_ADDRESS (([0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2})
Value IP_ADDRESS ((\d+\.){3}\d+)

Start
  ^Index\s+Status\s+Mac\sAddress\s+IP\sAddress\s*$$ -> DHCPTable
  ^\s*$$
  ^. -> Error

DHCPTable
  ^${INDEX}\s+${STATUS}\s+${MAC_ADDRESS}\s+${IP_ADDRESS}\s*$$ -> Record
  ^Command\sSuccessful.\s*$$
  ^\s*$$
  ^. -> Error
`