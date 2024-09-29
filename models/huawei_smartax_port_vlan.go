package models

type HuaweiSmartaxPortVlan struct {
	Success	string	`json:"SUCCESS"`
	Failed	string	`json:"FAILED"`
}

var HuaweiSmartaxPortVlan_Template = `Value SUCCESS (\d+)
Value FAILED (\d+)

Start
  ^Set\sONT\sport\(s\)\sVLAN\sconfiguration,\ssuccess:\s*${SUCCESS},\s*failed:\s*${FAILED} -> Record
  ^\s*$$
  ^. -> Error

`