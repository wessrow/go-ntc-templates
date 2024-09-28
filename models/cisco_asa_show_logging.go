package models

type CiscoAsaShowLogging struct {
	Number	string	`json:"NUMBER"`
	Month	string	`json:"MONTH"`
	Day	string	`json:"DAY"`
	Year	string	`json:"YEAR"`
	Time	string	`json:"TIME"`
	Hostname	string	`json:"HOSTNAME"`
	Timezone	string	`json:"TIMEZONE"`
	Facility	string	`json:"FACILITY"`
	Severity	string	`json:"SEVERITY"`
	Mnemonic	string	`json:"MNEMONIC"`
	Message	[]string	`json:"MESSAGE"`
}