package cisco_xr

type Dir struct {
	File_system string `json:"FILE_SYSTEM"`
	Permissions string `json:"PERMISSIONS"`
	Size        string `json:"SIZE"`
	Total_size  string `json:"TOTAL_SIZE"`
	Date_day    string `json:"DATE_DAY"`
	Date_time   string `json:"DATE_TIME"`
	Total_free  string `json:"TOTAL_FREE"`
	Date_month  string `json:"DATE_MONTH"`
	Date_year   string `json:"DATE_YEAR"`
	Name        string `json:"NAME"`
}

var Dir_Template string = `Value Filldown FILE_SYSTEM (\S+)
Value PERMISSIONS (.+?)
Value SIZE (\d+)
Value Fillup TOTAL_SIZE (\d+)
Value Fillup TOTAL_FREE (\d+)
Value DATE_MONTH (\w+)
Value DATE_DAY (\d+)
Value DATE_TIME (\d+:\d+)
Value DATE_YEAR (\d+)
Value NAME (\S+)

Start
  ^\s*$$
  ^\S{3}\s+\S{3}\s+\d{1,2}\s+\d+:\d+:\d+
  ^Directory\s+of\s+${FILE_SYSTEM}
  ^${PERMISSIONS}\.\s+\d+\s+${SIZE}\s+${DATE_MONTH}\s+${DATE_DAY}\s+(${DATE_TIME}|${DATE_YEAR})\s+${NAME} -> Record
  ^${TOTAL_SIZE}\s+kbytes\s+total\s+\(${TOTAL_FREE}\s+kbytes\s+free\)
  ^.* -> Error

EOF
`
