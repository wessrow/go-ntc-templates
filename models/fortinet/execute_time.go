package fortinet 

type ExecuteTime struct {
	Hour	string	`json:"HOUR"`
	Minute	string	`json:"MINUTE"`
	Second	string	`json:"SECOND"`
	Last_ntp_sync	string	`json:"LAST_NTP_SYNC"`
}

var ExecuteTime_Template = `Value HOUR (\d{2})
Value MINUTE (\d{2})
Value SECOND (\d{2})
Value LAST_NTP_SYNC (.*?)

Start
  ^\s*current\s+time\s+is:\s+${HOUR}:${MINUTE}:${SECOND}\s*$$
  ^\s*last\s+ntp\s+sync:\s+${LAST_NTP_SYNC}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`