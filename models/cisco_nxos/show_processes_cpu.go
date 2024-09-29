package cisco_nxos 

type ShowProcessesCpu struct {
	Cpu_5_sec	string	`json:"CPU_5_SEC"`
	Cpu_1_min	string	`json:"CPU_1_MIN"`
	Cpu_5_min	string	`json:"CPU_5_MIN"`
	Interrupts	string	`json:"INTERRUPTS"`
}

var ShowProcessesCpu_Template = `Value CPU_5_SEC (\d+)
Value CPU_1_MIN (\d+)
Value CPU_5_MIN (\d+)
Value INTERRUPTS (\d+)

Start
  ^CPU utilization for five seconds:\W+${CPU_5_SEC}%/${INTERRUPTS}%; one minute:\W+${CPU_1_MIN}%; five minutes:\W+${CPU_5_MIN}% -> Record
`