package mikrotik_routeros 

type LogPrintDetailWithoutPaging struct {
	Month	string	`json:"MONTH"`
	Day	string	`json:"DAY"`
	Year	string	`json:"YEAR"`
	Hour	string	`json:"HOUR"`
	Minute	string	`json:"MINUTE"`
	Second	string	`json:"SECOND"`
	Topic	string	`json:"TOPIC"`
	Message	string	`json:"MESSAGE"`
	Datetime	string	`json:"DATETIME"`
	Whole_message	string	`json:"WHOLE_MESSAGE"`
}

var LogPrintDetailWithoutPaging_Template = `Value MONTH (\w{3})
Value DAY (\d{2})
Value YEAR (\d{4})
Value HOUR (\d{2})
Value MINUTE (\d{2})
Value SECOND (\d{2})
Value TOPIC (\S+)
Value MESSAGE (.+)
# Aggregating values
Value DATETIME ((\w{3}/\d{2}(/\d{4})?\s+)?\d{2}:\d{2}:\d{2})
Value WHOLE_MESSAGE (.+?)

Start
  # Check the general structure to skip all the data that are not log messages
  ^\s*time=${DATETIME}\s+${WHOLE_MESSAGE}\s*$$ -> LogMessages

LogMessages
  # Here we do more general parsing to get the log messages' time and text as a whole
  ^\s*time=${DATETIME}\s+${WHOLE_MESSAGE}\s*$$ -> Continue
  # and then extract specific values
  ^\s*time=(${MONTH}/${DAY}(/${YEAR})?\s+)?${HOUR}:${MINUTE}:${SECOND}\s+topics=${TOPIC}\s+message=\"${MESSAGE}\"\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`