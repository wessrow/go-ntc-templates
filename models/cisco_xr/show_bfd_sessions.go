package cisco_xr

type ShowBfdSessions struct {
	Dstaddress string `json:"DSTADDRESS"`
	State      string `json:"STATE"`
	Interface  string `json:"INTERFACE"`
}

var ShowBfdSessions_Template string = `Value INTERFACE (.+?)
Value DSTADDRESS (\d+\.\d+\.\d+\.\d+)
Value STATE (\S+)

Start
  ^\-+ -> Begin1

Begin1
  ^\s*${INTERFACE}\s+${DSTADDRESS}\s+\S+\s+\S+\s+${STATE} -> Record

`
