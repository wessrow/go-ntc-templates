package models

type BrocadeNetironShowMonitorActual struct {
	Monitoredport	string	`json:"monitoredport"`
	Inputmirror	string	`json:"inputmirror"`
	Outputmirror	string	`json:"outputmirror"`
}

var BrocadeNetironShowMonitorActual_Template = `Value monitoredport ([0-9\/]+)
Value inputmirror ([0-9\/\(\)\ UM]+)
Value outputmirror ([0-9\/\(\)\ UM]+)

Start
  ^\w -> Continue.Record
  ^Monitored Port ${monitoredport} -> Continue
  ^\s+Input traffic mirrored to:\s+${inputmirror} -> Continue
  ^\s+Output traffic mirrored to:\s+${outputmirror} -> Continue

`