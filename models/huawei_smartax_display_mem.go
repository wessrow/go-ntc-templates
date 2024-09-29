package models

type HuaweiSmartaxDisplayMem struct {
	Memory_occupancy	string	`json:"MEMORY_OCCUPANCY"`
}

var HuaweiSmartaxDisplayMem_Template = `Value MEMORY_OCCUPANCY (\d+)

Start
  ^\s+Memory\s+occupancy:\s+${MEMORY_OCCUPANCY}%
  ^\s*$$
  ^. -> Error
`