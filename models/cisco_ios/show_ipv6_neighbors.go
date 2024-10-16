package cisco_ios

type ShowIpv6Neighbors struct {
	Type        string `json:"TYPE"`
	Interface   string `json:"INTERFACE"`
	Address     string `json:"ADDRESS"`
	Age         string `json:"AGE"`
	Mac_address string `json:"MAC_ADDRESS"`
}

var ShowIpv6Neighbors_Template string = `Value ADDRESS (\S+)
Value AGE (\d+|-)
Value MAC_ADDRESS (\S+|-)
Value TYPE (\S+)
Value INTERFACE (\S+)

Start
  ^IPv6\s+Address\s+Age\s+Link-layer\s+Addr\s+State\s+Interface\s*$$
  ^${ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${TYPE}\s+${INTERFACE} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^\s*$$
  ^. -> Error
`
