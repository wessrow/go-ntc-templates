package models

type CiscoIosShowAlertCounters struct {
	Interface	string	`json:"interface"`
	Errorcode	string	`json:"errorcode"`
	Timestamp	string	`json:"timestamp"`
	Description	string	`json:"description"`
	Recommendation	string	`json:"recommendation"`
}