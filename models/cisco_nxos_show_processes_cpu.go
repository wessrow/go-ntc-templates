package models

type CiscoNxosShowProcessesCpu struct {
	Cpu_5_sec	string	`json:"CPU_5_SEC"`
	Cpu_1_min	string	`json:"CPU_1_MIN"`
	Cpu_5_min	string	`json:"CPU_5_MIN"`
	Interrupts	string	`json:"INTERRUPTS"`
}