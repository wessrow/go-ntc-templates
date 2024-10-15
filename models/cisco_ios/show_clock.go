package cisco_ios

type ShowClock struct {
	Day      string `json:"DAY"`
	Year     string `json:"YEAR"`
	Time     string `json:"TIME"`
	Timezone string `json:"TIMEZONE"`
	Dayweek  string `json:"DAYWEEK"`
	Month    string `json:"MONTH"`
}

var ShowClock_Template string = `Value TIME (\d+:\d+:\d+\.\d+)
Value TIMEZONE (\w+)
Value DAYWEEK (\w+)
Value MONTH (\w+)
Value DAY (\d+)
Value YEAR (\d+)

Start
  ^\s*(\*)?\.?${TIME}\s${TIMEZONE}\s${DAYWEEK}\s${MONTH}\s${DAY}\s${YEAR} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`
