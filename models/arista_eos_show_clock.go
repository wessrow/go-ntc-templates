package models

type AristaEosShowClock struct {
	Time	string	`json:"TIME"`
	Timezone	string	`json:"TIMEZONE"`
	Dayweek	string	`json:"DAYWEEK"`
	Month	string	`json:"MONTH"`
	Day	string	`json:"DAY"`
	Year	string	`json:"YEAR"`
}

var AristaEosShowClock_Template = `Value TIME (\d+:\d+:\d+)
Value TIMEZONE (\S+)
Value DAYWEEK (\w+)
Value MONTH (\w+)
Value DAY (\d+)
Value YEAR (\d+)

Start
  ^${DAYWEEK}\s+${MONTH}\s+${DAY}\s+${TIME}\s+${YEAR}
  ^[t|T]imezone(:|\sis)\s+${TIMEZONE} -> Record

`