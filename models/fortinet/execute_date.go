package fortinet

type ExecuteDate struct {
	Year  string `json:"YEAR"`
	Month string `json:"MONTH"`
	Day   string `json:"DAY"`
}

var ExecuteDate_Template string = `Value YEAR (\d{4})
Value MONTH (\d{2})
Value DAY (\d{2})

Start
  ^\s*current\s+date\s+is:\s+${YEAR}-${MONTH}-${DAY}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
