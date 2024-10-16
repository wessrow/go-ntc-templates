package cisco_nxos

type ShowInventory struct {
	Pid   string `json:"PID"`
	Vid   string `json:"VID"`
	Sn    string `json:"SN"`
	Name  string `json:"NAME"`
	Descr string `json:"DESCR"`
}

var ShowInventory_Template string = `Value NAME (.*)
Value DESCR (.*)
Value PID ([^,]\S+)
Value VID ([\d+\w-]*)
Value SN ([\d+\w+/]+)

Start
  ^NAME:\s+"${NAME}",\s+DESCR:\s+"${DESCR}"
  ^NAME:\s+${NAME},\s+DESCR:\s+${DESCR}
  ^PID:\s+${PID}.*,.*VID:\s+${VID}.*SN:\s+${SN} -> Record
  ^PID:\s+${PID}.*,.*VID:\s+${VID}.*SN: -> Record
  ^PID:\s+,.*VID:\s+${VID}.*SN:\s+${SN} -> Record
  ^PID:\s+,.*VID:\s+${VID}.*SN: -> Record
  ^PID:\s+${PID}.*,.*VID:\s+${VID}.*
  ^PID:\s+,.*VID:\s+${VID}.*
  ^.*SN:\s+${SN} -> Record
  ^.*SN: -> Record
`
