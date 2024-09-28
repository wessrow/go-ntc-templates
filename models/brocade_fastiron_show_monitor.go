package models

type BrocadeFastironShowMonitor struct {
	Monitoredport	string	`json:"monitoredport"`
	Inputmirror	string	`json:"inputmirror"`
	Outputmirror	string	`json:"outputmirror"`
}