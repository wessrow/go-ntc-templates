package huawei_smartax 

type DisplayMem struct {
	Memory_occupancy	string	`json:"MEMORY_OCCUPANCY"`
}

var DisplayMem_Template = `Value MEMORY_OCCUPANCY (\d+)

Start
  ^\s+Memory\s+occupancy:\s+${MEMORY_OCCUPANCY}%
  ^\s*$$
  ^. -> Error`