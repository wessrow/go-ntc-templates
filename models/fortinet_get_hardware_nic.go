package models

type FortinetGetHardwareNic struct {
	Nic	string	`json:"NIC"`
}

var FortinetGetHardwareNic_Template = `Value NIC (\S+)

Start
  ^The\s+following\s+NICs\s+are\s+available:\s*$$
  ^\s*${NIC}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`