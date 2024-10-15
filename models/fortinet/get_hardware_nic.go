package fortinet

type GetHardwareNic struct {
	Nic string `json:"NIC"`
}

var GetHardwareNic_Template string = `Value NIC (\S+)

Start
  ^The\s+following\s+NICs\s+are\s+available:\s*$$
  ^\s*${NIC}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
