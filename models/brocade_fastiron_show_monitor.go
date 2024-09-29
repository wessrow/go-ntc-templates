package models

type BrocadeFastironShowMonitor struct {
	Monitoredport	string	`json:"monitoredport"`
	Inputmirror	string	`json:"inputmirror"`
	Outputmirror	string	`json:"outputmirror"`
}

var BrocadeFastironShowMonitor_Template = `Value monitoredport ([0-9\/]+)
Value inputmirror ([0-9\/\(\)\ UM]+)
Value outputmirror ([0-9\/\(\)\ UM]+)

Start
  ^\w -> Continue.Record
  ^Monitored Port ${monitoredport} -> Continue
  ^\s+Input mirrored by\s+:\s+${inputmirror} -> Continue
  ^\s+Output mirrored by\s+:\s+${outputmirror} -> Continue

`