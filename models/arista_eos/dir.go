package arista_eos

type Dir struct {
	Total_size  string `json:"TOTAL_SIZE"`
	Total_free  string `json:"TOTAL_FREE"`
	File_system string `json:"FILE_SYSTEM"`
	Permissions string `json:"PERMISSIONS"`
	Size        string `json:"SIZE"`
	Date_time   string `json:"DATE_TIME"`
	Name        string `json:"NAME"`
}

var Dir_Template string = `Value Filldown FILE_SYSTEM (\S+)
Value PERMISSIONS (\S+)
Value SIZE (\d+)
Value DATE_TIME (\S+\s+\d+\s+((\d+)|(\d+:\d+))|\<no date\>)
Value NAME (\S+)
Value Fillup TOTAL_SIZE (\d+)
Value Fillup TOTAL_FREE (\d+)

Start
  ^Directory\s+of\s+${FILE_SYSTEM} -> Dispatcher
  ^\s*$$
  ^. -> Error "LINE NOT FOUND"

Dispatcher
  ^\s+${PERMISSIONS}\s+${SIZE}\s+${DATE_TIME}\s+${NAME} -> Record Files
  ^Directory\s+of -> Continue.Record
  ^Directory\s+of\s+${FILE_SYSTEM}
  ^No\s+files -> NoFiles
  ^\s*$$
  ^\. -> Error "LINE NOT FOUND"

Files
  ^\s+${PERMISSIONS}\s+${SIZE}\s+${DATE_TIME}\s+${NAME} -> Record
  ^${TOTAL_SIZE}\s+\S+\s+\S+\s+\(${TOTAL_FREE}\s+\S+\s+\S+\) -> Start
  ^Directory\s+of\s+${FILE_SYSTEM} -> Dispatcher
  ^No\s+space\s+information -> Start
  ^\s*$$
  ^. -> Error "LINE NOT FOUND"

NoFiles
  ^${TOTAL_SIZE}\s+\S+\s+\S+\s+\(${TOTAL_FREE}\s+\S+\s+\S+\) -> Record Start
  ^Directory\s+of -> Continue.Record
  ^Directory\s+of\s+${FILE_SYSTEM} -> Dispatcher
  ^\s*$$
  ^\. -> Error "LINE NOT FOUND"

EOF
`
