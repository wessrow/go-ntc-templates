package mikrotik_routeros

type IpArpPrint struct {
	Num         string `json:"NUM"`
	Flags       string `json:"FLAGS"`
	Ip_address  string `json:"IP_ADDRESS"`
	Mac_address string `json:"MAC_ADDRESS"`
	Interface   string `json:"INTERFACE"`
}

var IpArpPrint_Template string = `Value NUM (\d+)
Value FLAGS ([XIHDPC]+)
Value IP_ADDRESS (\S+)
Value MAC_ADDRESS ((?:[0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2})
Value INTERFACE (\S+)

Start
  ^\s*Flags:\s+X\s+-\s+disabled,\s+I\s+-\s+invalid,\s+H\s+-\s+DHCP,\s+D\s+-\s+dynamic,\s+P\s+-\s+published,\s+C\s+-\s+complete\s*$$
  ^\s*#\s+ADDRESS\s+MAC-ADDRESS\s+INTERFACE\s*$$ -> EntriesTable
  ^\s*$$
  ^. -> Error

EntriesTable
  ^\s*${NUM}(?:\s+${FLAGS})?\s+${IP_ADDRESS}(?:\s+${MAC_ADDRESS})?(?:\s+${INTERFACE})?\s*$$ -> Record
  ^\s*(?:\d{2}:){2}\d{2}\s+echo:\s*.*$$
  ^\s*$$
  ^. -> Error
`
