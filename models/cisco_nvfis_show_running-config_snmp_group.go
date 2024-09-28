package models

type CiscoNvfisShowRunningConfigSnmpGroup struct {
	Groupname	string	`json:"GROUPNAME"`
	Version	string	`json:"VERSION"`
	Priviledge	string	`json:"PRIVILEDGE"`
	Write	string	`json:"WRITE"`
	Read	string	`json:"READ"`
	Notify	string	`json:"NOTIFY"`
}