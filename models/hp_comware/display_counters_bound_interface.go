package hp_comware

type DisplayCountersBoundInterface struct {
	Interface         string `json:"INTERFACE"`
	Packets_total     string `json:"PACKETS_TOTAL"`
	Packets_broadcast string `json:"PACKETS_BROADCAST"`
	Packets_multicast string `json:"PACKETS_MULTICAST"`
	Packets_error     string `json:"PACKETS_ERROR"`
}

var DisplayCountersBoundInterface_Template string = `Value INTERFACE (\S+)
Value PACKETS_TOTAL (\d+)
Value PACKETS_BROADCAST (\d+)
Value PACKETS_MULTICAST (\d+)
Value PACKETS_ERROR (\d+)

Start
  ^\s*$$
  ^Interface\s+Total\s+\(pkts\)\s+Broadcast\s+\(pkts\)\s+Multicast\s+\(pkts\)\s+Err\s+\(pkts\) -> InterfaceTable
  ^. -> Error

InterfaceTable
  ^${INTERFACE}\s+${PACKETS_TOTAL}\s+${PACKETS_BROADCAST}\s+${PACKETS_MULTICAST}\s+${PACKETS_ERROR} -> Record
  ^\s*$$
  ^\s*Overflow -> InterfaceTableFooter
  ^. -> Error

InterfaceTableFooter
  # ignore the footer's content
`
