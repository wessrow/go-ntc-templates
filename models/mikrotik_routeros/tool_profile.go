package mikrotik_routeros 

type ToolProfile struct {
	Name	string	`json:"NAME"`
	Cpu	string	`json:"CPU"`
	Usage	string	`json:"USAGE"`
}

var ToolProfile_Template = `Value NAME (\S+)
Value CPU (\S+)
Value USAGE (\d+|\d+.\d+)

Start
  ^\s*NAME\s+CPU\s+USAGE\s*$$ -> ProfilingTable
  ^\s*$$
  ^. -> Error

ProfilingTable
  ^\s*${NAME}\s*(\s+${CPU})?\s+${USAGE}%\s*$$ -> Record
  ^\s*--\s+\[Q\s+quit\|D\s+dump\|C-z\s+pause\]\s*$$
  ^\s*--\s+\[Q\s+quit\|C-z\s+pause\]\s*$$
  ^\s*$$
  ^. -> Error
`