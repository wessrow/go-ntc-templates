package fortinet_execute 

type DhcpLeaseList struct {
	Interface	string	`json:"INTERFACE"`
	Ip	string	`json:"IP"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Hostname	string	`json:"HOSTNAME"`
	Vci	string	`json:"VCI"`
	Ssid	string	`json:"SSID"`
	Ap	string	`json:"AP"`
	Expiry	string	`json:"EXPIRY"`
}

var DhcpLeaseList_Template = `Value Filldown INTERFACE (\S+)
Value Required IP (\S+)
Value MAC_ADDRESS (([a-f0-9]{2}:){5}[a-f0-9]{2})
Value HOSTNAME ([^ ].+?(?=\s{2}))
Value VCI ([^ ].+?(?=\s{2}))
Value SSID (\S+)
Value AP (\S+)
Value EXPIRY ((Mon|Tue|Wed|Thu|Fri|Sat|Sun)\s+(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\s+\d+\s+\d+\:\d+\:\d+\s+\d+)

Start
  ^\s*${INTERFACE}\s*$$
  ^\s*IP\s+MAC-Address\s+Hostname\s+VCI\s+SSID\s+AP\s+Expiry\s*$$ -> SevenColumns
  ^\s*IP\s+MAC-Address\s+Hostname\s+VCI\s+Expiry\s*$$ -> FiveColumns
  ^\s*$$
  ^. -> Error

SevenColumns
  ^\s*IP\s+MAC-Address\s+Hostname\s+VCI\s+SSID\s+AP\s+Expiry\s*$$
  ^\s*${INTERFACE}\s*$$
  ^\s*${IP}\s+${MAC_ADDRESS}\s+${HOSTNAME}(\s+${VCI})?(\s+${SSID})?(\s+${AP})?\s+${EXPIRY}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

FiveColumns
  ^\s*IP\s+MAC-Address\s+Hostname\s+VCI\s+Expiry\s*$$
  ^\s*${INTERFACE}\s*$$
  ^\s*${IP}\s+${MAC_ADDRESS}\s+${HOSTNAME}(\s+${VCI})?\s+${EXPIRY}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`