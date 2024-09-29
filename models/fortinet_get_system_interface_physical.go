package models

type FortinetGetSystemInterfacePhysical struct {
	Name	string	`json:"NAME"`
	Mode	string	`json:"MODE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Ipv6netmask	string	`json:"IPV6NETMASK"`
	Status	string	`json:"STATUS"`
	Speed	string	`json:"SPEED"`
	Duplex	string	`json:"DUPLEX"`
}

var FortinetGetSystemInterfacePhysical_Template = `#
# Refried Jello, Stephen Kiely
# 2021-01-06 - Initial Version
#
# FG Version: 5.6, 6.0, 6.2
# HW        : variety
# HA        : AP
# VDOMS     : ENABLED
Value Required NAME (\S+)
Value MODE (\S+)
Value IP_ADDRESS (\d+?\.\d+?\.\d+?\.\d+?)
Value NETMASK (\S+)
Value IPV6_ADDRESS (\S+)
Value IPV6NETMASK (\S+)
Value STATUS (\S+)
Value SPEED (\d+|n\/a)
Value DUPLEX (\S+)

Start
  ^==\s+\[onboard\]$$
  ^\s+==\[${NAME}\]$$
  ^\s+mode:\s+${MODE}$$
  ^\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}$$
  ^\s+ipv6:\s+${IPV6_ADDRESS}/${IPV6NETMASK}$$
  ^\s+status:\s+${STATUS}$$
  ^\s+speed:\s+${SPEED}(|Mbps)\s+\(Duplex:\s+${DUPLEX}\)$$ -> Record
  ^\s+speed:\s+${SPEED}$$ -> Record
  ^\s*$$
  ^. -> Error

`