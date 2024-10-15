package huawei_smartax

type DisplayCpu struct {
	Cpu_occupancy string `json:"CPU_OCCUPANCY"`
}

var DisplayCpu_Template string = `Value CPU_OCCUPANCY (\d+)

Start
  ^\s+CPU\s+occupancy:\s+${CPU_OCCUPANCY}%
  ^\s*$$
  ^. -> Error`
