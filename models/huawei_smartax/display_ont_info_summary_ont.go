package huawei_smartax 

type DisplayOntInfoSummaryOnt struct {
	Ont_id	string	`json:"ONT_ID"`
	Run_state	string	`json:"RUN_STATE"`
	Last_uptime	string	`json:"LAST_UPTIME"`
	Last_downtime	string	`json:"LAST_DOWNTIME"`
	Last_downcause	string	`json:"LAST_DOWNCAUSE"`
}

var DisplayOntInfoSummaryOnt_Template = `Value Key ONT_ID (\d+)
Value RUN_STATE (\w+)
Value LAST_UPTIME (\S+\s*\S+)
Value LAST_DOWNTIME (\S+\s*\S+)
Value LAST_DOWNCAUSE (\S+)

Start
  ^\s+-
  ^\s+ONT\s+Run\s+Last\s+Last\s+Last
  ^\s+ID\s+State\s+UpTime\s+DownTime\s+DownCause -> ONTs

ONTs
  ^\s+${ONT_ID}\s+${RUN_STATE}\s+(-|${LAST_UPTIME})\s*(-|${LAST_DOWNTIME})\s+(-|${LAST_DOWNCAUSE}) -> Record
  ^\s+-
  ^\s+ONT\s+SN\s+Type\s+Distance\s+Rx\/Tx\s+power\s+Description
  ^\s+ID\s+\(m\)\s+\(dBm\) -> EOF
  ^. -> Error
`