package models

type CiscoIosShowProcessesCpu struct {
	Cpu_usage_5_sec	string	`json:"CPU_USAGE_5_SEC"`
	Cpu_interruption_5_sec	string	`json:"CPU_INTERRUPTION_5_SEC"`
	Cpu_usage_1_min	string	`json:"CPU_USAGE_1_MIN"`
	Cpu_usage_5_min	string	`json:"CPU_USAGE_5_MIN"`
	Process_pid	string	`json:"PROCESS_PID"`
	Process_runtime	string	`json:"PROCESS_RUNTIME"`
	Process_invoked	string	`json:"PROCESS_INVOKED"`
	Process_cpu_time_for_invocation	string	`json:"PROCESS_CPU_TIME_FOR_INVOCATION"`
	Process_cpu_usage_5_sec	string	`json:"PROCESS_CPU_USAGE_5_SEC"`
	Process_cpu_usage_1_min	string	`json:"PROCESS_CPU_USAGE_1_MIN"`
	Process_cpu_usage_5_min	string	`json:"PROCESS_CPU_USAGE_5_MIN"`
	Process_tty	string	`json:"PROCESS_TTY"`
	Process_name	string	`json:"PROCESS_NAME"`
}