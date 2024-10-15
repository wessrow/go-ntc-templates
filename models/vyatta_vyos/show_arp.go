package vyatta_vyos

type ShowArp struct {
	Ip_address  string `json:"IP_ADDRESS"`
	Type        string `json:"TYPE"`
	Mac_address string `json:"MAC_ADDRESS"`
	Flags       string `json:"FLAGS"`
	Interface   string `json:"INTERFACE"`
}

var ShowArp_Template string = `Value Required IP_ADDRESS ([A-Fa-f0-9:\.]+)
Value TYPE (\S+)
Value MAC_ADDRESS (\S+)
Value FLAGS (\S+)
Value INTERFACE (\S+)

Start
  ^${IP_ADDRESS}\s+(${TYPE}|)\s+${MAC_ADDRESS}\s+(${FLAGS}|)\s+${INTERFACE}$$ -> Record
  ^Address\s+HWtype\s+HWaddress\s+Flags\s+Mask\s+Iface\s*$$
  ^\s*$$
  ^. -> Error
`
