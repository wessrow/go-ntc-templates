package zyxel_os 

type CfgFirewallAclGet struct {
	Index	string	`json:"INDEX"`
	Enable	string	`json:"ENABLE"`
	Order	string	`json:"ORDER"`
	Name	string	`json:"NAME"`
	Source_ip	string	`json:"SOURCE_IP"`
	Destination_ip	string	`json:"DESTINATION_IP"`
	Protocol	string	`json:"PROTOCOL"`
	Ports	string	`json:"PORTS"`
	Action	string	`json:"ACTION"`
}

var CfgFirewallAclGet_Template = `Value INDEX (\d+)
Value ENABLE (true|false)
Value ORDER (\d+)
Value NAME (.+\S)
Value SOURCE_IP ((Any|\d+\.\d+\.\d+\.\d+)/\d+)
Value DESTINATION_IP ((Any|\d+\.\d+\.\d+\.\d+)/\d+)
Value PROTOCOL (ALL|ICMP|TCP|UDP|TCP/UDP)
Value PORTS ((\d+|Any)(:\d+)?)
Value ACTION (Accept|Drop)

Start
  ^Index\s+(Enable\s+)?Order\s+Name\s+Source\sIP\s+Destination\sIP\s+Service\s+Action\s*$$ -> FIREWALLTable
  ^\s*$$
  ^. -> Error

FIREWALLTable
  ^${INDEX}\s+(${ENABLE}\s+)?${ORDER}\s+${NAME}\s+${SOURCE_IP}\s+${DESTINATION_IP}\s+${PROTOCOL}(:Any-->${PORTS})?\s+${ACTION}\s*$$ -> Record
  ^Command\sSuccessful.\s*$$
  ^\s*$$
  ^. -> Error
`