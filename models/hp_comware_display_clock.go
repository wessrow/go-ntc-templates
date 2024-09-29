package models

type HpComwareDisplayClock struct {
	Time	string	`json:"TIME"`
	Timezone	string	`json:"TIMEZONE"`
	Dayweek	string	`json:"DAYWEEK"`
	Month	string	`json:"MONTH"`
	Day	string	`json:"DAY"`
	Year	string	`json:"YEAR"`
}

var HpComwareDisplayClock_Template = `Value TIME (\d+:\d+:\d+)
Value TIMEZONE (\S+)
Value DAYWEEK (\w+)
Value MONTH (\d+)
Value DAY (\d+)
Value YEAR (\d+)

Start
  ^${TIME}\s+${TIMEZONE}\s+${DAYWEEK}\s+${MONTH}/${DAY}/${YEAR} -> Record



`