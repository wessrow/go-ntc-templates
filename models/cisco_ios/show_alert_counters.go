package cisco_ios

type ShowAlertCounters struct {
	Description    string `json:"description"`
	Recommendation string `json:"recommendation"`
	Interface      string `json:"interface"`
	Errorcode      string `json:"errorcode"`
	Timestamp      string `json:"timestamp"`
}

var ShowAlertCounters_Template string = `Value Filldown interface (.+) 
Value errorcode (.+)
Value timestamp (.+)
Value description (.+)
Value recommendation (.+)

Start
  ^Interface:\s+${interface} -> Clear
  ^Error Code:\s+${errorcode}
  ^Timestamp:\s+${timestamp}
  ^Description:\s+${description}
  ^Recommendation:\s+${recommendation} -> Record
  ^\s*Global\s+Errors
  ^\s*$$
  ^. -> Error

EOF
`
