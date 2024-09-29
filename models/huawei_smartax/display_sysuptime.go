package huawei_smartax 

type DisplaySysuptime struct {
	Day	string	`json:"DAY"`
	Hour	string	`json:"HOUR"`
	Minute	string	`json:"MINUTE"`
	Second	string	`json:"SECOND"`
}

var DisplaySysuptime_Template = `Value DAY (\d+)
Value HOUR (\d+)
Value MINUTE (\d+)
Value SECOND (\d+)

Start
  ^\s+System\s+up\s+time:\s+${DAY}\s+day\s+${HOUR}\s+hour\s+${MINUTE}\s+minute\s+${SECOND}\s+
  ^\s*$$ -> EOF
  ^. -> Error`