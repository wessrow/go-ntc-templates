package cisco_nxos

type Dir struct {
	Total_size  string `json:"TOTAL_SIZE"`
	Total_free  string `json:"TOTAL_FREE"`
	Total_used  string `json:"TOTAL_USED"`
	File_system string `json:"FILE_SYSTEM"`
	Size        string `json:"SIZE"`
	Date_time   string `json:"DATE_TIME"`
	Item_name   string `json:"ITEM_NAME"`
}

var Dir_Template string = `Value SIZE (\d+)
Value DATE_TIME (\w+\s+\d+\s+\d+\:\d+\:\d+\s+\d+)
Value ITEM_NAME (\S+)
Value Fillup TOTAL_SIZE (\d+)
Value Fillup TOTAL_FREE (\d+)
Value Fillup TOTAL_USED (\d+)
Value Fillup FILE_SYSTEM (\w+)

Start
  ^Input\/output\s+error
  ^\s*${SIZE}\s+${DATE_TIME}\s+${ITEM_NAME} -> Record
  ^\s*Usage\s+for\s+${FILE_SYSTEM}://
  ^\s*${TOTAL_USED}\s+bytes\s+used
  ^\s*${TOTAL_FREE}\s+bytes\s+free
  ^\s*${TOTAL_SIZE}\s+bytes\s+total
  ^.+ -> Error

EOF
`
