package models

type HuaweiSmartaxDisplayCpu struct {
	Cpu_occupancy	string	`json:"CPU_OCCUPANCY"`
}

var HuaweiSmartaxDisplayCpu_Template = `Value CPU_OCCUPANCY (\d+)

Start
  ^\s+CPU\s+occupancy:\s+${CPU_OCCUPANCY}%
  ^\s*$$
  ^. -> Error
`