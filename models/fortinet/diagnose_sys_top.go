package fortinet 

type DiagnoseSysTop struct {
	Uptime_days	string	`json:"UPTIME_DAYS"`
	Uptime_hours	string	`json:"UPTIME_HOURS"`
	Uptime_minutes	string	`json:"UPTIME_MINUTES"`
	Userspace_app_cpu	string	`json:"USERSPACE_APP_CPU"`
	N_cpu	string	`json:"N_CPU"`
	System_app_cpu	string	`json:"SYSTEM_APP_CPU"`
	Idle_cpu	string	`json:"IDLE_CPU"`
	Wa_cpu	string	`json:"WA_CPU"`
	Hi_cpu	string	`json:"HI_CPU"`
	Si_cpu	string	`json:"SI_CPU"`
	St_cpu	string	`json:"ST_CPU"`
	Total_memory	string	`json:"TOTAL_MEMORY"`
	Free_memory	string	`json:"FREE_MEMORY"`
	Total_shared_memory_pages	string	`json:"TOTAL_SHARED_MEMORY_PAGES"`
	Process_name	string	`json:"PROCESS_NAME"`
	Process_id	string	`json:"PROCESS_ID"`
	Process_state	string	`json:"PROCESS_STATE"`
	Process_cpu_usage	string	`json:"PROCESS_CPU_USAGE"`
	Process_memory_usage	string	`json:"PROCESS_MEMORY_USAGE"`
	Process_cpu_core	string	`json:"PROCESS_CPU_CORE"`
}

var DiagnoseSysTop_Template = `Value UPTIME_DAYS (\d+)
Value UPTIME_HOURS (\d+)
Value UPTIME_MINUTES (\d+)
Value USERSPACE_APP_CPU (\d+)
Value N_CPU (\d+)
Value SYSTEM_APP_CPU (\d+)
Value IDLE_CPU (\d+)
Value WA_CPU (\d+)
Value HI_CPU (\d+)
Value SI_CPU (\d+)
Value ST_CPU (\d+)
Value TOTAL_MEMORY (\d+)
Value FREE_MEMORY (\d+)
Value TOTAL_SHARED_MEMORY_PAGES (\d+)
Value PROCESS_NAME (\S+)
Value PROCESS_ID (\d+)
Value PROCESS_STATE ((R|S|Z|D)(\s+(<|N))?)
Value PROCESS_CPU_USAGE (\d+\.\d+)
Value PROCESS_MEMORY_USAGE (\d+\.\d+)
Value PROCESS_CPU_CORE (\d+)

Start
  ^\s*Run\s+Time:\s+${UPTIME_DAYS}\s+days,\s+${UPTIME_HOURS}\s+hours\s+and\s+${UPTIME_MINUTES}\s+minutes\s*$$
  ^\s*${USERSPACE_APP_CPU}U(,\s+${N_CPU}N)?,\s+${SYSTEM_APP_CPU}S,\s+${IDLE_CPU}I(,\s+${WA_CPU}WA)?(,\s+${HI_CPU}HI)?(,\s+${SI_CPU}SI)?(,\s+${ST_CPU}ST)?;\s+${TOTAL_MEMORY}T,\s+${FREE_MEMORY}F(,\s+${TOTAL_SHARED_MEMORY_PAGES}KF)?\s*$$
  ^\s*${PROCESS_NAME}\s+${PROCESS_ID}\s+${PROCESS_STATE}\s+${PROCESS_CPU_USAGE}\s+${PROCESS_MEMORY_USAGE}(\s+${PROCESS_CPU_CORE})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`