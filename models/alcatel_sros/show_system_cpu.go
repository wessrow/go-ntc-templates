package alcatel_sros

type ShowSystemCpu struct {
	Name           string `json:"NAME"`
	Cpu_time       string `json:"CPU_TIME"`
	Cpu_usage      string `json:"CPU_USAGE"`
	Capacity_usage string `json:"CAPACITY_USAGE"`
}

var ShowSystemCpu_Template string = `Value NAME (\S.*?)
Value CPU_TIME (\S+)
Value CPU_USAGE (\S+)
Value CAPACITY_USAGE (\S+)

Start
  ^=+
  ^CPU\sUtil.+
  ^Name\s+CPU\s+Time\s+CPU\s+Usage\s+Capacity\s*$$
  ^\s+\(uSec\)\s+Usage
  ^--.+ -> Resources
  ^\s*$$
  ^. -> Error

Resources
  ^${NAME}\s{4,}${CPU_TIME}\s+${CPU_USAGE}\s+${CAPACITY_USAGE}\s*$$ -> Record
  ^--.+ -> Total
  ^\s*$$
  ^. -> Error

Total
  ^${NAME}\s{4,}${CPU_TIME}\s+${CPU_USAGE}\s+  -> Record
  ^\s+${NAME}\s{4,}${CPU_TIME}\s+${CPU_USAGE}\s+  -> Record
  ^=+ -> Done
  ^. -> Error

Done
`
