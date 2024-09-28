package models

type CiscoIosShowProcessesMemorySorted struct {
	Memory_total	string	`json:"MEMORY_TOTAL"`
	Memory_used	string	`json:"MEMORY_USED"`
	Memory_free	string	`json:"MEMORY_FREE"`
	Process_id	[]string	`json:"PROCESS_ID"`
	Process_allocated	[]string	`json:"PROCESS_ALLOCATED"`
	Process_freed	[]string	`json:"PROCESS_FREED"`
	Process_holding	[]string	`json:"PROCESS_HOLDING"`
	Process	[]string	`json:"PROCESS"`
}