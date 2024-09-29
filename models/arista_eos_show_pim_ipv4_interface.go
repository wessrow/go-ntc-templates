package models

type AristaEosShowPimIpv4Interface struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Mode	string	`json:"MODE"`
	Neighbors	string	`json:"NEIGHBORS"`
	Hello_interval	string	`json:"HELLO_INTERVAL"`
	Dr_priority	string	`json:"DR_PRIORITY"`
	Dr_address	string	`json:"DR_ADDRESS"`
	Packets_queued	string	`json:"PACKETS_QUEUED"`
	Packets_dropped	string	`json:"PACKETS_DROPPED"`
}

var AristaEosShowPimIpv4Interface_Template = `Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Required INTERFACE ([\w\./-]+)
Value MODE (\S+)
Value NEIGHBORS (\d+)
Value HELLO_INTERVAL (\d+)
Value DR_PRIORITY (\d+)
Value DR_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value PACKETS_QUEUED (\d+)
Value PACKETS_DROPPED (\d+)

Start
  ^\s*Address\s+Interface\s+Mode\s+Neighbor\s+Hello\s+DR\s+DR Address\s+PktsQed\s+PktsDropped\s*$$
  ^\s*Count\s+Intvl\s+Pri\s*$$
  ^\s*${IP_ADDRESS}\s+${INTERFACE}\s+${MODE}\s+${NEIGHBORS}\s+${HELLO_INTERVAL}\s+${DR_PRIORITY}\s+${DR_ADDRESS}\s+${PACKETS_QUEUED}\s+${PACKETS_DROPPED}\s*$$ -> Record

`