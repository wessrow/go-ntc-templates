package fortinet_get 

type HardwareNic struct {
	Nic	string	`json:"NIC"`
}

var HardwareNic_Template = `Value NIC (\S+)

Start
  ^The\s+following\s+NICs\s+are\s+available:\s*$$
  ^\s*${NIC}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`