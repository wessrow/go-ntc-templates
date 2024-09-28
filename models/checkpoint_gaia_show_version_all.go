package models

type CheckpointGaiaShowVersionAll struct {
	Version	string	`json:"VERSION"`
	Build	string	`json:"BUILD"`
	Kernel	string	`json:"KERNEL"`
	Architecture	string	`json:"ARCHITECTURE"`
}