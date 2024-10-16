package mikrotik_routeros

type UserActivePrint struct {
	When    string `json:"WHEN"`
	Name    string `json:"NAME"`
	Address string `json:"ADDRESS"`
	Via     string `json:"VIA"`
	Id      string `json:"ID"`
}

var UserActivePrint_Template string = `Value ID (\d+)
Value WHEN ([a-z]{3}/\d{1,2}/\d{4}\s+\d{1,2}:\d{1,2}:\d{1,2})
Value NAME ([a-zA-Z0-9*_.@]+)
Value ADDRESS (\S+)
Value VIA (local|telnet|ssh|winbox|api|web|tikapp|ftp|dude)

Start
  ^\s*Flags:\s+R\s+-\s+radius,\s+M\s+-\s+by-romon\s*$$
  ^\s*#\s+WHEN\s+NAME\s+ADDRESS\s+VIA\s*$$ -> ActiveUsersTable
  ^\s*$$
  ^. -> Error

ActiveUsersTable
  ^\s*${ID}\s+${WHEN}\s+${NAME}\s+${ADDRESS}\s+${VIA}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
