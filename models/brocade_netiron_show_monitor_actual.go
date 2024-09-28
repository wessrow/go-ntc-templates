package models

type BrocadeNetironShowMonitorActual struct {
	Monitoredport	string	`json:"monitoredport"`
	Inputmirror	string	`json:"inputmirror"`
	Outputmirror	string	`json:"outputmirror"`
}