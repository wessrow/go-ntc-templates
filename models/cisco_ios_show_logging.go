package models

type CiscoIosShowLogging struct {
	Number	string	`json:"NUMBER"`
	Month	string	`json:"MONTH"`
	Day	string	`json:"DAY"`
	Time	string	`json:"TIME"`
	Timezone	string	`json:"TIMEZONE"`
	Facility	string	`json:"FACILITY"`
	Severity	string	`json:"SEVERITY"`
	Mnemonic	string	`json:"MNEMONIC"`
	Message	[]string	`json:"MESSAGE"`
}