package cisco_ios

type ShowIpArp struct {
	Type        string `json:"TYPE"`
	Interface   string `json:"INTERFACE"`
	Protocol    string `json:"PROTOCOL"`
	Ip_address  string `json:"IP_ADDRESS"`
	Age         string `json:"AGE"`
	Mac_address string `json:"MAC_ADDRESS"`
}

var ShowIpArp_Template string = `Value Required PROTOCOL (\S+)
Value Required IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Required AGE (\S+)
Value Required MAC_ADDRESS (\S+)
Value Required TYPE (\S+)
Value INTERFACE (\S+)

Start
  ^Protocol\s+Address\s+Age\s*\(min\)\s+Hardware Addr\s+Type\s+Interface
  ^${PROTOCOL}\s+${IP_ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${TYPE}\s+${INTERFACE} -> Record
  ^${PROTOCOL}\s+${IP_ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${TYPE} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error  
`
