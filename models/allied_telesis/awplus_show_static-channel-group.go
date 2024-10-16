package allied_telesis

type AwplusShowStaticChannelGroup struct {
	Interfaces []string `json:"INTERFACES"`
	Po_name    string   `json:"PO_NAME"`
}

var AwplusShowStaticChannelGroup_Template string = `Value PO_NAME (\S+)
Value List INTERFACES (\w+\.\d\.\w+)

Start
  ^LAG\s+Maximum
  ^LAG\s+Static\s+Maximum
  ^LAG\s+Dynamic\s+Maximum
  ^LAG\s+Static\s+Count
  ^LAG\s+Dynamic\s+Count
  ^LAG\s+Total\s+Count
  ^% -> AlternativeOutput
  ^Static\s+Aggregator -> Continue.Record
  ^Static\s+Aggregator:\s+${PO_NAME}
  ^Member
  ^\s+${INTERFACES}
  ^. -> Error

AlternativeOutput
  ^%\s+LAG\s+Maximum
  ^%\s+LAG\s+Static\s+Maximum
  ^%\s+LAG\s+Dynamic\s+Maximum
  ^%\s+LAG\s+Static\s+Count
  ^%\s+LAG\s+Dynamic\s+Count
  ^%\s+LAG\s+Total\s+Count
  ^%\s+Static\s+Aggregator -> Continue.Record
  ^%\s+Static\s+Aggregator:\s+${PO_NAME}
  ^%\s+Member:\s+${PO_NAME}
  ^%\s+Member:$$
  ^\s+${INTERFACES}
  ^. -> Error
`
