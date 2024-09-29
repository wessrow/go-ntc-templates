package models

type AlliedTelesisAwplusShowEtherchannelSummary struct {
	Po_name	string	`json:"PO_NAME"`
	Interfaces	[]string	`json:"INTERFACES"`
}

var AlliedTelesisAwplusShowEtherchannelSummary_Template = `Value PO_NAME (\S+)
Value List INTERFACES (\w+\.\d\.\w+)

Start
  ^Aggregator -> Continue.Record
  ^Aggregator\s+${PO_NAME}
  ^\s+Admin\s+Key
  ^\s+Link:\s+${INTERFACES}
  ^. -> Error

`