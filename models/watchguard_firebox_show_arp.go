package models

type WatchguardFireboxShowArp struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
}

var WatchguardFireboxShowArp_Template = `Value INTERFACE (\S+)
Value IP_ADDRESS ([0-9.]+)
Value MAC_ADDRESS ([0-9a-fA-F:]+|incomplete)

Start
  ^--
  ^--\s+ARP\s+Table
  ^--
  ^\?\s+\(${IP_ADDRESS}\)\s+at\s+<*${MAC_ADDRESS}>*\s+(\[\S+\]\s*)?([A-Z]+\s*)?on\s+${INTERFACE} -> Record
  ^\s*$$
  ^. -> Error

`