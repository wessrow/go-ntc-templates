package huawei_smartax 

type PortVlan struct {
	Success	string	`json:"SUCCESS"`
	Failed	string	`json:"FAILED"`
}

var PortVlan_Template = `Value SUCCESS (\d+)
Value FAILED (\d+)

Start
  ^Set\sONT\sport\(s\)\sVLAN\sconfiguration,\ssuccess:\s*${SUCCESS},\s*failed:\s*${FAILED} -> Record
  ^\s*$$
  ^. -> Error
`