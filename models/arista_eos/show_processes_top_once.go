package arista_eos

type ShowProcessesTopOnce struct {
	Cpu_time                       string `json:"CPU_TIME"`
	Global_users                   string `json:"GLOBAL_USERS"`
	Global_cpu_percent_idle        string `json:"GLOBAL_CPU_PERCENT_IDLE"`
	Global_mem_used                string `json:"GLOBAL_MEM_USED"`
	Nice                           string `json:"NICE"`
	Global_load_average_5_minutes  string `json:"GLOBAL_LOAD_AVERAGE_5_MINUTES"`
	Global_mem_total               string `json:"GLOBAL_MEM_TOTAL"`
	Percent_cpu                    string `json:"PERCENT_CPU"`
	Global_cpu_percent_iowait      string `json:"GLOBAL_CPU_PERCENT_IOWAIT"`
	Global_load_average_1_minutes  string `json:"GLOBAL_LOAD_AVERAGE_1_MINUTES"`
	Global_tasks_running           string `json:"GLOBAL_TASKS_RUNNING"`
	Global_tasks_sleeping          string `json:"GLOBAL_TASKS_SLEEPING"`
	Global_cpu_percent_user        string `json:"GLOBAL_CPU_PERCENT_USER"`
	Global_swap_mem_used           string `json:"GLOBAL_SWAP_MEM_USED"`
	Resident_memory_size           string `json:"RESIDENT_MEMORY_SIZE"`
	Shared_memory_size             string `json:"SHARED_MEMORY_SIZE"`
	Process_status                 string `json:"PROCESS_STATUS"`
	Pid                            string `json:"PID"`
	Command                        string `json:"COMMAND"`
	Global_tasks_zombie            string `json:"GLOBAL_TASKS_ZOMBIE"`
	Global_cpu_percent_system      string `json:"GLOBAL_CPU_PERCENT_SYSTEM"`
	Global_mem_buffers             string `json:"GLOBAL_MEM_BUFFERS"`
	Global_swap_mem_free           string `json:"GLOBAL_SWAP_MEM_FREE"`
	Global_swap_mem_unit           string `json:"GLOBAL_SWAP_MEM_UNIT"`
	User                           string `json:"USER"`
	Virtual_memory_size            string `json:"VIRTUAL_MEMORY_SIZE"`
	Global_load_average_15_minutes string `json:"GLOBAL_LOAD_AVERAGE_15_MINUTES"`
	Global_tasks_total             string `json:"GLOBAL_TASKS_TOTAL"`
	Global_cpu_percent_si          string `json:"GLOBAL_CPU_PERCENT_SI"`
	Global_cpu_percent_stolen      string `json:"GLOBAL_CPU_PERCENT_STOLEN"`
	Global_mem_free                string `json:"GLOBAL_MEM_FREE"`
	Global_swap_mem_total          string `json:"GLOBAL_SWAP_MEM_TOTAL"`
	Global_swap_mem_cached         string `json:"GLOBAL_SWAP_MEM_CACHED"`
	Percent_memory                 string `json:"PERCENT_MEMORY"`
	Global_timestamp               string `json:"GLOBAL_TIMESTAMP"`
	Global_tasks_stopped           string `json:"GLOBAL_TASKS_STOPPED"`
	Global_cpu_percent_nice        string `json:"GLOBAL_CPU_PERCENT_NICE"`
	Global_mem_unit                string `json:"GLOBAL_MEM_UNIT"`
	Global_cpu_percent_hi          string `json:"GLOBAL_CPU_PERCENT_HI"`
	Priority                       string `json:"PRIORITY"`
}

var ShowProcessesTopOnce_Template string = `# Reference for 'top' output can be found here:
# https://man7.org/linux/man-pages/man1/top.1.html
# 
# These are top level summary values we "Filldown" to each record.
# This is because TextFSM does not support storing these as a separate
# dicitonary within the output
#
# Global values from line 1
Value Filldown GLOBAL_TIMESTAMP (\d+:\d+:\d+)
Value Filldown GLOBAL_USERS (\d+)
Value Filldown GLOBAL_LOAD_AVERAGE_1_MINUTES (\d+[.]?\d*)
Value Filldown GLOBAL_LOAD_AVERAGE_5_MINUTES (\d+[.]?\d*)
Value Filldown GLOBAL_LOAD_AVERAGE_15_MINUTES (\d+[.]?\d*)
# Global values from line 2
Value Filldown GLOBAL_TASKS_TOTAL (\d+)
Value Filldown GLOBAL_TASKS_RUNNING (\d+)
Value Filldown GLOBAL_TASKS_SLEEPING (\d+)
Value Filldown GLOBAL_TASKS_STOPPED (\d+)
Value Filldown GLOBAL_TASKS_ZOMBIE (\d+)
# Global values from line 3
Value Filldown GLOBAL_CPU_PERCENT_USER (\d+[.]\d+)
Value Filldown GLOBAL_CPU_PERCENT_SYSTEM (\d+[.]\d+)
Value Filldown GLOBAL_CPU_PERCENT_NICE (\d+[.]\d+)
Value Filldown GLOBAL_CPU_PERCENT_IDLE (\d+[.]\d+)
Value Filldown GLOBAL_CPU_PERCENT_IOWAIT (\d+[.]\d+)
Value Filldown GLOBAL_CPU_PERCENT_HI (\d+[.]\d+)
Value Filldown GLOBAL_CPU_PERCENT_SI (\d+[.]\d+)
Value Filldown GLOBAL_CPU_PERCENT_STOLEN (\d+[.]\d+)  
# Global values from line 4       
Value Filldown GLOBAL_MEM_UNIT (KiB|MiB|GiB|TiB|k|m|g|t)
Value Filldown GLOBAL_MEM_TOTAL (\d+[.]?\d*)
Value Filldown GLOBAL_MEM_USED (\d+[.]?\d*)
Value Filldown GLOBAL_MEM_FREE (\d+[.]?\d*)
Value Filldown GLOBAL_MEM_BUFFERS (\d+[.]?\d*)
# Global values from line 5
Value Filldown GLOBAL_SWAP_MEM_UNIT (KiB|MiB|GiB|TiB|k|m|g|t)
Value Filldown GLOBAL_SWAP_MEM_TOTAL (\d+[.]?\d*)
Value Filldown GLOBAL_SWAP_MEM_USED (\d+[.]?\d*)
Value Filldown GLOBAL_SWAP_MEM_FREE (\d+[.]?\d*)
Value Filldown GLOBAL_SWAP_MEM_CACHED (\d+[.]?\d*)
# Per process records start here
Value Required PID (\d+)
Value USER (\S+)
Value PRIORITY (\d+|rt|RT)
Value NICE (-?\d+)
Value VIRTUAL_MEMORY_SIZE (\d+[.]?\d*m?)
Value RESIDENT_MEMORY_SIZE (\d+[.]?\d*m?)
Value SHARED_MEMORY_SIZE (\d+[.]?\d*m?)
Value PROCESS_STATUS (D|I|R|S|T|t|Z)
Value PERCENT_CPU (\d+[.]\d+)
Value PERCENT_MEMORY (\d+[.]\d+)
Value CPU_TIME (\d+:\d+\S+)
Value COMMAND (\S+)


Start  
  # The top two lines are the same for all tested versions
  ^top - ${GLOBAL_TIMESTAMP} up .*,\s+${GLOBAL_USERS} use.*,\s+load average:\s+${GLOBAL_LOAD_AVERAGE_1_MINUTES},\s+${GLOBAL_LOAD_AVERAGE_5_MINUTES},\s+${GLOBAL_LOAD_AVERAGE_15_MINUTES}$$
  ^Tasks:\s+${GLOBAL_TASKS_TOTAL} total,\s+${GLOBAL_TASKS_RUNNING} running,\s+${GLOBAL_TASKS_SLEEPING} sleeping,\s+${GLOBAL_TASKS_STOPPED} stopped,\s+${GLOBAL_TASKS_ZOMBIE} zombie$$
  # CPU - 1st format
  # %Cpu(s):  2.8 us,  0.7 sy,  0.0 ni, 96.3 id,  0.0 wa,  0.2 hi,  0.0 si,  0.0 st
  ^%Cpu[(]s[)]:\s*${GLOBAL_CPU_PERCENT_USER}\sus,\s*${GLOBAL_CPU_PERCENT_SYSTEM}\ssy,\s*${GLOBAL_CPU_PERCENT_NICE}\sni,\s*${GLOBAL_CPU_PERCENT_IDLE}\sid,\s*${GLOBAL_CPU_PERCENT_IOWAIT}\swa,\s*${GLOBAL_CPU_PERCENT_HI}\shi,\s*${GLOBAL_CPU_PERCENT_SI}\ssi,\s*${GLOBAL_CPU_PERCENT_STOLEN}\sst$$
  # CPU - 2nd format
  # Cpu(s): 12.5%us,  2.3%sy,  0.0%ni, 84.5%id,  0.0%wa,  0.7%hi,  0.1%si,  0.0%st
  ^Cpu[(]s[)]:\s*${GLOBAL_CPU_PERCENT_USER}%us,\s*${GLOBAL_CPU_PERCENT_SYSTEM}%sy,\s*${GLOBAL_CPU_PERCENT_NICE}%ni,\s*${GLOBAL_CPU_PERCENT_IDLE}%id,\s*${GLOBAL_CPU_PERCENT_IOWAIT}%wa,\s*${GLOBAL_CPU_PERCENT_HI}%hi,\s*${GLOBAL_CPU_PERCENT_SI}%si,\s*${GLOBAL_CPU_PERCENT_STOLEN}%st$$
  # Memory - 1st format
  # KiB Mem:   2014520 total,  1970928 used,    43592 free,   171340 buffers
  ^${GLOBAL_MEM_UNIT}\s+Mem.*:\s+${GLOBAL_MEM_TOTAL}\stotal,\s+${GLOBAL_MEM_USED}\sused,\s+${GLOBAL_MEM_FREE}\sfree,\s+${GLOBAL_MEM_BUFFERS}\sbuffers$$
  # Memory - 2nd format
  # MiB Mem :  7956.2 total,  1755.3 free,  2052.9 used,  4148.0 buff/cache
  ^${GLOBAL_MEM_UNIT}\s+Mem.*:\s+${GLOBAL_MEM_TOTAL}\stotal,\s+${GLOBAL_MEM_FREE}\sfree,\s+${GLOBAL_MEM_USED}\sused,\s+${GLOBAL_MEM_BUFFERS}\sbuff[/]cache$$
  # Memory - 3rd format
  # Mem:   3981336k total,  2726640k used,  1254696k free,   156856k buffers
  ^Mem.*:\s+${GLOBAL_MEM_TOTAL}${GLOBAL_MEM_UNIT}\stotal,\s+${GLOBAL_MEM_USED}[kmgt]\sused,\s+${GLOBAL_MEM_FREE}[kmgt]\sfree,\s+${GLOBAL_MEM_BUFFERS}[kmgt]\sbuffers$$
  # Swap - 1st format
  # KiB Swap:        0 total,        0 used,        0 free,  1156836 cached
  ^${GLOBAL_SWAP_MEM_UNIT}\s+Swap.*:\s+${GLOBAL_SWAP_MEM_TOTAL}\stotal,\s+${GLOBAL_SWAP_MEM_USED}\sused,\s+${GLOBAL_SWAP_MEM_FREE}\sfree,\s+${GLOBAL_SWAP_MEM_CACHED}\scached$$
  # Swap - 2nd format
  # MiB Swap:   0.0 total,   0.0 free,   0.0 used.  5325.2 avail Mem
  ^${GLOBAL_SWAP_MEM_UNIT}\s+Swap.*:\s+${GLOBAL_SWAP_MEM_TOTAL}\stotal,\s+${GLOBAL_SWAP_MEM_FREE}\sfree,\s+${GLOBAL_SWAP_MEM_USED}\sused.\s+${GLOBAL_SWAP_MEM_CACHED}\savail Mem\s*$$
  # Swap - 3rd format
  # Swap:        0k total,        0k used,        0k free,  1521744k cached
  ^Swap.*:\s+${GLOBAL_SWAP_MEM_TOTAL}${GLOBAL_SWAP_MEM_UNIT}\stotal,\s+${GLOBAL_SWAP_MEM_USED}[kmgt]\sused,\s+${GLOBAL_SWAP_MEM_FREE}[kmgt]\sfree,\s+${GLOBAL_SWAP_MEM_CACHED}[kmgt]\scached$$
  ^$$     
  # Process table header
  ^\s*PID\s+USER\s+PR\s+NI\s+VIRT\s+RES\s+SHR\s+S\s+.CPU\s+.MEM\s+TIME.\s+COMMAND\s*$$
  # Process line items in table
  ^\s*${PID}\s+${USER}\s*${PRIORITY}\s*${NICE}\s+${VIRTUAL_MEMORY_SIZE}\s+${RESIDENT_MEMORY_SIZE}\s+${SHARED_MEMORY_SIZE}\s+${PROCESS_STATUS}\s+${PERCENT_CPU}\s+${PERCENT_MEMORY}\s+${CPU_TIME}\s+${COMMAND}\s*$$ -> Record
  ^. -> Error
`
