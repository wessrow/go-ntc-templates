package arista_eos

type ShowClock struct {
	Month    string `json:"MONTH"`
	Day      string `json:"DAY"`
	Year     string `json:"YEAR"`
	Time     string `json:"TIME"`
	Timezone string `json:"TIMEZONE"`
	Dayweek  string `json:"DAYWEEK"`
}

var ShowClock_Template string = `Value TIME (\d+:\d+:\d+)
Value TIMEZONE (\S+)
Value DAYWEEK (\w+)
Value MONTH (\w+)
Value DAY (\d+)
Value YEAR (\d+)

Start
  ^${DAYWEEK}\s+${MONTH}\s+${DAY}\s+${TIME}\s+${YEAR}
  ^[t|T]imezone(:|\sis)\s+${TIMEZONE} -> Record
`
