package models

type CiscoNxosShowLicenseUsage struct {
	Feature	string	`json:"FEATURE"`
	Installed	string	`json:"INSTALLED"`
	License_count	string	`json:"LICENSE_COUNT"`
	Status	string	`json:"STATUS"`
	Expiry_date	string	`json:"EXPIRY_DATE"`
	Comments	string	`json:"COMMENTS"`
}