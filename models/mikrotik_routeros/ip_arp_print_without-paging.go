package mikrotik_routeros 

type IpArpPrintWithoutPaging struct {
	Index	string	`json:"INDEX"`
	Flags	string	`json:"FLAGS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
}

var IpArpPrintWithoutPaging_Template = `Value Key INDEX (\d+)
Value FLAGS ([XIHDPC]+)
Value IP_ADDRESS (\S+)
Value MAC_ADDRESS (([0-9a-fA-F]{2}[:]){5}([0-9a-fA-F]{2}))
Value INTERFACE (\S+)

Start
  ^\s*#\s*ADDRESS\s*MAC-ADDRESS\s*INTERFACE\s*$$ -> EntriesTable

EntriesTable
  ^\s*${INDEX}\s*(${FLAGS})?\s*${IP_ADDRESS}\s*(${MAC_ADDRESS})?\s+(${INTERFACE})?\s*$$ -> Record
  ^\s*(?:\d{2}:){2}\d{2}\s+echo:\s*.*$$ -> Next
  ^. -> Error
`