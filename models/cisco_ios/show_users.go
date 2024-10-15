package cisco_ios

type ShowUsers struct {
	Location string `json:"LOCATION"`
	Line     string `json:"LINE"`
	User     string `json:"USER"`
	Hosts    string `json:"HOSTS"`
	Idle     string `json:"IDLE"`
}

var ShowUsers_Template string = `Value LINE ((\*\s+)?\d+\s+(con|aux|tty|vty)\s+\d+)
Value USER (\S+)
Value HOSTS (\S+)
Value IDLE (\S+)
Value LOCATION (\S+)

Start
  ^\s*Line\s+User\s+Host\(s\)\s+Idle\s+Location\s*$$ -> UsersTable
  ^\s*$$
  ^. -> Error

UsersTable
  ^\s*${LINE}(?:\s+${USER})?(?:\s+${HOSTS})?\s+${IDLE}(?:\s+${LOCATION})?\s*$$ -> Record
  ^\s*Interface\s+User\s+Mode\s+Idle\s+Peer\s+Address\s*$$
  ^\s*$$
  ^. -> Error
`
