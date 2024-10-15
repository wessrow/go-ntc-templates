package mikrotik_routeros

type SystemResourcePrint struct {
	Cpu_frequency           string `json:"CPU_FREQUENCY"`
	Write_sect_total        string `json:"WRITE_SECT_TOTAL"`
	Architecture_name       string `json:"ARCHITECTURE_NAME"`
	Uptime_minutes          string `json:"UPTIME_MINUTES"`
	Uptime_seconds          string `json:"UPTIME_SECONDS"`
	Version                 string `json:"VERSION"`
	Cpu_load                string `json:"CPU_LOAD"`
	Write_sect_since_reboot string `json:"WRITE_SECT_SINCE_REBOOT"`
	Platform                string `json:"PLATFORM"`
	Free_memory             string `json:"FREE_MEMORY"`
	Total_memory            string `json:"TOTAL_MEMORY"`
	Free_hdd_space          string `json:"FREE_HDD_SPACE"`
	Bad_blocks              string `json:"BAD_BLOCKS"`
	Board_name              string `json:"BOARD_NAME"`
	Uptime                  string `json:"UPTIME"`
	Uptime_weeks            string `json:"UPTIME_WEEKS"`
	Factory_software        string `json:"FACTORY_SOFTWARE"`
	Cpu                     string `json:"CPU"`
	Cpu_count               string `json:"CPU_COUNT"`
	Total_hdd_space         string `json:"TOTAL_HDD_SPACE"`
	Uptime_days             string `json:"UPTIME_DAYS"`
	Uptime_hours            string `json:"UPTIME_HOURS"`
	Build_time              string `json:"BUILD_TIME"`
}

var SystemResourcePrint_Template string = `Value UPTIME (\w+)
Value UPTIME_WEEKS (\d+)
Value UPTIME_DAYS (\d+)
Value UPTIME_HOURS (\d+)
Value UPTIME_MINUTES (\d+)
Value UPTIME_SECONDS (\d+)
Value VERSION (.+)
Value BUILD_TIME ([A-Za-z]{3}\/\d{1,2}\/\d{4} \d{2}:\d{2}:\d{2})
Value FACTORY_SOFTWARE (\S+)
Value FREE_MEMORY (\d+\.\d{1}(K|M|G)iB)
Value TOTAL_MEMORY (\d+\.\d{1}(K|M|G)iB)
Value CPU (.+)
Value CPU_COUNT (\d+)
Value CPU_FREQUENCY (\d+(M|G)Hz)
Value CPU_LOAD (\d{1,3}%)
Value FREE_HDD_SPACE (\d+\.\d{1}(K|M|G)iB)
Value TOTAL_HDD_SPACE (\d+\.\d{1}(K|M|G)iB)
Value WRITE_SECT_SINCE_REBOOT (\d+)
Value WRITE_SECT_TOTAL (\d+)
Value BAD_BLOCKS (\d{1,3}(\.\d+)?%)
Value ARCHITECTURE_NAME (\S*)
Value BOARD_NAME (.+)
Value PLATFORM (.+)

Start
  ^\s*uptime\:\s+${UPTIME}\s*$$ -> Continue
  ^\s*uptime\:\s+(?:${UPTIME_WEEKS}w)?(?:${UPTIME_DAYS}d)?(?:${UPTIME_HOURS}h)?(?:${UPTIME_MINUTES}m)?(?:${UPTIME_SECONDS}s)?\s*$$
  ^\s*version\:\s+${VERSION}\s*$$
  ^\s*build\-time\:\s+${BUILD_TIME}\s*$$
  ^\s*factory\-software\:\s+${FACTORY_SOFTWARE}\s*$$
  ^\s*free\-memory\:\s+${FREE_MEMORY}\s*$$
  ^\s*total\-memory\:\s+${TOTAL_MEMORY}\s*$$
  ^\s*cpu\:\s+${CPU}\s*$$
  ^\s*cpu\-count\:\s+${CPU_COUNT}\s*$$
  ^\s*cpu\-frequency\:\s+${CPU_FREQUENCY}\s*$$
  ^\s*cpu\-load\:\s+${CPU_LOAD}\s*$$
  ^\s*free\-hdd\-space\:\s+${FREE_HDD_SPACE}\s*$$
  ^\s*total\-hdd\-space\:\s+${TOTAL_HDD_SPACE}\s*$$
  ^\s*write\-sect\-since\-reboot\:\s+${WRITE_SECT_SINCE_REBOOT}\s*$$
  ^\s*write\-sect\-total\:\s+${WRITE_SECT_TOTAL}\s*$$
  ^\s*bad\-blocks\:\s+${BAD_BLOCKS}\s*$$
  ^\s*architecture\-name\:\s+${ARCHITECTURE_NAME}\s*$$
  ^\s*board\-name\:\s+${BOARD_NAME}\s*$$
  ^\s*platform\:\s+${PLATFORM}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
