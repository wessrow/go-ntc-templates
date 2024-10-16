package arista_eos

type ShowInventory struct {
	Port  string `json:"PORT"`
	Pid   string `json:"PID"`
	Sn    string `json:"SN"`
	Descr string `json:"DESCR"`
	Vid   string `json:"VID"`
}

var ShowInventory_Template string = `Value PORT ([0-9\/]+)
Value PID (\S+)
Value SN (\S+)
Value DESCR (.+)
Value VID (\d+\.\d+|\d+)

Start
  ^\s*System\s+\S+?$$ -> Chassis

Chassis
  ^\s+Model
  ^\s+-
  ^\s+HW
  ^\s+${VID}\s+${SN}\s+\d+-\d+-\d+ -> Record
  ^\s+${PID}?\s+${DESCR}$$
  ^\s*System.+(power supply|power-supply) -> Power_Supply


Power_Supply
  ^\s+Slot
  ^\s+-
  ^\s+${PORT}\s+${PID}\s+${SN} -> Record
  ^\s*System.+(fan) -> Fan

Fan
  ^\s+Module
  ^\s+-
  ^\s+${PORT}?\s+\d+?\s+${PID}?\s+${SN} -> Record
  ^\s*System.+ports -> Ports

Ports
  ^\s+${DESCR}\s+${PORT} -> Record
  ^\s*System.+transceiver -> Transceiver

Transceiver
  ^\s+${PORT}\s+${DESCR}\s+${PID}\s+${SN}\s+${VID} -> Record
`
