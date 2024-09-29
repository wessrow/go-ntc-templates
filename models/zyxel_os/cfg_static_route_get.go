package zyxel_os 

type CfgStaticRouteGet struct {
	Index	string	`json:"INDEX"`
	Name	string	`json:"NAME"`
	Enabled	string	`json:"ENABLED"`
	Ip_version	string	`json:"IP_VERSION"`
	Destination	string	`json:"DESTINATION"`
	Netmask	string	`json:"NETMASK"`
	Interface	string	`json:"INTERFACE"`
	Gateway	string	`json:"GATEWAY"`
}

var CfgStaticRouteGet_Template = `Value INDEX (\d+)
Value NAME (.+?)
Value ENABLED (0|1)
Value IP_VERSION (IPv4|IPv6)
Value DESTINATION (((([0-9]{1,3}\.){3}[0-9]{1,3}(/\d{1-2})?)|[0-9a-f:]+/\d{1,3})?)
Value NETMASK ((([0-9]{1,3}\.){3}[0-9]{1,3})?)
Value INTERFACE (\S+?)
Value GATEWAY (((([0-9]{1,3}\.){3}[0-9]{1,3}|[0-9a-f:]+)?)?)

Start
  ^Index\s+Name\s+Enable\s+IPver\s+DestIPAddress/DestPrefix\s+DestMask\s+Interface\s+Gateway/NextHop\s*$$ -> ROUTETable
  ^\s*$$
  ^. -> Error

ROUTETable
  ^${INDEX}\s+${NAME}\s+${ENABLED}\s+${IP_VERSION}\s+${DESTINATION}\s+${NETMASK}\s+${INTERFACE}\s+${GATEWAY}\s*$$ -> Record
  ^Command Successful.\s*$$
  ^\s*$$
  ^. -> Error
`