package models

type CiscoAsaDir struct {
	File_system	string	`json:"FILE_SYSTEM"`
	Id	string	`json:"ID"`
	Permissions	string	`json:"PERMISSIONS"`
	Size	string	`json:"SIZE"`
	Total_size	string	`json:"TOTAL_SIZE"`
	Total_free	string	`json:"TOTAL_FREE"`
	Total_percent_free	string	`json:"TOTAL_PERCENT_FREE"`
	Date_time	string	`json:"DATE_TIME"`
	Name	string	`json:"NAME"`
}

var CiscoAsaDir_Template = `Value Filldown FILE_SYSTEM (\S+)
Value ID (\d+)
Value PERMISSIONS (\S+)
Value SIZE (\d+)
Value Fillup TOTAL_SIZE (\d+)
Value Fillup TOTAL_FREE (\d+)
Value Fillup TOTAL_PERCENT_FREE (\d+)
Value DATE_TIME ((<no date>)|(\S+\s\w+\s\d+\s\d+))
Value NAME (\S+)

Start
  ^Directory of\s+${FILE_SYSTEM} -> DIR
  ^. -> Error

DIR
  ^Directory of\s+${FILE_SYSTEM} -> DIR
  ^((\s+)*${ID})\s+${PERMISSIONS}\s+${SIZE}\s+${DATE_TIME}\s+${NAME} -> Record
  ^${TOTAL_SIZE}\s+\S+\s+\S+\s\(${TOTAL_FREE} bytes free\)
  # Accounts for X files total size: X bytes
  ^\d+\s+\S+\s+\S+\s+\d+\s+\S+
  ^${TOTAL_SIZE} bytes total \(${TOTAL_FREE} bytes free/${TOTAL_PERCENT_FREE}\% free\)
  ^.*$$
  ^. -> Error DIR

#
EOF

`