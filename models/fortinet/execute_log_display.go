package fortinet

type ExecuteLogDisplay struct {
	Month          string `json:"MONTH"`
	Day            string `json:"DAY"`
	Second         string `json:"SECOND"`
	Message        string `json:"MESSAGE"`
	Logs_found     string `json:"LOGS_FOUND"`
	Logs_returned  string `json:"LOGS_RETURNED"`
	Logs_searched  string `json:"LOGS_SEARCHED"`
	Message_number string `json:"MESSAGE_NUMBER"`
	Year           string `json:"YEAR"`
	Hour           string `json:"HOUR"`
	Minute         string `json:"MINUTE"`
}

var ExecuteLogDisplay_Template string = `Value LOGS_FOUND (\d+)
Value LOGS_RETURNED (\d+)
Value LOGS_SEARCHED (\d+(?:\.\d+)?)
Value MESSAGE_NUMBER (\d+)
Value YEAR (\d{4})
Value MONTH (\d{2})
Value DAY (\d{2})
Value HOUR (\d{2})
Value MINUTE (\d{2})
Value SECOND (\d{2})
Value MESSAGE (.+?)

Start
  ^\s*${LOGS_FOUND}\s+logs\s+found\.\s*$$
  ^\s*${LOGS_RETURNED}\s+logs\s+returned\.\s*$$
  ^\s*${LOGS_SEARCHED}%\s+of\s+logs\s+has\s+been\s+searched\.\s*$$
  ^\s*${MESSAGE_NUMBER}:\s+date=${YEAR}-${MONTH}-${DAY}\s+time=${HOUR}:${MINUTE}:${SECOND}\s+${MESSAGE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
