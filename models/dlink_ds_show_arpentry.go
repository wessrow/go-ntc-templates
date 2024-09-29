package models

type DlinkDsShowArpentry struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
}

var DlinkDsShowArpentry_Template = `Value INTERFACE (\S+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value MAC_ADDRESS (\S+)
Value TYPE (\S+)

Start
  ^\s+ARP\s+Aging\s+Time\s+:\s+\d+
  ^Interface\s+IP Address\s+MAC Address\s+Type
  ^\-+\s+\-+\s+\-+\s+\-+$$
  ^Total Entries:\s+\d+
  ^${INTERFACE}\s+${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${TYPE} -> Record
  ^\s*$$
  ^. -> Error

`