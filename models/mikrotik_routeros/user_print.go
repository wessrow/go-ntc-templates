package mikrotik_routeros

type UserPrint struct {
	Group          string `json:"GROUP"`
	Address        string `json:"ADDRESS"`
	Last_logged_in string `json:"LAST_LOGGED_IN"`
	Id             string `json:"ID"`
	Comment        string `json:"COMMENT"`
	Name           string `json:"NAME"`
}

var UserPrint_Template string = `Value ID (\d+)
Value COMMENT (.*)
Value NAME ([a-zA-Z0-9*_.@]+)
Value GROUP (\S+)
Value ADDRESS (\S+)
Value LAST_LOGGED_IN ([a-z]{3}/\d{1,2}/\d{4} \d{1,2}:\d{1,2}:\d{1,2})

Start
  ^\s*Flags:\s+X\s+-\s+disabled\s*$$
  ^\s*#\s+NAME\s+GROUP\s+ADDRESS\s+LAST-LOGGED-IN\s*$$ -> UsersTable
  ^\s*$$
  ^. -> Error

UsersTable
  ^\s*${ID}\s+${NAME}\s+${GROUP}\s+(?:${ADDRESS})?\s+${LAST_LOGGED_IN}\s*$$ -> Record
  ^\s*${ID}\s+;;;\s+${COMMENT}\s*$$
  ^\s*${NAME}\s+${GROUP}\s+(?:${ADDRESS})?\s+${LAST_LOGGED_IN}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
