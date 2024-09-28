package models

type AristaEosShowProcessesTopOnce struct {
	Global_timestamp	string	`json:"GLOBAL_TIMESTAMP"`
	Global_users	string	`json:"GLOBAL_USERS"`
	Global_load_average_1_minutes	string	`json:"GLOBAL_LOAD_AVERAGE_1_MINUTES"`
	Global_load_average_5_minutes	string	`json:"GLOBAL_LOAD_AVERAGE_5_MINUTES"`
	Global_load_average_15_minutes	string	`json:"GLOBAL_LOAD_AVERAGE_15_MINUTES"`
	Global_tasks_total	string	`json:"GLOBAL_TASKS_TOTAL"`
	Global_tasks_running	string	`json:"GLOBAL_TASKS_RUNNING"`
	Global_tasks_sleeping	string	`json:"GLOBAL_TASKS_SLEEPING"`
	Global_tasks_stopped	string	`json:"GLOBAL_TASKS_STOPPED"`
	Global_tasks_zombie	string	`json:"GLOBAL_TASKS_ZOMBIE"`
	Global_cpu_percent_user	string	`json:"GLOBAL_CPU_PERCENT_USER"`
	Global_cpu_percent_system	string	`json:"GLOBAL_CPU_PERCENT_SYSTEM"`
	Global_cpu_percent_nice	string	`json:"GLOBAL_CPU_PERCENT_NICE"`
	Global_cpu_percent_idle	string	`json:"GLOBAL_CPU_PERCENT_IDLE"`
	Global_cpu_percent_iowait	string	`json:"GLOBAL_CPU_PERCENT_IOWAIT"`
	Global_cpu_percent_hi	string	`json:"GLOBAL_CPU_PERCENT_HI"`
	Global_cpu_percent_si	string	`json:"GLOBAL_CPU_PERCENT_SI"`
	Global_cpu_percent_stolen	string	`json:"GLOBAL_CPU_PERCENT_STOLEN"`
	Global_mem_unit	string	`json:"GLOBAL_MEM_UNIT"`
	Global_mem_total	string	`json:"GLOBAL_MEM_TOTAL"`
	Global_mem_used	string	`json:"GLOBAL_MEM_USED"`
	Global_mem_free	string	`json:"GLOBAL_MEM_FREE"`
	Global_mem_buffers	string	`json:"GLOBAL_MEM_BUFFERS"`
	Global_swap_mem_unit	string	`json:"GLOBAL_SWAP_MEM_UNIT"`
	Global_swap_mem_total	string	`json:"GLOBAL_SWAP_MEM_TOTAL"`
	Global_swap_mem_used	string	`json:"GLOBAL_SWAP_MEM_USED"`
	Global_swap_mem_free	string	`json:"GLOBAL_SWAP_MEM_FREE"`
	Global_swap_mem_cached	string	`json:"GLOBAL_SWAP_MEM_CACHED"`
	Pid	string	`json:"PID"`
	User	string	`json:"USER"`
	Priority	string	`json:"PRIORITY"`
	Nice	string	`json:"NICE"`
	Virtual_memory_size	string	`json:"VIRTUAL_MEMORY_SIZE"`
	Resident_memory_size	string	`json:"RESIDENT_MEMORY_SIZE"`
	Shared_memory_size	string	`json:"SHARED_MEMORY_SIZE"`
	Process_status	string	`json:"PROCESS_STATUS"`
	Percent_cpu	string	`json:"PERCENT_CPU"`
	Percent_memory	string	`json:"PERCENT_MEMORY"`
	Cpu_time	string	`json:"CPU_TIME"`
	Command	string	`json:"COMMAND"`
}