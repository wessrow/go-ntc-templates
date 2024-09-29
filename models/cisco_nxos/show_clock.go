package cisco_nxos 

type ShowClock struct {
	Time	string	`json:"TIME"`
	Timezone	string	`json:"TIMEZONE"`
	Dayweek	string	`json:"DAYWEEK"`
	Month	string	`json:"MONTH"`
	Day	string	`json:"DAY"`
	Year	string	`json:"YEAR"`
}

var ShowClock_Template = `Value TIME (\d+:\d+:\d+\.\d+)
Value TIMEZONE (\w+)
Value DAYWEEK (\w+)
Value MONTH (\w+)
Value DAY (\d+)
Value YEAR (\d+)

Start
  ^${TIME}\s${TIMEZONE}\s${DAYWEEK}\s${MONTH}\s${DAY}\s${YEAR} -> Record 
`