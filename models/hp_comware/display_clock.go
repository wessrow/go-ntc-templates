package hp_comware

type DisplayClock struct {
	Day      string `json:"DAY"`
	Year     string `json:"YEAR"`
	Time     string `json:"TIME"`
	Timezone string `json:"TIMEZONE"`
	Dayweek  string `json:"DAYWEEK"`
	Month    string `json:"MONTH"`
}

var DisplayClock_Template string = `Value TIME (\d+:\d+:\d+)
Value TIMEZONE (\S+)
Value DAYWEEK (\w+)
Value MONTH (\d+)
Value DAY (\d+)
Value YEAR (\d+)

Start
  ^${TIME}\s+${TIMEZONE}\s+${DAYWEEK}\s+${MONTH}/${DAY}/${YEAR} -> Record


`
