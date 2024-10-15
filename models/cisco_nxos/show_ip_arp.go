package cisco_nxos

type ShowIpArp struct {
	Ip_address  string `json:"IP_ADDRESS"`
	Age         string `json:"AGE"`
	Mac_address string `json:"MAC_ADDRESS"`
	Interface   string `json:"INTERFACE"`
}

var ShowIpArp_Template string = `Value Required IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Required AGE (\S+)
Value Required MAC_ADDRESS (\S+)
Value INTERFACE (\S+)

Start
  ^Address\s+Age\s+MAC Address\s+Interface -> Start_record

Start_record
  ^${IP_ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${INTERFACE} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
`
