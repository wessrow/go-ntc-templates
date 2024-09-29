package ciena_saos 

type SshServerShowKey struct {
	Username	string	`json:"USERNAME"`
	Key_status	string	`json:"KEY_STATUS"`
}

var SshServerShowKey_Template = `Value USERNAME (\S+)
Value KEY_STATUS (\S+)

Start
  # Accounts for tabular delineations and table names
  ^\+\-
  # Account for SSH Server Key block
  ^\|\s+[Kk]ey
  ^\|\s+[Uu]sername.*[Kk]ey\s+[Ss]tatus
  ^\|\s*${USERNAME}\s*\S*\s*\|\s*${KEY_STATUS}\s*\| -> Record
  ^\s*$$
  ^. -> Error`