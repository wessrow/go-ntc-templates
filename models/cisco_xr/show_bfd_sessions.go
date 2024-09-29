package cisco_xr 

type ShowBfdSessions struct {
	Interface	string	`json:"INTERFACE"`
	Dstaddress	string	`json:"DSTADDRESS"`
	State	string	`json:"STATE"`
}

var ShowBfdSessions_Template = `Value INTERFACE (.+?)
Value DSTADDRESS (\d+\.\d+\.\d+\.\d+)
Value STATE (\S+)

Start
  ^\-+ -> Begin1

Begin1
  ^\s*${INTERFACE}\s+${DSTADDRESS}\s+\S+\s+\S+\s+${STATE} -> Record

`