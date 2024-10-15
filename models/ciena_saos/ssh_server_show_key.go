package ciena_saos

type SshServerShowKey struct {
	Key_status string `json:"KEY_STATUS"`
	Username   string `json:"USERNAME"`
}

var SshServerShowKey_Template string = `Value USERNAME (\S+)
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
