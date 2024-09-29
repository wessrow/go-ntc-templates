package models

type CiscoXrShowProcessesCpu struct {
	Cpu_1_min	string	`json:"CPU_1_min"`
	Cpu_5_min	string	`json:"CPU_5_MIN"`
	Cpu_15_min	string	`json:"CPU_15_MIN"`
}

var CiscoXrShowProcessesCpu_Template = `Value CPU_1_min (\d+)
Value CPU_5_MIN (\d+)
Value CPU_15_MIN (\d+)

Start
  ^CPU utilization for one minute:\s${CPU_1_min}%;\s+five minutes:\s+${CPU_5_MIN}%;\s+fifteen minutes:\s+${CPU_15_MIN}% -> Record


`