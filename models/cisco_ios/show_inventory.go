package cisco_ios

type ShowInventory struct {
	Name  string `json:"NAME"`
	Descr string `json:"DESCR"`
	Pid   string `json:"PID"`
	Vid   string `json:"VID"`
	Sn    string `json:"SN"`
}

var ShowInventory_Template string = `Value NAME (.*)
Value DESCR (.*)
Value PID (([\S+]+|.*))
Value VID (.*)
Value SN ([\w+\d+]+)

Start
  ^NAME:\s+"${NAME}",\s+DESCR:\s+"${DESCR}"
  ^PID:\s+${PID}.*,.*VID:\s+${VID},.*SN:\s+${SN} -> Record
  ^PID:\s+,.*VID:\s+${VID},.*SN: -> Record
  ^PID:\s+${PID}.*,.*VID:\s+${VID},.*SN: -> Record
  ^PID:\s+,.*VID:\s+${VID},.*SN:\s+${SN} -> Record
  ^PID:\s+${PID}.*,.*VID:\s+${VID}.*
  ^PID:\s+,.*VID:\s+${VID}.*
  ^.*SN:\s+${SN} -> Record
  ^.*SN: -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`
