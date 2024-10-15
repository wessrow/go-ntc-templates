package fortinet

type GetSystemArp struct {
	Ip_address  string `json:"IP_ADDRESS"`
	Age         string `json:"AGE"`
	Mac_address string `json:"MAC_ADDRESS"`
	Interface   string `json:"INTERFACE"`
}

var GetSystemArp_Template string = `Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value AGE (\d+)
Value MAC_ADDRESS (\S+)
Value INTERFACE (.*)

Start
  ^${IP_ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${INTERFACE} -> Record
  ^Address\s+Age(min)\s+Hardware\s+Addr\s+Interface\s*$$
  ^\s*
  ^. -> Error
`
