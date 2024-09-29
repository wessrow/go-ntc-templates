package cisco_ios 

type ShowProcessesCpu struct {
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

var ShowProcessesCpu_Template = `Value Filldown CPU_USAGE_5_SEC (\d+)
Value Filldown CPU_INTERRUPTION_5_SEC (\d+)
Value Filldown CPU_USAGE_1_MIN (\d+)
Value Filldown CPU_USAGE_5_MIN (\d+)
Value Required PROCESS_PID (\d+)
Value PROCESS_RUNTIME (\d+)
Value PROCESS_INVOKED (\d+)
Value PROCESS_CPU_TIME_FOR_INVOCATION (\d+)
Value PROCESS_CPU_USAGE_5_SEC (\d+\.\d+)
Value PROCESS_CPU_USAGE_1_MIN (\d+\.\d+)
Value PROCESS_CPU_USAGE_5_MIN (\d+\.\d+)
Value PROCESS_TTY (\d+)
Value PROCESS_NAME (.*?)

Start
  ^\s*CPU\s+utilization\s+for\s+five\s+seconds:\s+${CPU_USAGE_5_SEC}%/${CPU_INTERRUPTION_5_SEC}%;\s+one\s+minute:\s+${CPU_USAGE_1_MIN}%;\s+five\s+minutes:\s+${CPU_USAGE_5_MIN}%\s*$$
  ^\s*PID\s+Runtime\(ms\)\s+Invoked\s+uSecs\s+5Sec\s+1Min\s+5Min\s+TTY\s+Process\s*$$
  ^\s*${PROCESS_PID}\s+${PROCESS_RUNTIME}\s+${PROCESS_INVOKED}\s+${PROCESS_CPU_TIME_FOR_INVOCATION}\s+${PROCESS_CPU_USAGE_5_SEC}%\s+${PROCESS_CPU_USAGE_1_MIN}%\s+${PROCESS_CPU_USAGE_5_MIN}%\s+${PROCESS_TTY}\s+${PROCESS_NAME}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`