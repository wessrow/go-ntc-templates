package arista_eos

type BashDfH struct {
	Mounted_on  string `json:"MOUNTED_ON"`
	File_system string `json:"FILE_SYSTEM"`
	Size        string `json:"SIZE"`
	Used        string `json:"USED"`
	Avail       string `json:"AVAIL"`
	Use_percent string `json:"USE_PERCENT"`
}

var BashDfH_Template string = `Value FILE_SYSTEM (\S+)
Value SIZE (\S+)
Value USED (\S+)
Value AVAIL (\S+)
Value USE_PERCENT (\S+)
Value MOUNTED_ON (\S+)

Start
  ^Filesystem.*Mounted on$$ -> Data
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"  
  
Data
  ^${FILE_SYSTEM}\s+${SIZE}\s+${USED}\s+${AVAIL}\s+${USE_PERCENT}\s+${MOUNTED_ON} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
EOF  
`
