package avaya_ers 

type ShowMacAddressTable struct {
	Mac_address	string	`json:"MAC_ADDRESS"`
	Vid	string	`json:"VID"`
	Type	string	`json:"TYPE"`
	Port	string	`json:"PORT"`
}

var ShowMacAddressTable_Template = `Value MAC_ADDRESS (\S+)
Value VID (\d+)
Value TYPE (\S+)
Value PORT (\d+)

Start
  ^\s+MAC Address.*
  ^-+ -+ -+ -+
  ^${MAC_ADDRESS}\s+${VID}\s+${TYPE}\s+\w*:*\s*${PORT}* -> Record

Done`