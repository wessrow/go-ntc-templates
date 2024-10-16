package cisco_ios

type Dir struct {
	Permissions string `json:"PERMISSIONS"`
	Size        string `json:"SIZE"`
	Total_size  string `json:"TOTAL_SIZE"`
	Total_free  string `json:"TOTAL_FREE"`
	Date_time   string `json:"DATE_TIME"`
	Name        string `json:"NAME"`
	File_system string `json:"FILE_SYSTEM"`
	Id          string `json:"ID"`
}

var Dir_Template string = `Value Filldown FILE_SYSTEM (\S+)
Value ID (\d+)
Value PERMISSIONS (.+?)
Value SIZE (\d+)
Value Fillup TOTAL_SIZE (\d+)
Value Fillup TOTAL_FREE (\d+)
Value DATE_TIME (.+?)
Value NAME (\S+)

Start
  ^\s*${ID}\s+${PERMISSIONS}\s+${SIZE}\s+${DATE_TIME}\s+${NAME}\s*$$ -> Record
  ^Directory of\s+${FILE_SYSTEM}
  ^${TOTAL_SIZE}\s+\S+\s+\S+\s\(${TOTAL_FREE} bytes free\)
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^No\s+files\s+in\s+directory -> EmptyFileSystem

EmptyFileSystem
  ^${TOTAL_SIZE}\s+\S+\s+\S+\s\(${TOTAL_FREE} bytes free\) -> Record

EOF
`
