package models

type CiscoNxosShowClock struct {
	Time	string	`json:"TIME"`
	Timezone	string	`json:"TIMEZONE"`
	Dayweek	string	`json:"DAYWEEK"`
	Month	string	`json:"MONTH"`
	Day	string	`json:"DAY"`
	Year	string	`json:"YEAR"`
}