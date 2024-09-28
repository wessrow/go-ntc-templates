package models

type CiscoIosShowLicense struct {
	Feature	string	`json:"FEATURE"`
	Period_left	string	`json:"PERIOD_LEFT"`
	Period_used	string	`json:"PERIOD_USED"`
	License_type	string	`json:"LICENSE_TYPE"`
	License_state	string	`json:"LICENSE_STATE"`
	License_count	string	`json:"LICENSE_COUNT"`
	License_priority	string	`json:"LICENSE_PRIORITY"`
}