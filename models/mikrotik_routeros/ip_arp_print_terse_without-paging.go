package mikrotik_routeros 

type IpArpPrintTerseWithoutPaging struct {
	Index	string	`json:"INDEX"`
	Flags	string	`json:"FLAGS"`
	Address	string	`json:"ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Published	string	`json:"PUBLISHED"`
}

var IpArpPrintTerseWithoutPaging_Template = `Value INDEX (\d+)
Value FLAGS ([XIHDPC]+)
Value ADDRESS (\S+)
Value MAC_ADDRESS (\S+)
Value INTERFACE (\S+)
Value PUBLISHED (yes|no)

Start
  ^\s*${INDEX}(?:\s+${FLAGS})?\s+address=${ADDRESS}\s+mac-address=${MAC_ADDRESS}\s+interface=${INTERFACE}\s+published=${PUBLISHED}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`