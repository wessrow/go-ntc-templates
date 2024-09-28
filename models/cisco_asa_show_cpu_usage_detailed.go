package models

type CiscoAsaShowCpuUsageDetailed struct {
	Core	string	`json:"CORE"`
	Cpu_5_sec	string	`json:"CPU_5_SEC"`
	Cpu_5_sec_dp	string	`json:"CPU_5_SEC_DP"`
	Cpu_5_sec_cp	string	`json:"CPU_5_SEC_CP"`
	Cpu_1_min	string	`json:"CPU_1_MIN"`
	Cpu_1_min_dp	string	`json:"CPU_1_MIN_DP"`
	Cpu_1_min_cp	string	`json:"CPU_1_MIN_CP"`
	Cpu_5_min	string	`json:"CPU_5_MIN"`
	Cpu_5_min_dp	string	`json:"CPU_5_MIN_DP"`
	Cpu_5_min_cp	string	`json:"CPU_5_MIN_CP"`
}