package cisco_ios

type ShowProcessesMemorySorted struct {
	Memory_free       string   `json:"MEMORY_FREE"`
	Process_id        []string `json:"PROCESS_ID"`
	Process_allocated []string `json:"PROCESS_ALLOCATED"`
	Process_freed     []string `json:"PROCESS_FREED"`
	Process_holding   []string `json:"PROCESS_HOLDING"`
	Process           []string `json:"PROCESS"`
	Memory_total      string   `json:"MEMORY_TOTAL"`
	Memory_used       string   `json:"MEMORY_USED"`
}

var ShowProcessesMemorySorted_Template string = `Value MEMORY_TOTAL (\d+)
Value MEMORY_USED (\d+)
Value MEMORY_FREE (\d+)
Value List PROCESS_ID (\d+)
Value List PROCESS_ALLOCATED (\d+)
Value List PROCESS_FREED (\d+)
Value List PROCESS_HOLDING (\d+)
Value List PROCESS (.+?)

Start
  ^Processor\s+Pool\s+Total:\s+${MEMORY_TOTAL}\s+Used:\s+${MEMORY_USED}\s+Free:\s+${MEMORY_FREE}
  ^\s*lsmpi_io\s+Pool
  ^\s*PID\s+TTY\s+Allocated\s+Freed\s+Holding\s+Getbufs\s+Retbufs\s+Process\s*$$ -> Process
  ^\s*$$
  ^. -> Error

Process
  ^\s*${PROCESS_ID}\s+\d+\s+${PROCESS_ALLOCATED}\s+${PROCESS_FREED}\s+${PROCESS_HOLDING}\s+\d+\s+\d+\s+${PROCESS}\s*$$
  ^\s*\d+\s+Total\s*$$
  ^\s*$$
  ^. -> Error
`
