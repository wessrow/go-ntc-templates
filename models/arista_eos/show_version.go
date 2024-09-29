package arista_eos 

type ShowVersion struct {
	Model	string	`json:"MODEL"`
	Hw_version	string	`json:"HW_VERSION"`
	Serial_number	string	`json:"SERIAL_NUMBER"`
	Sys_mac	string	`json:"SYS_MAC"`
	Image	string	`json:"IMAGE"`
	Total_memory	string	`json:"TOTAL_MEMORY"`
	Free_memory	string	`json:"FREE_MEMORY"`
}

var ShowVersion_Template = `Value MODEL (\S+)
Value HW_VERSION (\S+)
Value SERIAL_NUMBER (\S+)
Value SYS_MAC (\S+)
Value IMAGE (\S+)
Value TOTAL_MEMORY (\d+)
Value FREE_MEMORY (\d+)

Start
  ^Arista\s+${MODEL}
  ^Hardware\s+version:\s+${HW_VERSION}
  ^Serial\s+number:\s+${SERIAL_NUMBER}
  ^System\s+MAC\s+address:\s+${SYS_MAC}
  ^Software\s+image\s+version:\s+${IMAGE}
  ^Total\s+memory:\s+${TOTAL_MEMORY}
  ^Free\s+memory:\s+${FREE_MEMORY} -> Record
`