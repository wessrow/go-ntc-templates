package ubiquiti_edgerouter

type ShowArp struct {
	Hardware_type string `json:"HARDWARE_TYPE"`
	Mac_address   string `json:"MAC_ADDRESS"`
	Flags         string `json:"FLAGS"`
	Interface     string `json:"INTERFACE"`
	Ip_address    string `json:"IP_ADDRESS"`
}

var ShowArp_Template string = `Value Required IP_ADDRESS ((?:[0-9]{1,3}\.){3}[0-9]{1,3})
Value HARDWARE_TYPE (\S+)
Value MAC_ADDRESS (((?:[0-9a-fA-F]{2}\:){5}[0-9a-fA-F]{2})|\(incomplete\))
Value FLAGS (\S+)
Value INTERFACE (\S+)

Start
  ^Address\s+HWtype\s+HWaddress\s+Flags\sMask\s+Iface\s*$$ -> ARPTable
  ^\s*$$
  ^. -> Error

ARPTable
  ^${IP_ADDRESS}\s+(${HARDWARE_TYPE})?\s+${MAC_ADDRESS}\s+(${FLAGS})?\s+${INTERFACE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
